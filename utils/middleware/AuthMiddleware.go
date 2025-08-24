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
	userRepo repository.UserRepository
}

func NewAuthMiddleware(userRepo repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepo: userRepo}
}

func (a *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.HttpError(w, http.StatusUnauthorized, "Unauthorized: No token provided")
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helper.ParseJWT(tokenString)
		if err != nil || !claims.Verified {
			helper.HttpError(w, http.StatusForbidden, "Unauthorized: Invalid or unverified user")
			return
		}

		if err := a.userRepo.ValidateToken(claims.UserID); err != nil {
			helper.HttpError(w, http.StatusUnauthorized, "Unauthorized: Token not valid")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
