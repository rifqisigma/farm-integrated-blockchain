package main

import (
	"farm-integrated-web3/cmd/database"
	"farm-integrated-web3/cmd/route"
	_ "farm-integrated-web3/docs"
	"os/signal"
	"syscall"
	"time"

	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/internal/worker"
	"farm-integrated-web3/utils/middleware"
	"log"
	"net/http"
	"os"
)

// @title Agrichain - Distribution API
// @version 1.0
// @description This is a Backend for Agricultural of Distribution for reach transparant, decentralization, and immutable.
// @termsOfService http://swagger.io/terms/
// @contact.name Project Menager
// @contact.email ipb_rifqi@apps.ipb.ac.id
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	db, rdb, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	// auth
	authRepo := repository.NewAuthRepository(db, rdb)
	authUC := usecase.NewAuthUsecase(authRepo, rdb)
	authHandler := handler.NewAuthHandler(authUC)

	// user
	userRepo := repository.NewUserRepository(db, rdb, authRepo)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	// farmer
	farmerRepo := repository.NewFarmerRepository(db, rdb)
	farmerUC := usecase.NewFarmerUsecase(farmerRepo, rdb)
	farmerHandler := handler.NewFarmerHandler(farmerUC)

	//collector
	collectorRepo := repository.NewCollectorRepository(db, rdb)
	collectorUC := usecase.NewCollectorUsecase(collectorRepo, rdb)
	collectorHandler := handler.NewCollectorHandler(collectorUC)

	//processor
	processorRepo := repository.NewProcessorRepository(db, rdb)
	processorUC := usecase.NewProcessoUsecase(processorRepo, rdb)
	processorHandler := handler.NewProcessorHandler(processorUC)

	// distributor
	distributorRepo := repository.NewDistributorRepository(db, rdb)
	distributorUC := usecase.NewDistributorUsecase(distributorRepo, rdb)
	distributorHandler := handler.NewDistributorHandler(distributorUC)

	// seller
	sellerRepo := repository.NewSellerRepository(db, rdb)
	sellerUC := usecase.NewSellerUsecase(sellerRepo, rdb)
	SellerHandler := handler.NewSellerHandler(sellerUC)

	// blockchain
	blockchainUC := usecase.NewBlockchainRepository()
	blockchainHandler := handler.NewBlockchainHandler(blockchainUC)

	// middleware
	middlewareAuth := middleware.NewAuthMiddleware(authRepo, rdb)

	// routes
	r := route.NewRoutes(
		authHandler,
		userHandler,
		farmerHandler,
		distributorHandler,
		SellerHandler,
		collectorHandler,
		processorHandler,
		blockchainHandler,
		middlewareAuth,
	)

	w := worker.NewWorker(farmerRepo, distributorRepo, sellerRepo, collectorRepo, processorRepo, rdb)
	w.StartFlushWorker(1 * time.Second)
	log.Println("Worker started...")
	defer w.StopFlushWorker()

	// jalankan server di background
	go func() {
		port := os.Getenv("PORT")
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()

	// tunggu signal untuk stop
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Println("Stopping worker...")

}
