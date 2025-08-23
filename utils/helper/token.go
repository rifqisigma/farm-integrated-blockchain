package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// hash pw
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func ValidateToken(hashedToken, rawToken string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(rawToken)); err != nil {
		return false, errors.New("invalid refresh token")
	}
	return true, nil
}
