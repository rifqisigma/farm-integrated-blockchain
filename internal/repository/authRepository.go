package repository

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"

	"gorm.io/gorm"
)

type AuthRepository interface {

	//gmail traditional
	Login(input *dto.Login) (*dto.LoginResponse, error)
	Register(input *dto.Register) (*dto.RegisterResponse, error)
	ValidateUser(id uint, email string) (bool, error)
	ChangePassword(email, password string) error
	CheckUserExist(email string) (bool, error)
}
type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(input *dto.Register) (*dto.RegisterResponse, error) {

	newUser := entity.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     entity.Consumer,
		Provider: "gmail",
	}

	if err := r.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	response := dto.RegisterResponse{
		Id:       newUser.ID,
		Email:    newUser.Email,
		Verified: newUser.IsVerified,
		Provider: newUser.Provider,
	}

	return &response, nil
}

func (r *authRepository) Login(input *dto.Login) (*dto.LoginResponse, error) {
	var user entity.User

	if err := r.db.Where("email = ? AND is_verified = ?", input.Email, true).First(&user).Error; err != nil {
		return nil, err
	}

	result := dto.LoginResponse{
		Id:           user.ID,
		Email:        user.Email,
		Role:         string(user.Role),
		IsVerified:   user.IsVerified,
		PasswordHash: user.Password,
	}

	return &result, nil
}

func (r *authRepository) ValidateUser(id uint, email string) (bool, error) {
	if err := r.db.Model(&entity.User{}).Where("id = ? AND email = ? AND is_verified = ?", id, email, false).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, helper.ErrUserNotFound
		}

		return false, err
	}

	return true, nil
}

func (r *authRepository) ChangePassword(email, password string) error {
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *authRepository) CheckUserExist(email string) (bool, error) {
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, helper.ErrUserNotFound
		}
		return false, err
	}

	return true, nil
}

func (r *authRepository) DeleteAcccount(email string) error {
	if err := r.db.Where("email = ?", email).Delete(&entity.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}
