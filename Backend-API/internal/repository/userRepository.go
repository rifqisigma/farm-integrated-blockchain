package repository

import (
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	Me(ctx context.Context, id uint) (*dto.GetUser, error)
	ChangePassword(ctx context.Context, email, password string) error
	CreateProfile(ctx context.Context, input *dto.CreateProfileRequest) error
	UpdateProfile(ctx context.Context, input *dto.UpdateProfileRequest) error
	UpdateRole(ctx context.Context, input *dto.UpdateRoleRequest) error
	CheckUserExist(ctx context.Context, userId uint, email string) (bool, error)
}
type userRepository struct {
	db       *gorm.DB
	redis    *redis.Client
	authRepo AuthRepository
}

func NewUserRepository(db *gorm.DB, redis *redis.Client, authRepo AuthRepository) UserRepository {
	return &userRepository{db, redis, authRepo}
}

func (r *userRepository) ChangePassword(ctx context.Context, email, password string) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *userRepository) CheckUserExist(ctx context.Context, userId uint, email string) (bool, error) {
	return r.authRepo.CheckUserExist(ctx, email)
}

// err
func (r *userRepository) CreateProfile(ctx context.Context, input *dto.CreateProfileRequest) error {
	switch input.Role {
	case entity.Consumer:
		newConsumer := entity.ConsumerProfile{
			UserID: input.UserId,
			Name:   input.Name,
		}
		if err := r.db.WithContext(ctx).Model(&entity.ConsumerProfile{}).Create(&newConsumer).Error; err != nil {
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
		if err := r.db.WithContext(ctx).Model(&entity.DistributorProfile{}).Create(&newDistributor).Error; err != nil {
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
		if err := r.db.WithContext(ctx).Model(&entity.FarmerProfile{}).Create(&newFarmer).Error; err != nil {
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
		if err := r.db.WithContext(ctx).Model(&entity.RetailerProfile{}).Create(&newRetailer).Error; err != nil {
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

func (r *userRepository) UpdateProfile(ctx context.Context, input *dto.UpdateProfileRequest) error {

	switch input.Role {
	case entity.Consumer:
		if err := r.db.WithContext(ctx).Model(&entity.ConsumerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	case entity.Distributor:
		if err := r.db.WithContext(ctx).Model(&entity.DistributorProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	case entity.Farmer:
		if err := r.db.WithContext(ctx).Model(&entity.FarmerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {

			return err
		}
	case entity.Retailer:
		if err := r.db.WithContext(ctx).Model(&entity.RetailerProfile{}).Where("user_id = ?", input.UserId).UpdateColumns(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	default:
		return gorm.ErrRecordNotFound
	}

	key := fmt.Sprintf("user:%d", input.UserId)
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateRole(ctx context.Context, input *dto.UpdateRoleRequest) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()
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

	if err := r.authRepo.UpdateRevokeToken(ctx, input.UserId); err != nil {
		tx.Rollback()
		return err
	}

	key := fmt.Sprintf("user:%d", input.UserId)
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepository) Me(ctx context.Context, id uint) (*dto.GetUser, error) {
	var result dto.GetUser
	key := fmt.Sprintf("user:%d", id)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil && val != "" {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return &result, nil
		}
	}
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	res := tx.Model(&entity.User{}).Select("id as id", "email as email", "is_verified as is_verified", "role as role", "provider as provider").Where("id = ?", id).Select(&result)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	switch result.Role {
	case entity.Consumer:
		if err := tx.Model(&entity.ConsumerProfile{}).Select("id", "name").Where("user_id = ?", result.Id).Scan(&result.Data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Distributor:
		if err := tx.Model(&entity.ConsumerProfile{}).Select("id", "name").Where("user_id = ?", result.Id).Scan(&result.Data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Retailer:
		if err := tx.Model(&entity.ConsumerProfile{}).Select("id", "name").Where("user_id = ?", result.Id).Scan(&result.Data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Farmer:
		if err := tx.Model(&entity.ConsumerProfile{}).Select("id", "name").Where("user_id = ?", result.Id).Scan(&result.Data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	jsonData, _ := json.Marshal(result)

	if err := r.redis.Set(ctx, key, jsonData, 30*time.Minute).Err(); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &result, nil
}
