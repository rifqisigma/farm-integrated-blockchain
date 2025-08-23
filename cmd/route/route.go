package route

import (
	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/utils/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRoutes(auth *handler.AuthHandler, user *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	authPath := r.PathPrefix("/gmail").Subrouter()

	authPath.HandleFunc("/register", auth.Register).Methods(http.MethodPost)
	authPath.HandleFunc("/login", auth.Login).Methods(http.MethodPost)
	authPath.HandleFunc("/forgot-password", auth.ResetPassword).Methods(http.MethodGet)
	authPath.HandleFunc("/reset-password", auth.ResetPassword).Methods(http.MethodPost)
	authPath.HandleFunc("/verification", auth.ValidateUser).Methods(http.MethodGet)
	authPath.HandleFunc("/refresh-token", auth.RefreshLongToken).Methods(http.MethodGet)

	// User routes
	userPath := r.PathPrefix("/user").Subrouter()
	userPath.Use(middleware.AuthMiddleware)

	return r
}
