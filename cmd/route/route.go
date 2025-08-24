package route

import (
	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/utils/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRoutes(auth *handler.AuthHandler, user *handler.UserHandler, middlewareAuth *middleware.AuthMiddleware) *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	authPath := r.PathPrefix("/gmail").Subrouter()

	authPath.HandleFunc("/register", auth.Register).Methods(http.MethodPost)
	authPath.HandleFunc("/login", auth.Login).Methods(http.MethodPost)
	authPath.HandleFunc("/forgot-password", auth.ResetPassword).Methods(http.MethodPost)
	authPath.HandleFunc("/reset-password", auth.ResetPassword).Methods(http.MethodPost)
	authPath.HandleFunc("/verification", auth.ValidateUser).Methods(http.MethodGet)
	authPath.HandleFunc("/refresh-token", auth.RefreshLongToken).Methods(http.MethodPost)
	authPath.HandleFunc("/resend-verification", auth.ResendVerificationEmail).Methods(http.MethodPost)

	// User routes
	userPath := r.PathPrefix("/user").Subrouter()
	userPath.Use(middlewareAuth.Auth)

	//create profile
	r.HandleFunc("/user/consumer", user.CreateConsumer).Methods(http.MethodPost)
	r.HandleFunc("/user/farmer", user.CreateFarmer).Methods(http.MethodPost)
	r.HandleFunc("/user/distributor", user.CreateDistributor).Methods(http.MethodPost)
	r.HandleFunc("/user/retailer", user.CreateRetailer).Methods(http.MethodPost)

	//update profile
	userPath.HandleFunc("/consumer", user.UpdateConsumerProfile).Methods(http.MethodPut)
	userPath.HandleFunc("/farmer", user.UpdateFarmerProfile).Methods(http.MethodPut)
	userPath.HandleFunc("/distributor", user.UpdateDistributorProfile).Methods(http.MethodPut)
	userPath.HandleFunc("/retailer", user.UpdateRetailerProfile).Methods(http.MethodPut)
	userPath.HandleFunc("/change-password", user.ChangePassword).Methods(http.MethodPost)
	userPath.HandleFunc("/delete-account", user.DeleteAccount).Methods(http.MethodDelete)
	userPath.HandleFunc("/logout", user.Logout).Methods(http.MethodPost)

	return r
}
