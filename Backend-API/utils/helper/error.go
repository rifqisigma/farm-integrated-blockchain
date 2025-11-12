package helper

import "errors"

var (
	ErrLoginNotSuccess = errors.New("email or password is wrong")
	ErrInvalidToken    = errors.New("token invliad or expired")
	ErrBadRequest      = errors.New("something wrong with your request")
	ErrDuplicateToken  = errors.New("token already exists")

	//farm
	ErrQuantityNotEnough = errors.New("quantity not enough")
	ErrRoleNotFound      = errors.New("role not found")
)
