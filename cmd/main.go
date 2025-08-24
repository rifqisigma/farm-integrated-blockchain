package main

import (
	"farm-integrated-web3/cmd/database"
	"farm-integrated-web3/cmd/route"
	"farm-integrated-web3/internal/handler"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env : %v", err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	//auth
	authRepo := repository.NewAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUC)

	//user
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	//middleware
	middlewareAuth := middleware.NewAuthMiddleware(userRepo)

	route := route.NewRoutes(authHandler, userHandler, middlewareAuth)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, route))

}
