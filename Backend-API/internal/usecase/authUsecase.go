package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"

	"farm-integrated-web3/utils/helper"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	Register(ctx context.Context, input *dto.RegisterRequest) error
	Login(ctx context.Context, input *dto.LoginRequest) (string, string, error)
	ValidateUser(ctx context.Context, token string) error
	RefreshLongToken(ctx context.Context, userId uint, validate bool) (string, error)
	ResetPassword(ctx context.Context, input *dto.UserResetPasswordRequest) error
	RequestResetPassword(ctx context.Context, email string) error
	ResendVerificationEmail(ctx context.Context, email string) error
	CreateAccessToken(ctx context.Context, id uint) (string, error)

	Logout(ctx context.Context, userId uint) error
	DeleteAccount(ctx context.Context, userId uint, role string) error
}

type authUsecase struct {
	authRepo repository.AuthRepository
	redis    *redis.Client
}

func NewAuthUsecase(authRepo repository.AuthRepository, redis *redis.Client) AuthUsecase {
	return &authUsecase{authRepo, redis}
}

func (u *authUsecase) Register(ctx context.Context, input *dto.RegisterRequest) error {
	hashpw := helper.HashPassword(input.Password)

	input.Password = hashpw
	dataUser, err := u.authRepo.Register(ctx, input)
	if err != nil {
		return err
	}

	tokenJwt, err := helper.GenerateJWTShortExp(dataUser.Email)
	if err != nil {
		return err
	}

	key := fmt.Sprintln("behind:pending:email_verify")
	mapping := map[string]interface{}{
		"op": "email_verify",
		"data": map[string]interface{}{
			"email": dataUser.Email,
			"token": tokenJwt,
		},
	}

	dataMap, err := json.Marshal(mapping)
	if err != nil {
		return err
	}
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataMap).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)

	return nil
}

func (u *authUsecase) Login(ctx context.Context, input *dto.LoginRequest) (string, string, error) {
	dataUser, err := u.authRepo.Login(ctx, input)
	if err != nil {
		return "", "", err
	}

	valid, err := helper.ValidateToken(dataUser.PasswordHash, input.Password)
	if err != nil {
		return "", "", err
	}

	if !valid {
		return "", "", helper.ErrLoginNotSuccess
	}

	tokenJwt, err := helper.GenerateJWT(dataUser.Email, string(dataUser.Role), dataUser.Id, dataUser.ProfileId, dataUser.IsVerified)
	fmt.Printf("token: %s", tokenJwt)

	if err != nil {

		return "", "", err
	}

	tokenJwtLongExp, err := helper.GenerateJWTLongExp(dataUser.Id, dataUser.IsVerified)
	fmt.Printf("token besar: %s", tokenJwtLongExp)
	if err != nil {
		return "", "", err
	}

	if err := u.authRepo.CreateToken(ctx, dataUser.Id, tokenJwt); err != nil {
		return "", "", err
	}

	return tokenJwt, tokenJwtLongExp, nil
}

func (u *authUsecase) ValidateUser(ctx context.Context, token string) error {
	userClaims, err := helper.ParseJWTShortExp(token)
	if err != nil {
		return err
	}

	valid, err := u.authRepo.ValidateUser(ctx, userClaims.Email)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("user already verified")
	}

	return nil
}

func (u *authUsecase) RefreshLongToken(ctx context.Context, userId uint, validate bool) (string, error) {
	return helper.GenerateJWTLongExp(userId, validate)
}

func (u *authUsecase) ResetPassword(ctx context.Context, input *dto.UserResetPasswordRequest) error {
	jwtClaims, err := helper.ParseJWTShortExp(input.Token)
	if err != nil {
		return helper.ErrInvalidToken
	}
	if input.ConfirmNewPassword != input.NewPassword {
		return helper.ErrBadRequest
	}
	if err := u.authRepo.ChangePassword(ctx, jwtClaims.Email, input.NewPassword); err != nil {
		return err
	}

	return nil
}

func (u *authUsecase) RequestResetPassword(ctx context.Context, email string) error {
	isExist, err := u.authRepo.CheckUserExist(ctx, email)
	if err != nil {
		return err
	}

	if !isExist {
		return gorm.ErrRecordNotFound
	}
	token, err := helper.GenerateJWTShortExp(email)
	if err != nil {
		return err
	}

	key := fmt.Sprintln("behind:pending:reset_password")
	mapping := map[string]interface{}{
		"op": "reset_password",
		"data": map[string]interface{}{

			"email": email,
			"token": token,
		},
	}

	dataMap, err := json.Marshal(mapping)
	if err != nil {
		return err
	}
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataMap).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}

func (u *authUsecase) ResendVerificationEmail(ctx context.Context, email string) error {
	isExist, err := u.authRepo.CheckUserExist(ctx, email)
	if err != nil {
		return err
	}

	if !isExist {
		return gorm.ErrRecordNotFound
	}
	token, err := helper.GenerateJWTShortExp(email)
	if err != nil {
		return err
	}

	key := fmt.Sprintln("behind:pending:email_verify")
	mapping := map[string]interface{}{
		"op": "email_verify",
		"data": map[string]interface{}{
			"email": email,
			"token": token,
		},
	}

	dataMap, err := json.Marshal(mapping)
	if err != nil {
		return err
	}
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataMap).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}

func (u *authUsecase) CreateAccessToken(ctx context.Context, id uint) (string, error) {
	dataUser, err := u.authRepo.GetUserInfo(ctx, id)
	if err != nil {
		return "", err
	}

	token, err := helper.GenerateJWT(dataUser.Email, string(dataUser.Role), dataUser.Id, dataUser.ProfileId, dataUser.IsVerified)
	if err != nil {
		return "", err
	}

	if err := u.authRepo.CreateToken(ctx, dataUser.Id, token); err != nil {
		return "", err
	}

	return token, nil
}

func (u *authUsecase) Logout(ctx context.Context, userId uint) error {
	return u.authRepo.UpdateRevokeToken(ctx, userId)

}

func (u *authUsecase) DeleteAccount(ctx context.Context, userId uint, role string) error {
	if err := u.authRepo.DeleteAccount(ctx, userId, role); err != nil {
		return err
	}

	return nil
}
