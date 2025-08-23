package usecase

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"
)

type AuthUsecase interface {
	Register(input *dto.Register) error
	Login(input *dto.Login) (string, string, error)
	ValidateUser(token string) error
	RefreshLongToken(token string) (string, error)
	ResetPassword(input *dto.UserNewPassword) error
	RequestResetPassword(email string) error
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{authRepo}
}

func (u *authUsecase) Register(input *dto.Register) error {
	hashpw := helper.HashPassword(input.Password)

	input.Password = hashpw
	dataUser, err := u.authRepo.Register(input)
	if err != nil {
		return err
	}

	tokenJwt, err := helper.GenerateJWT(dataUser.Email, dataUser.Provider, dataUser.Id, dataUser.Verified)
	if err != nil {
		return err
	}

	helper.SendEmailValidateEmail(dataUser.Email, tokenJwt)
	return nil
}

func (u *authUsecase) Login(input *dto.Login) (string, string, error) {
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

	tokenJwt, err := helper.GenerateJWT(dataUser.Email, dataUser.Role, dataUser.Id, dataUser.IsVerified)
	if err != nil {
		return "", "", err
	}

	tokenJwtLongExp, err := helper.GenerateJWTLongExp(dataUser.Id, dataUser.IsVerified)
	if err != nil {
		return "", "", err
	}
	return tokenJwt, tokenJwtLongExp, nil
}

func (u *authUsecase) ValidateUser(token string) error {
	userClaims, err := helper.ParseJWT(token)
	if err != nil {
		return err
	}

	valid, err := u.authRepo.ValidateUser(userClaims.UserID, userClaims.Email)
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

func (u *authUsecase) ResetPassword(input *dto.UserNewPassword) error {
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
	isExist, err := u.authRepo.CheckUserExist(email)
	if err != nil {
		return err
	}

	if !isExist {
		return helper.ErrUserNotFound
	}
	token, err := helper.GenerateJWTShortExp(email)
	if err != nil {
		return err
	}

	helper.SendEmailResetPassword(email, token)
	return nil
}
