package dto

import "farm-integrated-web3/entity"

//create
type CreateProfileRequest struct {
	UserId uint          `json:"-"`
	Name   string        `json:"name" validate:"required"`
	Role   entity.Status `json:"role" validate:"required,oneof=consumer farmer collector processor distributor seller"`
}

//update
type UpdateProfileRequest struct {
	UserId uint          `json:"-"`
	Name   string        `json:"name" validate:"required"`
	Role   entity.Status `json:"-" validate:"required,oneof=consumer collector processor farmer distributor seller"`
}

type UserChangePasswordRequest struct {
	Email              string `json:"-" validate:"required"`
	UserId             uint   `json:"-" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}

type UpdateRoleRequest struct {
	UserId  uint          `json:"-"`
	OldRole entity.Status `json:"-" validate:"required,oneof=consumer farmer collector processor distributor seller"`
	NewRole NewRole       `json:"new_role"`
}

type NewRole struct {
	Role entity.Status `json:"role"`
	Name string        `json:"name"`
}

type GetUser struct {
	Id         uint          `json:"id" gorm:"column:id"`
	Email      string        `json:"email" gorm:"column:email"`
	IsVerified bool          `json:"is_verified" gorm:"column:is_verified"`
	Provider   string        `json:"provider" gorm:"column:provider"`
	Role       entity.Status `json:"role" gorm:"column:role"`
	Data       DataProfile   `json:"data"`
}

type DataProfile struct {
	Id   uint   `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}
