package main

import (
	"farm-integrated-web3/cmd/database"
	"farm-integrated-web3/cmd/route"
	_ "farm-integrated-web3/docs"

	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env : %v", err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	// auth
	authRepo := repository.NewAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUC)

	// user
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	// farmer
	farmerRepo := repository.NewFarmerRepository(db)
	farmerUC := usecase.NewFarmerUsecase(farmerRepo)
	farmerHandler := handler.NewFarmerHandler(farmerUC)

	// distributor
	distributorRepo := repository.NewDistributorRepository(db)
	distributorUC := usecase.NewDistributorUsecase(distributorRepo)
	distributorHandler := handler.NewDistributorHandler(distributorUC)

	// retailer
	retailerRepo := repository.NewRetailerRepository(db)
	retailerUC := usecase.NewRetailerUsecase(retailerRepo)
	retailerHandler := handler.NewRetailerHandler(retailerUC)

	// middleware
	middlewareAuth := middleware.NewAuthMiddleware(authRepo)

	// routes
	r := route.NewRoutes(authHandler, userHandler, farmerHandler, distributorHandler, retailerHandler, middlewareAuth)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
