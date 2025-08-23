package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTclaims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type JWTclaimsLongExp struct {
	UserID   uint `json:"user_id"`
	Verified bool `json:"verified"`
	jwt.RegisteredClaims
}

type JWTclaimsShortExp struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("SECRET_KEY")

func GenerateJWT(email, role string, id uint, verified bool) (string, error) {
	claims := JWTclaims{
		UserID:   id,
		Email:    email,
		Verified: verified,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenstring string) (*JWTclaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &JWTclaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTclaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func GenerateJWTLongExp(id uint, verified bool) (string, error) {
	claims := JWTclaimsLongExp{
		UserID:   id,
		Verified: verified,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * 24 * 7 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWTLongExp(tokenstring string) (*JWTclaimsLongExp, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &JWTclaimsLongExp{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTclaimsLongExp)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func GenerateJWTShortExp(email string) (string, error) {
	claims := JWTclaimsShortExp{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWTShortExp(tokenstring string) (*JWTclaimsShortExp, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &JWTclaimsShortExp{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTclaimsShortExp)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
