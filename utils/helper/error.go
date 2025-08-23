package helper

import "errors"

var (
	ErrLoginNotSuccess = errors.New("email or password is wrong")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidToken    = errors.New("token invliad or expired")
	ErrBadRequest      = errors.New("something wrong with your request")
)
