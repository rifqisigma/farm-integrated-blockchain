package repository

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {

	//gmail traditional
	Login(input *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(input *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ValidateUser(email string) (bool, error)
	ChangePassword(email, password string) error
	CheckUserExist(userId uint, email string) (bool, error)
	CreateToken(userId uint, token string) error
	GetUserInfo(id uint) (*dto.LoginResponse, error)
	ValidateToken(token string) error
	UpdateRevokeToken(userId uint) error
	DeleteAccount(userId uint) error
}
type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(input *dto.RegisterRequest) (*dto.RegisterResponse, error) {

	newUser := entity.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     entity.Consumer,
		Provider: "gmail",
	}

	if err := r.db.Create(&newUser).Error; err != nil {
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

func (r *authRepository) Login(input *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user dto.LoginResponse

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("email = ? AND is_verified = ?", input.Email, true).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	switch user.Role {
	case entity.Consumer:
		if err := tx.Model(&entity.ConsumerProfile{}).Select("id").Where("user_id = ?", user.Id).Limit(1).Scan(&user.ProfileId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		}
	case entity.Distributor:
		if err := tx.Model(&entity.DistributorProfile{}).Select("id").Where("user_id = ?", user.Id).Limit(1).Scan(&user.ProfileId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		}
	case entity.Farmer:
		if err := tx.Model(&entity.FarmerProfile{}).Select("id").Where("user_id = ?", user.Id).Limit(1).Scan(&user.ProfileId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		}
	case entity.Retailer:
		if err := tx.Model(&entity.FarmerProfile{}).Select("id").Where("user_id = ?", user.Id).Limit(1).Scan(&user.ProfileId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return nil, gorm.ErrRecordNotFound
			}
		}
	default:
		return nil, gorm.ErrRecordNotFound
	}
	return &user, tx.Commit().Error
}

func (r *authRepository) ValidateUser(email string) (bool, error) {
	if err := r.db.Model(&entity.User{}).Where("email = ? AND is_verified = ?", email, false).UpdateColumn("is_verified", true).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}

		return false, err
	}

	return true, nil
}

func (r *authRepository) ChangePassword(email, password string) error {
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *authRepository) CreateToken(userId uint, token string) error {
	newToken := entity.Token{
		UserID: userId,
		Token:  token,
	}

	if err := r.db.Create(&newToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetUserInfo(id uint) (*dto.LoginResponse, error) {
	var user dto.LoginResponse
	if err := r.db.Model(&entity.User{}).Select("id", "email", "role", "is_verified", "password").Where("id = ? AND is_verified = ?", id, true).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) UpdateRevokeToken(userId uint) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ?", userId).UpdateColumn("is_revoked", true).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return gorm.ErrRecordNotFound
		}

		return err
	}

	return nil

}

func (r *authRepository) CheckUserExist(userId uint, email string) (bool, error) {
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

func (r *authRepository) ValidateToken(token string) error {
	if err := r.db.Model(&entity.Token{}).Where("token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *authRepository) DeleteAccount(userId uint) error {
	tx := r.db.Commit()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(&entity.User{}).Where("id = ?", userId).UpdateColumn("is_canceled", true).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}
