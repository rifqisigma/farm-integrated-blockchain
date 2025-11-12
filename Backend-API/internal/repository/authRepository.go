package repository

import (
	"context"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository interface {

	//gmail traditional
	Login(ctx context.Context, input *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, input *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ValidateUser(ctx context.Context, email string) (bool, error)
	ChangePassword(ctx context.Context, email, password string) error
	CheckUserExist(ctx context.Context, email string) (bool, error)
	CreateToken(ctx context.Context, userId uint, token string) error
	GetUserInfo(ctx context.Context, id uint) (*dto.LoginResponse, error)
	ValidateToken(ctx context.Context, userId uint, token string) (bool, error)
	UpdateRevokeToken(ctx context.Context, userId uint) error
	DeleteAccount(ctx context.Context, userId uint, role string) error
}
type authRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthRepository(db *gorm.DB, redis *redis.Client) AuthRepository {
	return &authRepository{db, redis}
}

func (r *authRepository) Register(ctx context.Context, input *dto.RegisterRequest) (*dto.RegisterResponse, error) {

	newUser := entity.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     entity.None,
		Provider: "gmail",
	}

	if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, gorm.ErrDuplicatedKey
		}
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

func (r *authRepository) Login(ctx context.Context, input *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user dto.LoginResponse

	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	res := tx.Debug().Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("email = ? AND is_verified = ? and role IS NOT NULL", input.Email, true).Scan(&user)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, helper.ErrLoginNotSuccess
	}

	var entityDestination interface{}

	if user.Role != "" {
		switch user.Role {
		case entity.Consumer:
			entityDestination = entity.ConsumerProfile{}
		case entity.Distributor:
			entityDestination = entity.DistributorProfile{}
		case entity.Farmer:
			entityDestination = entity.FarmerProfile{}
		case entity.Seller:
			entityDestination = entity.SellerProfile{}
		case entity.Processor:
			entityDestination = entity.ProcessorProfile{}
		case entity.Collector:
			entityDestination = entity.CollectorProfile{}
		default:
			return nil, helper.ErrRoleNotFound
		}
	}

	res2 := tx.Debug().Model(&entityDestination).Where("user_id = ?", user.Id).Pluck("id", &user.ProfileId)
	if res2.Error != nil {
		return nil, res2.Error
	}

	keys, _ := r.redis.Keys(ctx, "user:%d:token:*").Result()
	if len(keys) > 0 {
		r.redis.Del(ctx, keys...)
	}

	return &user, tx.Commit().Error
}

func (r *authRepository) ValidateUser(ctx context.Context, email string) (bool, error) {
	res := r.db.WithContext(ctx).Debug().Model(&entity.User{}).Where("email = ? AND is_verified = ?", email, false).Update("is_verified", true)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {

		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *authRepository) ChangePassword(ctx context.Context, email, password string) error {
	res := r.db.WithContext(ctx).Debug().Model(&entity.User{}).Where("email = ?", email).Update("password", password)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {

		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *authRepository) CreateToken(ctx context.Context, userId uint, token string) error {
	tx := r.db.Debug().WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := r.UpdateRevokeToken(ctx, userId); err != nil {
		tx.Rollback()
		return err
	}
	newToken := entity.Token{
		UserId:    userId,
		Token:     token,
		IsRevoked: false,
	}

	if err := tx.Model(&entity.Token{}).Create(&newToken).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *authRepository) GetUserInfo(ctx context.Context, id uint) (*dto.LoginResponse, error) {
	tx := r.db.Begin().WithContext(ctx)
	var user dto.LoginResponse
	res := tx.Debug().Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("id = ? AND is_verified = ?", id, true).Scan(&user)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	var model interface{}
	switch user.Role {
	case entity.Consumer:
		model = entity.ConsumerProfile{}
	case entity.Collector:
		model = entity.CollectorProfile{}
	case entity.Processor:
		model = entity.ProcessorProfile{}
	case entity.Farmer:
		model = entity.FarmerProfile{}
	case entity.Distributor:
		model = entity.DistributorProfile{}
	case entity.Seller:
		model = entity.SellerProfile{}
	default:
		return nil, gorm.ErrRecordNotFound
	}

	res2 := tx.Debug().Model(&model).Where("user_id = ?", user.Id).Pluck("id", &user.ProfileId)
	if res2.Error != nil {
		tx.Rollback()
		return nil, res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (r *authRepository) UpdateRevokeToken(ctx context.Context, userId uint) error {
	res := r.db.Debug().Model(&entity.Token{}).
		Where("user_id = ?", userId).Update("is_revoked", true)
	if res.Error != nil {
		return res.Error
	}

	keys, _ := r.redis.Keys(ctx, "user:%d:token:*").Result()
	if len(keys) > 0 {
		r.redis.Del(ctx, keys...)
	}
	return nil

}

func (r *authRepository) CheckUserExist(ctx context.Context, email string) (bool, error) {
	var count int64
	res := r.db.WithContext(ctx).Debug().Model(&entity.User{}).Where("email = ?", email).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return count > 0, nil

}

func (r *authRepository) ValidateToken(ctx context.Context, userId uint, token string) (bool, error) {
	key := fmt.Sprintf("user:%d:token:%s", userId, token)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil && val == "true" {
		return false, nil
	} else if err == nil && val == "false" {
		return false, errors.New("invalid token")
	}

	var count int64
	if err := r.db.WithContext(ctx).Debug().Model(&entity.Token{}).Where("token = ? AND is_revoked = ?", token, false).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}

	if count <= 0 {
		return false, gorm.ErrRecordNotFound
	}

	err = r.redis.Set(ctx, key, count > 0, 5*time.Second).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *authRepository) DeleteAccount(ctx context.Context, userId uint, role string) error {
	res := r.db.WithContext(ctx).Debug().Model(&entity.User{}).Where("id = ?", userId).Update("is_canceled", true)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	switch entity.Status(role) {
	case entity.Consumer:
		res := r.db.WithContext(ctx).Debug().Model(&entity.ConsumerProfile{}).Where("user_id = ?", userId).Update("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

	case entity.Distributor:
		res := r.db.WithContext(ctx).Debug().Model(&entity.DistributorProfile{}).Where("user_id = ?", userId).Update("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	case entity.Farmer:
		res := r.db.WithContext(ctx).Debug().Model(&entity.FarmerProfile{}).Where("user_id = ?", userId).Update("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	case entity.Seller:
		res := r.db.WithContext(ctx).Debug().Model(&entity.SellerProfile{}).Where("user_id = ?", userId).Update("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	default:
		return gorm.ErrRecordNotFound
	}

	keyToken := fmt.Sprintf("user:%d:token:*", userId)
	keyUser := fmt.Sprintf("user:%d", userId)
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.Del(ctx, keyUser).Err(); err != nil {
			return err
		}
		if err := p.Del(ctx, keyToken).Err(); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
