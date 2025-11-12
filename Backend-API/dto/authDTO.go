package dto

import "farm-integrated-web3/entity"

// gmail traditional
type LoginRequest struct {
	Email    string `json:"email"  validate:"required,email" example:"rajaSukasari@gmail.com"`
	Password string `json:"password"  validate:"required" example:"12345678"`
}

type LoginResponse struct {
	Id           uint          `json:"id" gorm:"column:id"`
	ProfileId    uint          `json:"profile_id" gorm:"column:profile_id"`
	Email        string        `json:"email" gorm:"column:email"`
	Role         entity.Status `json:"role" gorm:"column:role"`
	PasswordHash string        `json:"password"  gorm:"column:password"`
	IsVerified   bool          `json:"is_verified" gorm:"column:is_verified"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"rajaSukasari@gmail.com"`
	Password string `json:"password"  validate:"required" example:"1234567"`
	Name     string `json:"name" validate:"required" example:"name"`
}

type RegisterResponse struct {
	Id       uint   `gorm:"column:id"`
	Email    string `gorm:"column:id"`
	Provider string `gorm:"provider"`
	Verified bool   `gorm:"is_verified"`
}

type UserResetPasswordRequest struct {
	Email              string `json:"-" validate:"required"`
	Token              string `json:"-" validate:"required"`
	UserId             uint   `json:"-" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required" example:"7654321"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required" example:"7654321"`
}
