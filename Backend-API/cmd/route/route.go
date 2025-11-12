package route

import (
	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/utils/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRoutes(
	auth *handler.AuthHandler,
	user *handler.UserHandler,
	farmer *handler.Farmerhandler,
	distributor *handler.DistributorHandler,
	seller *handler.SellerHandler,
	collector *handler.CollectorHandler,
	processor *handler.ProcessorHandler,
	blockchain *handler.BlockchainHandler,
	middlewareAuth *middleware.AuthMiddleware) *mux.Router {

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
	authPath.HandleFunc("/resend-verification", auth.ResendVerificationEmail).Methods(http.MethodPost)

	//auth secure
	authPathSecure := r.PathPrefix("/auth").Subrouter()
	authPathSecure.Use(middlewareAuth.Auth)
	authPathSecure.HandleFunc("/delete-account", auth.DeleteAccount).Methods(http.MethodDelete)
	authPathSecure.HandleFunc("/logout", auth.Logout).Methods(http.MethodPost)

	authPathRefresh := r.PathPrefix("/auth").Subrouter()
	authPathRefresh.Use(middlewareAuth.RefreshTokenMiddleware)
	authPathRefresh.HandleFunc("/refresh-token", auth.RefreshLongToken).Methods(http.MethodPost)

	authPathRefresh.HandleFunc("/access-token", auth.CreateAccessToken).Methods(http.MethodPost)

	// User routes
	userPath := r.PathPrefix("/user").Subrouter()
	userPath.Use(middlewareAuth.Auth)

	//create profile
	r.HandleFunc("/user/{id}/profile", user.CreateProfile).Methods(http.MethodPost)

	//update profile
	userPath.HandleFunc("/profile", user.UpdateProfile).Methods(http.MethodPatch)
	userPath.HandleFunc("/role", user.UpdateRole).Methods(http.MethodPatch)
	userPath.HandleFunc("/change-password", user.ChangePassword).Methods(http.MethodPost)
	userPath.HandleFunc("/me", user.Me).Methods(http.MethodGet)

	//farm
	farmerPath := r.PathPrefix("/farm").Subrouter()
	farmerPath.Use(middlewareAuth.Auth)

	//w
	farmerPath.HandleFunc("/harvest", farmer.CreateHarvest).Methods(http.MethodPost)
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.UpdateHarvest).Methods(http.MethodPatch)
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.DeleteHarvest).Methods(http.MethodDelete)
	farmerPath.HandleFunc("/status/{harvest}", farmer.UpdateStatusHarvest).Methods(http.MethodPatch)

	//accept
	farmerPath.HandleFunc("/collector/{collector}", farmer.AcceptHarvestCollector).Methods(http.MethodPatch)
	farmerPath.HandleFunc("/processor/{processor}", farmer.AcceptHarvestProcessor).Methods(http.MethodPatch)
	farmerPath.HandleFunc("/distribution/{distribution}", farmer.AcceptDistribution).Methods(http.MethodPatch)

	//get
	farmerPath.HandleFunc("/harvest/{harvest}", farmer.HarvestById).Methods(http.MethodGet)
	farmerPath.HandleFunc("/search", farmer.SearchHarvest).Methods(http.MethodGet)
	farmerPath.HandleFunc("/fyp", farmer.ListHarvestFYP).Methods(http.MethodGet)
	farmerPath.HandleFunc("", farmer.ListHarvestByFarmerId).Methods(http.MethodGet)

	//collector
	collectorPath := r.PathPrefix("/collector").Subrouter()
	collectorPath.Use(middlewareAuth.Auth)

	//w
	collectorPath.HandleFunc("/harvest/{harvest}", collector.CreateHarvestCollector).Methods(http.MethodPost)
	collectorPath.HandleFunc("/{collector}", collector.UpdateHarvestCollector).Methods(http.MethodPatch)
	collectorPath.HandleFunc("/{collector}", collector.CreateHarvestCollector).Methods(http.MethodDelete)

	//accept
	collectorPath.HandleFunc("/accept/distribution/{distribution}", collector.AcceptDistributor).Methods(http.MethodPatch)
	collectorPath.HandleFunc("/accept/processor/{processor}", collector.AcceptHarvestProcessor).Methods(http.MethodPatch)

	//get
	collectorPath.HandleFunc("/fyp", collector.ListHarvestCollectorFYP).Methods(http.MethodGet)
	collectorPath.HandleFunc("/search", collector.SearchHarvestCollector).Methods(http.MethodGet)
	collectorPath.HandleFunc("/id/{collector}", collector.GetHarvestCollectorById).Methods(http.MethodGet)
	collectorPath.HandleFunc("", collector.ListHarvestCollectorByCollectorId).Methods(http.MethodGet)

	//processor
	processorPath := r.PathPrefix("/processor").Subrouter()
	processorPath.Use(middlewareAuth.Auth)

	//w
	processorPath.HandleFunc("", processor.CreateProcessor).Methods(http.MethodPost)
	processorPath.HandleFunc("/{processor}", processor.UpdateProcessor).Methods(http.MethodPatch)
	processorPath.HandleFunc("/{processor}", processor.DeleteProcessor).Methods(http.MethodDelete)

	//accept
	processorPath.HandleFunc("/accept/distribution/{distribution}", processor.AcceptDistributor).Methods(http.MethodPatch)

	//get
	processorPath.HandleFunc("/fyp", processor.ListHarvestProcessorFYP).Methods(http.MethodGet)
	processorPath.HandleFunc("/search", processor.SearchHarvestProcessor).Methods(http.MethodGet)
	processorPath.HandleFunc("/id/{processor}", processor.GetHarvestProcessorById).Methods(http.MethodGet)
	processorPath.HandleFunc("", processor.ListHarvestProcessorByProcessorId).Methods(http.MethodGet)

	//distributor
	distributionPath := r.PathPrefix("/distribution").Subrouter()
	distributionPath.Use(middlewareAuth.Auth)

	distributionPath.HandleFunc("", distributor.CreateDistribution).Methods(http.MethodPost)
	distributionPath.HandleFunc("/{distribution}", distributor.UpdateDistribution).Methods(http.MethodPatch)
	distributionPath.HandleFunc("/{distribution}/status", distributor.UpdateStatusDistribution).Methods(http.MethodPatch)
	distributionPath.HandleFunc("/{distribution}", distributor.DeleteDistribution).Methods(http.MethodDelete)
	distributionPath.HandleFunc("/{distribution}/status", distributor.UpdateStatusDistribution).Methods(http.MethodPatch)

	//accept
	distributionPath.HandleFunc("/accept/seller/{sellerbox}", distributor.AcceptSeller).Methods(http.MethodPatch)

	//get
	distributionPath.HandleFunc("/search", distributor.SearchDistributions).Methods(http.MethodGet)
	distributionPath.HandleFunc("/fyp", distributor.GetListDistributionFYP).Methods(http.MethodGet)
	distributionPath.HandleFunc("/id/{distribution}", distributor.GetDistributionById).Methods(http.MethodGet)
	distributionPath.HandleFunc("", distributor.GetListDistributionsByDistributorId).Methods(http.MethodGet)

	distributionPath.HandleFunc("/fyp", distributor.GetListDistributionFYP).Methods(http.MethodGet)

	//seller
	sellerPath := r.PathPrefix("/seller").Subrouter()
	sellerPath.Use(middlewareAuth.Auth)
	sellerPath.HandleFunc("/distribution/{distribution}", seller.AddSellerBox).Methods(http.MethodPost)
	sellerPath.HandleFunc("/{seller}", seller.UpdateSellerBox).Methods(http.MethodPatch)
	sellerPath.HandleFunc("/{seller}", seller.DeleteSellerBox).Methods(http.MethodDelete)

	//get
	sellerPath.HandleFunc("/search", seller.SearchSellerBox).Methods(http.MethodGet)
	sellerPath.HandleFunc("/id/{seller}", seller.GetSellerBoxById).Methods(http.MethodGet)
	sellerPath.HandleFunc("/fyp", seller.ListGetSellerBoxsbySellerIdFYP).Methods(http.MethodGet)
	sellerPath.HandleFunc("", seller.ListGetSellerBoxsbySellerId).Methods(http.MethodGet)

	//blockchain
	blockchainPath := r.PathPrefix("/blockchain").Subrouter()
	blockchainPath.Use(middlewareAuth.Auth)
	blockchainPath.HandleFunc("/harvest/{harvest}", blockchain.GetHarvestBcByID).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/harvest-collector/{collector}", blockchain.GetHarvestCollectorBcByID).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/harvest-processor/{processor}", blockchain.GetHarvestProcessorBcByID).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/distribution/{distribution}", blockchain.GetDistributionByID).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/seller-box/{seller-box}", blockchain.GetSellerBoxBcByID).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/harvest", blockchain.GetAllHarvestBc).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/harvest-collector", blockchain.GetAllHarvestCollectorBc).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/harvest-processor", blockchain.GetAllHarvestProcessorBc).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/distribution", blockchain.GetAllDsitributionBc).Methods(http.MethodGet)
	blockchainPath.HandleFunc("/seller-box", blockchain.GetAllSellerBoxBc).Methods(http.MethodGet)

	return r
}
