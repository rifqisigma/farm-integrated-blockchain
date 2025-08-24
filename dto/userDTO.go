package dto

//create
type CreateConsumerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type CreateFarmerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type CreateDistributorProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type CreateRetailerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

//update
type UpdateConsumerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type UpdateFarmerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type UpdateDistributorProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type UpdateRetailerProfile struct {
	UserId uint   `json:"-"`
	Name   string `json:"name" validate:"required"`
}

type UserChangePassword struct {
	Email              string `json:"-" validate:"required"`
	Token              string `json:"-" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}
