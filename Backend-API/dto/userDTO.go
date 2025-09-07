package dto

import "farm-integrated-web3/entity"

//create
type CreateProfileRequest struct {
	UserId uint          `json:"-"`
	Name   string        `json:"name" validate:"required"`
	Role   entity.Status `json:"role" validate:"required,oneof=consumer farmer distributor retailer"`
}

//update
type UpdateProfileRequest struct {
	UserId uint          `json:"-"`
	Name   string        `json:"name" validate:"required"`
	Role   entity.Status `json:"role" validate:"required,oneof=consumer farmer distributor retailer"`
}

type UserChangePasswordRequest struct {
	Email              string `json:"-" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}

type UpdateRoleRequest struct {
	UserId  uint          `json:"-"`
	OldRole entity.Status `json:"-" validate:"required,oneof=farmer distributor retailer"`
	NewRole NewRole       `json:"new_role"`
}

type NewRole struct {
	Role entity.Status `json:"role"`
	Name string        `json:"name"`
}
