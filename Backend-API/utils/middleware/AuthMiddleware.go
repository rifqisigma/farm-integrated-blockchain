package middleware

import (
	"context"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type key int

const UserContextKey key = 1

type AuthMiddleware struct {
	authRepo repository.AuthRepository
	redis    *redis.Client
}

func NewAuthMiddleware(authRepo repository.AuthRepository, redis *redis.Client) *AuthMiddleware {
	return &AuthMiddleware{authRepo, redis}
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

		if claims.ExpiresAt.Before(time.Now()) {
			helper.HttpError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userKey := fmt.Sprintf("rate_limit:%d", claims.UserID)
		limit := 10
		window := time.Minute

		pipe := a.redis.TxPipeline()
		cnt := pipe.Incr(r.Context(), userKey)
		pipe.Expire(r.Context(), userKey, window)
		_, err = pipe.Exec(r.Context())
		if err != nil {
			helper.HttpError(w, http.StatusInternalServerError, "Rate limiter error")
			return
		}

		count, _ := cnt.Result()
		if int(count) > limit {
			helper.HttpError(w, http.StatusTooManyRequests, "Rate limit exceeded")
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

		userKey := fmt.Sprintf("rate_limit:%d", claims.UserID)
		limit := 10
		window := time.Minute

		pipe := a.redis.TxPipeline()
		cnt := pipe.Incr(r.Context(), userKey)
		pipe.Expire(r.Context(), userKey, window)
		_, err = pipe.Exec(r.Context())
		if err != nil {
			helper.HttpError(w, http.StatusInternalServerError, "Rate limiter error")
			return
		}

		count, _ := cnt.Result()
		if int(count) > limit {
			helper.HttpError(w, http.StatusTooManyRequests, "Rate limit exceeded")
			return
		}

		if err := a.authRepo.ValidateToken(r.Context(), claims.UserID, tokenString); err != nil {
			helper.HttpError(w, http.StatusUnauthorized, "Token expired")
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
