package middleware

import (
	"context"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"
	"net/http"
	"strings"
)

type key int

const UserContextKey key = 1

type AuthMiddleware struct {
	authRepo repository.AuthRepository
}

func NewAuthMiddleware(authRepo repository.AuthRepository) *AuthMiddleware {
	return &AuthMiddleware{authRepo}
}

func (a *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.HttpError(w, http.StatusUnauthorized, "No token provided")
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helper.ParseJWT(tokenString)
		if err != nil || !claims.Verified {
			helper.HttpError(w, http.StatusForbidden, "Invalid or unverified user")
			return
		}

		if err := a.authRepo.ValidateToken(tokenString); err != nil {
			helper.HttpError(w, http.StatusUnauthorized, "Token not expired")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func (a *AuthMiddleware) RefreshTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.HttpError(w, http.StatusUnauthorized, "No token provided")
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helper.ParseJWTLongExp(tokenString)
		if err != nil || !claims.Verified {
			helper.HttpError(w, http.StatusForbidden, "Invalid or unverified user")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
