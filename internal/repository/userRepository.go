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
	DeleteAccount(userId uint) error
	CheckUserExist(userId uint, email string) (bool, error)

	CreateConsumerProfile(input *dto.CreateConsumerProfile) error
	CreateFarmerProfile(input *dto.CreateFarmerProfile) error
	CreateDistributorProfile(input *dto.CreateDistributorProfile) error
	CreateRetailerProfile(input *dto.CreateRetailerProfile) error

	UpdateConsumerProfile(input *dto.UpdateConsumerProfile) error
	UpdateFarmerProfile(input *dto.UpdateFarmerProfile) error
	UpdateDistributorProfile(input *dto.UpdateDistributorProfile) error
	UpdateRetailerProfile(input *dto.UpdateRetailerProfile) error
	DeleteToken(userId uint) error
	ValidateToken(userId uint) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) DeleteAccount(userId uint) error {
	if err := r.db.Where("id = ?", userId).Delete(&entity.User{}).Error; err != nil {
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

	if err := r.db.Model(&entity.User{}).Where("id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, helper.ErrUserNotFound
		}
		return false, err
	}

	return true, nil
}

func (r *userRepository) CreateConsumerProfile(input *dto.CreateConsumerProfile) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Where("user_id = ? AND role = ?", input.UserId, entity.Consumer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return helper.ErrUserNotFound
		}
		tx.Rollback()
		return err
	}

	if err := tx.Create(&entity.ConsumerProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepository) CreateFarmerProfile(input *dto.CreateFarmerProfile) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Where("user_id = ? AND role = ?", input.UserId, entity.Farmer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return helper.ErrUserNotFound
		}
		tx.Rollback()
		return err
	}

	if err := r.db.Create(&entity.FarmerProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepository) CreateDistributorProfile(input *dto.CreateDistributorProfile) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Where("user_id = ? AND role = ?", input.UserId, entity.Distributor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return helper.ErrUserNotFound
		}
		tx.Rollback()
		return err
	}

	if err := r.db.Create(&entity.DistributorProfile{
		UserID: input.UserId,
		Name:   input.Name,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepository) CreateRetailerProfile(input *dto.CreateRetailerProfile) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Where("user_id = ? AND role = ?", input.UserId, entity.Retailer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return helper.ErrUserNotFound
		}
		tx.Rollback()
		return err
	}

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

func (r *userRepository) DeleteToken(userId uint) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ?", userId).Delete(&entity.Token{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return helper.ErrUserNotFound
		}

		return err
	}

	return nil

}

func (r *userRepository) ValidateToken(userId uint) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ? AND is_revoked = ?", userId, r.db.NowFunc().Local().Second()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}
