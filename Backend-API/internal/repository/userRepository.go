package repository

import (
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"

	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type UserRepository interface {
	Me(ctx context.Context, id uint) (*dto.GetUser, error)
	ChangePassword(ctx context.Context, userId uint, email, password string) error
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

func (r *userRepository) ChangePassword(ctx context.Context, userId uint, email, password string) error {
	tx := r.db.Begin().WithContext(ctx)
	if err := tx.Debug().Model(&entity.User{}).Where("email = ? and is_verified = ?", email, true).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Model(&entity.Token{}).Where("user_id", userId).Delete(&entity.Token{}).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func (r *userRepository) CheckUserExist(ctx context.Context, userId uint, email string) (bool, error) {
	return r.authRepo.CheckUserExist(ctx, email)
}

func (r *userRepository) CreateProfile(ctx context.Context, input *dto.CreateProfileRequest) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var quantity int64
	if res := tx.Debug().Model(&entity.User{}).Where("id = ? AND role = ?", input.UserId, entity.None).Count(&quantity); res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if quantity == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	fmt.Println(input.Role)
	switch input.Role {
	case entity.Consumer:
		newConsumer := entity.ConsumerProfile{
			UserId: input.UserId,
			Name:   input.Name,
		}
		if err := tx.Debug().Model(&entity.ConsumerProfile{}).Create(&newConsumer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	case entity.Distributor:
		newDistributor := entity.DistributorProfile{
			UserId: input.UserId,
			Name:   input.Name,
		}
		if err := tx.Debug().Model(&entity.DistributorProfile{}).Create(&newDistributor).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	case entity.Collector:
		newCollector := entity.CollectorProfile{
			UserId: input.UserId,
			Name:   input.Name,
		}
		if err := tx.Debug().Model(&entity.CollectorProfile{}).Create(&newCollector).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	case entity.Processor:
		newProcessor := entity.ProcessorProfile{
			UserId: input.UserId,
			Name:   input.Name,
		}
		if err := tx.Debug().Model(&entity.ProcessorProfile{}).Create(&newProcessor).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	case entity.Farmer:
		newFarmer := entity.FarmerProfile{
			Name:   input.Name,
			UserId: input.UserId,
		}
		if err := tx.Debug().Model(&entity.FarmerProfile{}).Create(&newFarmer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	case entity.Seller:
		newRetailer := entity.SellerProfile{
			UserId: input.UserId,
			Name:   input.Name,
		}
		if err := tx.Debug().Model(&entity.SellerProfile{}).Create(&newRetailer).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				tx.Rollback()
				return gorm.ErrDuplicatedKey
			}
			tx.Rollback()
			return err
		}
	default:
		return helper.ErrRoleNotFound
	}

	res := tx.Debug().Model(&entity.User{}).Where("id = ?", input.UserId).Update("role", input.Role)
	if res.Error != nil {
		tx.Rollback()
		return res.Error

	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	res2 := tx.Debug().Model(&entity.Token{}).Where("user_id = ?", input.UserId).Delete(entity.Token{})
	if res2.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateProfile(ctx context.Context, input *dto.UpdateProfileRequest) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var quantity int64
	if res := tx.Debug().Model(&entity.User{}).Where("id = ? AND is_verified = ?", input.UserId, true).Count(&quantity); res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if quantity == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	switch input.Role {
	case entity.Consumer:
		if err := tx.Debug().Model(&entity.ConsumerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.ConsumerProfile{
			Name: input.Name,
		}).Error; err != nil {
			return err
		}
	case entity.Distributor:
		if err := tx.Debug().Model(&entity.DistributorProfile{}).Where("user_id = ?", input.UserId).Updates(entity.DistributorProfile{
			Name: input.Name,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	case entity.Farmer:
		if err := tx.Debug().Model(&entity.FarmerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.FarmerProfile{
			Name: input.Name,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	case entity.Seller:
		if err := tx.Debug().Model(&entity.SellerProfile{}).Where("user_id = ?", input.UserId).Updates(entity.SellerProfile{
			Name: input.Name,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	case entity.Processor:
		if err := tx.Debug().Model(&entity.ProcessorProfile{}).Where("user_id = ?", input.UserId).Updates(entity.ProcessorProfile{
			Name: input.Name,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	case entity.Collector:
		if err := tx.Debug().Model(&entity.CollectorProfile{}).Where("user_id = ?", input.UserId).Updates(entity.CollectorProfile{
			Name: input.Name,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}

	default:
		tx.Rollback()
		return helper.ErrRoleNotFound
	}

	res2 := tx.Debug().Model(&entity.Token{}).Where("user_id = ?", input.UserId).Delete(entity.Token{})
	if res2.Error != nil {
		tx.Rollback()
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

	var quantity int64

	if res := tx.Debug().Model(&entity.User{}).Where("id = ? AND role = ?", input.UserId, entity.Consumer).Count(&quantity); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if quantity == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if input.OldRole == entity.Consumer {
		res := tx.Debug().Model(&entity.ConsumerProfile{}).Where("user_id = ?", input.UserId).Delete(&entity.ConsumerProfile{})

		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		if res.RowsAffected == 0 {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}
		switch input.NewRole.Role {
		case entity.Consumer:
			res := tx.Debug().Model(&entity.ConsumerProfile{}).Create(&entity.ConsumerProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})
			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}

			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		case entity.Farmer:
			res := tx.Debug().Model(&entity.FarmerProfile{}).Create(&entity.FarmerProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})
			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		case entity.Processor:
			res := tx.Debug().Model(&entity.ProcessorProfile{}).Create(&entity.ProcessorProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})
			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}

			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		case entity.Collector:
			res := tx.Debug().Model(&entity.CollectorProfile{}).Create(&entity.CollectorProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})
			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		case entity.Distributor:
			res := tx.Debug().Model(&entity.DistributorProfile{}).Create(&entity.DistributorProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})

			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}

			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		case entity.Seller:
			res := tx.Debug().Model(&entity.SellerProfile{}).Create(&entity.ConsumerProfile{
				UserId: input.UserId,
				Name:   input.NewRole.Name,
			})

			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}

			if res.RowsAffected == 0 {
				tx.Rollback()
				return gorm.ErrRecordNotFound
			}
		default:
			tx.Rollback()
			return helper.ErrRoleNotFound
		}

		res2 := tx.Debug().Model(&entity.User{}).Where("id = ?", input.UserId).Update("role", input.NewRole.Role)
		if res2.Error != nil {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}

		if res2.RowsAffected == 0 {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}

		res3 := tx.Debug().Model(&entity.Token{}).Where("user_id = ?", input.UserId).Delete(entity.Token{})
		if res3.Error != nil {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}

	} else {
		tx.Rollback()
		return helper.ErrQuantityNotEnough
	}

	key := fmt.Sprintf("user:%d", input.UserId)
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return tx.Commit().Error

}

func (r *userRepository) Me(ctx context.Context, id uint) (*dto.GetUser, error) {
	var result dto.GetUser

	var user entity.User
	var dataProfile dto.DataProfile
	key := fmt.Sprintf("user:%d", id)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil && val != "" {
		fmt.Println("im redis")
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return &result, nil
		}
	}
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	res := tx.Debug().Model(&entity.User{}).Select("id", "email", "is_verified", "role", "provider").Where("id = ?", id).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res.Error
	}

	switch user.Role {
	case entity.Consumer:
		if err := tx.Debug().Model(&entity.ConsumerProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Distributor:
		if err := tx.Debug().Model(&entity.DistributorProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Seller:
		if err := tx.Debug().Model(&entity.SellerProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Farmer:
		if err := tx.Debug().Model(&entity.FarmerProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

	case entity.Collector:
		if err := tx.Debug().Model(&entity.CollectorProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	case entity.Processor:
		if err := tx.Debug().Model(&entity.ProcessorProfile{}).Select("id", "name AS name").Where("user_id = ?", id).First(&dataProfile).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	default:
		tx.Rollback()
		return nil, helper.ErrRoleNotFound
	}

	result = dto.GetUser{
		Id:         user.ID,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		Role:       user.Role,
		Provider:   user.Provider,
	}
	result.Data = dataProfile

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
