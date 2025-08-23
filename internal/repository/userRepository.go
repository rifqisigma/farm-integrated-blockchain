package repository

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"

	"gorm.io/gorm"
)

type UserRepository interface {
	ChangePassword(email, password string) error
	DeleteAccount(email string) error
	CheckUserExist(userId uint, email string) (bool, error)

	//profile
	CreateConsumerProfile(input *dto.CreateConsumerProfile) error
	CreateFarmerProfile(input *dto.CreateFarmerProfile) error
	CreateDistributorProfile(input *dto.CreateDistributorProfile) error
	CreateRetailerProfile(input *dto.CreateRetailerProfile) error

	UpdateConsumerProfile(input *dto.UpdateConsumerProfile) error
	UpdateFarmerProfile(input *dto.UpdateFarmerProfile) error
	UpdateDistributorProfile(input *dto.UpdateDistributorProfile) error
	UpdateRetailerProfile(input *dto.UpdateRetailerProfile) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) DeleteAccount(email string) error {
	if err := r.db.Where("email = ?", email).Delete(&entity.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) ChangePassword(email, password string) error {
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) CheckUserExist(userId uint, email string) (bool, error) {
	if email != "" {
		if err := r.db.Model(&entity.User{}).Where("email = ?", email).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, helper.ErrUserNotFound
			}
			return false, err
		}

		return true, nil
	}

	if err := r.db.Model(&entity.User{}).Where("user_id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, helper.ErrUserNotFound
		}
		return false, err
	}

	return true, nil
}

func (r *userRepository) CreateConsumerProfile(input *dto.CreateConsumerProfile) error {
	if err := r.db.Create(&entity.ConsumerProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CreateFarmerProfile(input *dto.CreateFarmerProfile) error {
	if err := r.db.Create(&entity.FarmerProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CreateDistributorProfile(input *dto.CreateDistributorProfile) error {
	if err := r.db.Create(&entity.DistributorProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CreateRetailerProfile(input *dto.CreateRetailerProfile) error {
	if err := r.db.Create(&entity.RetailerProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateConsumerProfile(input *dto.UpdateConsumerProfile) error {
	if err := r.db.Model(&entity.ConsumerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.ConsumerProfile{
		Name: input.Name,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) UpdateFarmerProfile(input *dto.UpdateFarmerProfile) error {
	if err := r.db.Model(&entity.FarmerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.FarmerProfile{
		Name: input.Name,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) UpdateDistributorProfile(input *dto.UpdateDistributorProfile) error {
	if err := r.db.Model(&entity.DistributorProfile{}).Where("user_id = ?", input.UserId).Updates(entity.DistributorProfile{
		Name: input.Name,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) UpdateRetailerProfile(input *dto.UpdateRetailerProfile) error {
	if err := r.db.Model(&entity.RetailerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.RetailerProfile{
		Name: input.Name,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}
	return nil
}
