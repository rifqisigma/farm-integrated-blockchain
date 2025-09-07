package route

import (
	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/utils/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRoutes(auth *handler.AuthHandler, user *handler.UserHandler, farmer *handler.Farmerhandler, distributor *handler.DistributorHandler, retailer *handler.RetailerHandler, middlewareAuth *middleware.AuthMiddleware) *mux.Router {

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth routes
	authPathParent := r.PathPrefix("/auth").Subrouter()
	authPath := authPathParent.PathPrefix("/gmail").Subrouter()
	authPath.HandleFunc("/register", auth.Register).Methods(http.MethodPost)
	authPath.HandleFunc("/login", auth.Login).Methods(http.MethodPost)
	authPath.HandleFunc("/forgot-password", auth.RequestResetPassword).Methods(http.MethodPost)
	authPath.HandleFunc("/reset-password", auth.ResetPassword).Methods(http.MethodPost)
	authPath.HandleFunc("/verification", auth.ValidateUser).Methods(http.MethodGet)
	authPath.HandleFunc("/refresh-token", auth.RefreshLongToken).Methods(http.MethodPost)
	authPath.HandleFunc("/resend-verification", auth.ResendVerificationEmail).Methods(http.MethodPost)

	//auth secure
	authPathSecure := r.PathPrefix("/auth").Subrouter()
	authPathSecure.Use(middlewareAuth.Auth)
	authPathSecure.HandleFunc("/delete-account", auth.DeleteAccount).Methods(http.MethodDelete)
	authPathSecure.HandleFunc("/logout", auth.Logout).Methods(http.MethodPost)
	authPathSecure.HandleFunc("/access-token", auth.CreateAccessToken).Methods(http.MethodPost)

	// User routes
	userPath := r.PathPrefix("/user").Subrouter()
	userPath.Use(middlewareAuth.Auth)

	//create profile
	r.HandleFunc("/user/{id}/profile", user.CreateProfile).Methods(http.MethodPost)

	//update profile
	userPath.HandleFunc("/profile", user.UpdateProfile).Methods(http.MethodPatch)
	userPath.HandleFunc("/role", user.UpdateRole).Methods(http.MethodPost)
	userPath.HandleFunc("/change-password", user.ChangePassword).Methods(http.MethodPost)

	//farm
	farmerPath := r.PathPrefix("/farm").Subrouter()
	farmerPath.Use(middlewareAuth.Auth)

	farmerPath.HandleFunc("/harvest/crop/{crop}", farmer.CreateHarvest).Methods(http.MethodPost)
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.UpdateHarvest).Methods(http.MethodPatch)
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.DeleteHarvest).Methods(http.MethodDelete)
	farmerPath.HandleFunc("/distribution/{distribution}", farmer.AcceptedFarmerForDistributor).Methods(http.MethodPatch)
	farmerPath.HandleFunc("/harvest", farmer.ListHarvestByFarmerId).Methods(http.MethodGet)
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.HarvestById).Methods(http.MethodGet)
	farmerPath.HandleFunc("/harvest/search", farmer.SearchHarvest).Methods(http.MethodGet)

	//distributor
	distributionPath := r.PathPrefix("/distribution").Subrouter()
	distributionPath.Use(middlewareAuth.Auth)

	distributionPath.HandleFunc("/harvest/{harvest}", distributor.CreateDistribution).Methods(http.MethodPost)
	distributionPath.HandleFunc("/{distribution}", distributor.UpdateDistribution).Methods(http.MethodPatch)
	distributionPath.HandleFunc("/{distribution}/status", distributor.UpdateStatusDistribution).Methods(http.MethodPatch)
	distributionPath.HandleFunc("/{distribution}", distributor.DeleteDistribution).Methods(http.MethodDelete)
	distributionPath.HandleFunc("/search", distributor.SearchDistributions).Methods(http.MethodGet)
	distributionPath.HandleFunc("", distributor.GetDistributionsByDistributorId).Methods(http.MethodGet)
	distributionPath.HandleFunc("/{distribution}", distributor.GetDistributionById).Methods(http.MethodGet)
	distributionPath.HandleFunc("/{distribution}/status", distributor.UpdateStatusDistribution).Methods(http.MethodPatch)
	distributionPath.HandleFunc("/retailer-cart/{retailerCart", distributor.ApprovedRetailerCartForRetailer).Methods(http.MethodPatch)

	//retailer
	retailerPath := r.PathPrefix("/retail").Subrouter()
	retailerPath.Use(middlewareAuth.Auth)

	retailerPath.HandleFunc("/distribution/{distribution}", retailer.AddRetailerCart).Methods(http.MethodPost)
	retailerPath.HandleFunc("/{retailer}/distribution/{distribution}", retailer.UpdateRetailerCart).Methods(http.MethodPatch)
	retailerPath.HandleFunc("/{retailer}", retailer.DeleteRetailerCart).Methods(http.MethodDelete)
	retailerPath.HandleFunc("/search", retailer.SearchRetailerCart).Methods(http.MethodGet)
	retailerPath.HandleFunc("", retailer.GetRetailerCarts).Methods(http.MethodGet)
	retailerPath.HandleFunc("/{retailer}", retailer.GetRetailerCartById).Methods(http.MethodGet)

	return r
}
