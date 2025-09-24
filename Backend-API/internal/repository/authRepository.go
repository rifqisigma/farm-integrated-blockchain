package repository

import (
	"context"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"fmt"

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
	ValidateToken(ctx context.Context, userId uint, token string) error
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
		Role:     entity.Consumer,
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

	res := tx.Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("email = ? AND is_verified = ?", input.Email, true).Scan(&user)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	if user.ProfileId != 0 {
		switch user.Role {
		case entity.Consumer:
			res := tx.Model(&entity.ConsumerProfile{}).Select("id").Where("user_id = ?", user.Id).Scan(&user.ProfileId)
			if res.Error != nil {
				return nil, res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		case entity.Distributor:
			res := tx.Model(&entity.DistributorProfile{}).Select("id").Where("user_id = ?", user.Id).Scan(&user.ProfileId)
			if res.Error != nil {
				return nil, res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		case entity.Farmer:
			res := tx.Model(&entity.FarmerProfile{}).Select("id").Where("user_id = ?", user.Id).Scan(&user.ProfileId)
			if res.Error != nil {
				return nil, res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		case entity.Retailer:
			res := tx.Model(&entity.FarmerProfile{}).Select("id").Where("user_id = ?", user.Id).Scan(&user.ProfileId)
			if res.Error != nil {
				return nil, res.Error
			}
			if res.RowsAffected == 0 {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		default:
			return nil, gorm.ErrRecordNotFound
		}
	} else {
		return nil, fmt.Errorf("user %d not havent profile", user.ProfileId)
	}

	keys, _ := r.redis.Keys(ctx, "user:%d:token:*").Result()
	if len(keys) > 0 {
		r.redis.Del(ctx, keys...)
	}

	return &user, tx.Commit().Error
}

func (r *authRepository) ValidateUser(ctx context.Context, email string) (bool, error) {
	res := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ? AND is_verified = ?", email, false).UpdateColumn("is_verified", true)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {

		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}

func (r *authRepository) ChangePassword(ctx context.Context, email, password string) error {
	res := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Update("password", password)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {

		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *authRepository) CreateToken(ctx context.Context, userId uint, token string) error {
	newToken := entity.Token{
		UserID: userId,
		Token:  token,
	}

	if err := r.db.WithContext(ctx).Create(&newToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetUserInfo(ctx context.Context, id uint) (*dto.LoginResponse, error) {
	var user dto.LoginResponse
	res := r.db.WithContext(ctx).Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("id = ? AND is_verified = ?", id, true).Scan(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (r *authRepository) UpdateRevokeToken(ctx context.Context, userId uint) error {
	res := r.db.Model(&entity.Token{}).
		Where("user_id = ?", userId).UpdateColumn("is_revoked", true)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keys, _ := r.redis.Keys(ctx, "user:%d:token:*").Result()
	if len(keys) > 0 {
		r.redis.Del(ctx, keys...)
	}
	return nil

}

func (r *authRepository) CheckUserExist(ctx context.Context, email string) (bool, error) {
	res := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil

}

func (r *authRepository) ValidateToken(ctx context.Context, userId uint, token string) error {
	key := fmt.Sprintf("user:%d:token:%s", userId, token)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil && val == "true" {
		return nil
	} else if err == nil && val == "false" {
		return errors.New("invalid token")
	}

	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.Token{}).Where("token = ? AND is_validate = ?", token, true).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	valid := false
	if count > 0 {
		valid = true
	}

	err = r.redis.Set(ctx, key, valid, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) DeleteAccount(ctx context.Context, userId uint, role string) error {
	res := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userId).UpdateColumn("is_canceled", true)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	switch entity.Status(role) {
	case entity.Consumer:
		res := r.db.WithContext(ctx).Model(&entity.ConsumerProfile{}).Where("user_id = ?", userId).UpdateColumn("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

	case entity.Distributor:
		res := r.db.WithContext(ctx).Model(&entity.DistributorProfile{}).Where("user_id = ?", userId).UpdateColumn("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	case entity.Farmer:
		res := r.db.WithContext(ctx).Model(&entity.FarmerProfile{}).Where("user_id = ?", userId).UpdateColumn("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	case entity.Retailer:
		res := r.db.WithContext(ctx).Model(&entity.RetailerProfile{}).Where("user_id = ?", userId).UpdateColumn("is_deleted", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	default:
		return gorm.ErrRecordNotFound
	}

	keys, _ := r.redis.Keys(ctx, "user:%d:token:*").Result()
	if len(keys) > 0 {
		r.redis.Del(ctx, keys...)
	}

	return nil
}
