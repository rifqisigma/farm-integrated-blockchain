package usecase

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Register(input *dto.RegisterRequest) error
	Login(input *dto.LoginRequest) (string, string, error)
	ValidateUser(token string) error
	RefreshLongToken(token string) (string, error)
	ResetPassword(input *dto.UserResetPasswordRequest) error
	RequestResetPassword(email string) error
	ResendVerificationEmail(email string) error
	CreateAccessToken(id uint) (string, error)
	GetUserInfo(id uint) (*dto.LoginResponse, error)
	Logout(userId uint) error
	DeleteAccount(userId uint) error
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{authRepo}
}

func (u *authUsecase) Register(input *dto.RegisterRequest) error {
	hashpw := helper.HashPassword(input.Password)

	input.Password = hashpw
	dataUser, err := u.authRepo.Register(input)
	if err != nil {
		return err
	}

	tokenJwt, err := helper.GenerateJWTShortExp(dataUser.Email)
	if err != nil {
		return err
	}

	helper.SendEmailValidateEmail(dataUser.Email, tokenJwt)
	return nil
}

func (u *authUsecase) Login(input *dto.LoginRequest) (string, string, error) {
	dataUser, err := u.authRepo.Login(input)
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
	if err != nil {
		return "", "", err
	}

	tokenJwtLongExp, err := helper.GenerateJWTLongExp(dataUser.Id, dataUser.IsVerified)
	if err != nil {
		return "", "", err
	}

	if err := u.authRepo.CreateToken(dataUser.Id, tokenJwtLongExp); err != nil {
		return "", "", err
	}

	return tokenJwt, tokenJwtLongExp, nil
}

func (u *authUsecase) ValidateUser(token string) error {
	userClaims, err := helper.ParseJWTShortExp(token)
	if err != nil {
		return err
	}

	valid, err := u.authRepo.ValidateUser(userClaims.Email)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("user already verified")
	}

	return nil
}

func (u *authUsecase) RefreshLongToken(token string) (string, error) {
	userClaims, err := helper.ParseJWTLongExp(token)
	if err != nil {
		return "", err
	}

	return helper.GenerateJWTLongExp(userClaims.UserID, userClaims.Verified)
}

func (u *authUsecase) ResetPassword(input *dto.UserResetPasswordRequest) error {
	jwtClaims, err := helper.ParseJWTShortExp(input.Token)
	if err != nil {
		return helper.ErrInvalidToken
	}
	if input.ConfirmNewPassword != input.NewPassword {
		return helper.ErrBadRequest
	}
	if err := u.authRepo.ChangePassword(jwtClaims.Email, input.NewPassword); err != nil {
		return err
	}

	return nil
}

func (u *authUsecase) RequestResetPassword(email string) error {
	isExist, err := u.authRepo.CheckUserExist(0, email)
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

	helper.SendEmailResetPassword(email, token)
	return nil
}

func (u *authUsecase) ResendVerificationEmail(email string) error {
	isExist, err := u.authRepo.CheckUserExist(0, email)
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

	helper.SendEmailValidateEmail(email, token)
	return nil
}

func (u *authUsecase) CreateAccessToken(id uint) (string, error) {
	dataUser, err := u.authRepo.GetUserInfo(id)
	if err != nil {
		return "", err
	}
	token, err := helper.GenerateJWT(dataUser.Email, string(dataUser.Role), dataUser.Id, dataUser.ProfileId, dataUser.IsVerified)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *authUsecase) GetUserInfo(id uint) (*dto.LoginResponse, error) {
	result, err := u.authRepo.GetUserInfo(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *authUsecase) Logout(userId uint) error {
	valid, err := u.authRepo.CheckUserExist(userId, "")
	if err != nil {
		return err
	}

	if !valid {
		return gorm.ErrRecordNotFound
	}

	return u.authRepo.UpdateRevokeToken(userId)

}

func (u *authUsecase) DeleteAccount(userId uint) error {
	valid, err := u.authRepo.CheckUserExist(userId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.authRepo.DeleteAccount(userId); err != nil {
		return err
	}

	return nil
}
