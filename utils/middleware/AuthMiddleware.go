package middleware

import (
	"context"
	"farm-integrated-web3/utils/helper"
	"net/http"
	"strings"
)

type key int

const UserContextKey key = 1

func AuthMiddleware(next http.Handler) http.Handler {
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

		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
