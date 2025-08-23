package dto

// gmail traditional
type Login struct {
	Email    string `json:"email"  validate:"required,  email"`
	Password string `json:"password"  validate:"required"`
}
type LoginResponse struct {
	Id           uint
	Email        string
	Role         string
	PasswordHash string
	IsVerified   bool
}

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required"`
	Provider string `json:"-" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	Id       uint
	Email    string
	Provider string
	Verified bool
}

type UserNewPassword struct {
	Email              string `json:"-" validate:"required"`
	Token              string `json:"-" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}
