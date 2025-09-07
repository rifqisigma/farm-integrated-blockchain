package repository

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	ChangePassword(email, password string) error
	CheckUserExist(userId uint, email string) (bool, error)
	CreateProfile(input *dto.CreateProfileRequest) error
	UpdateProfile(input *dto.UpdateProfileRequest) error
	UpdateRole(input *dto.UpdateRoleRequest) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) ChangePassword(email, password string) error {
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) CheckUserExist(userId uint, email string) (bool, error) {
	if email != "" {
		if err := r.db.Model(&entity.User{}).Where("email = ?", email).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, gorm.ErrRecordNotFound
			}
			return false, err
		}

		return true, nil
	}

	if err := r.db.Model(&entity.User{}).Where("id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}

	return true, nil
}

// err
func (r *userRepository) CreateProfile(input *dto.CreateProfileRequest) error {
	switch input.Role {
	case entity.Consumer:
		newConsumer := entity.ConsumerProfile{
			UserID: input.UserId,
			Name:   input.Name,
		}
		if err := r.db.Model(&entity.ConsumerProfile{}).Create(&newConsumer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return gorm.ErrDuplicatedKey
			}
			return err
		}
	case entity.Distributor:
		newDistributor := entity.DistributorProfile{
			UserID: input.UserId,
			Name:   input.Name,
		}
		if err := r.db.Model(&entity.DistributorProfile{}).Create(&newDistributor).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return gorm.ErrDuplicatedKey
			}
			return err
		}
	case entity.Farmer:
		newFarmer := entity.FarmerProfile{
			Name:   input.Name,
			UserID: input.UserId,
		}
		if err := r.db.Model(&entity.FarmerProfile{}).Create(&newFarmer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return gorm.ErrDuplicatedKey
			}
			return err
		}
	case entity.Retailer:
		newRetailer := entity.RetailerProfile{
			UserID: input.UserId,
			Name:   input.Name,
		}
		if err := r.db.Model(&entity.RetailerProfile{}).Create(&newRetailer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return gorm.ErrDuplicatedKey
			}
			return err
		}
	default:
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) UpdateProfile(input *dto.UpdateProfileRequest) error {

	switch input.Role {
	case entity.Consumer:
		if err := r.db.Model(&entity.ConsumerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	case entity.Distributor:
		if err := r.db.Model(&entity.DistributorProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	case entity.Farmer:
		if err := r.db.Model(&entity.FarmerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {

			return err
		}
	case entity.Retailer:
		if err := r.db.Model(&entity.RetailerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	default:
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) DeleteToken(userId uint) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ?", userId).Delete(&entity.Token{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return gorm.ErrRecordNotFound
		}

		return err
	}

	return nil

}

func (r *userRepository) UpdateRole(input *dto.UpdateRoleRequest) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Where("id = ? AND role = ?", input.UserId, entity.Consumer).UpdateColumn("role", input.NewRole.Role).Error; err != nil {
		tx.Rollback()
		return err
	}

	switch input.OldRole {
	case entity.Distributor:
		if err := tx.Model(&entity.DistributorProfile{}).Where("user_id = ?", input.UserId).Delete(&entity.DistributorProfile{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		switch input.NewRole.Role {
		case entity.Farmer:
			newFarmer := entity.FarmerProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.FarmerProfile{}).Create(&newFarmer).Error; err != nil {
				tx.Rollback()
				return err
			}
		case entity.Retailer:
			newRetailer := entity.RetailerProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.RetailerProfile{}).Create(&newRetailer).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

	case entity.Farmer:
		if err := tx.Model(&entity.FarmerProfile{}).Where("user_id = ?", input.UserId).Delete(&entity.FarmerProfile{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		switch input.NewRole.Role {
		case entity.Distributor:
			newDistributor := entity.DistributorProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.DistributorProfile{}).Create(&newDistributor).Error; err != nil {
				tx.Rollback()
				return err
			}
		case entity.Retailer:
			newRetailer := entity.RetailerProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.RetailerProfile{}).Create(&newRetailer).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	case entity.Retailer:
		if err := tx.Model(&entity.RetailerProfile{}).Where("user_id = ?", input.UserId).Delete(&entity.RetailerProfile{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		switch input.NewRole.Role {
		case entity.Distributor:
			newDistributor := entity.DistributorProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.DistributorProfile{}).Create(&newDistributor).Error; err != nil {
				tx.Rollback()
				return err
			}
		case entity.Farmer:
			newFarmer := entity.FarmerProfile{
				UserID: input.UserId,
				Name:   input.NewRole.Name,
			}
			if err := tx.Model(&entity.FarmerProfile{}).Create(&newFarmer).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

	default:
		return gorm.ErrRecordNotFound
	}

	return nil
}
