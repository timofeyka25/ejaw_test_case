package main

import (
	"ejaw_test_case/internal/app/handler"
	"ejaw_test_case/internal/repository"
	"ejaw_test_case/internal/service"
	"ejaw_test_case/pkg/config"
	"ejaw_test_case/pkg/database"
	"ejaw_test_case/pkg/server"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalln(err)
	}

	err = database.InitDB(db)
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	sellerRepo := repository.NewSellerRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	sellerService := service.NewSellerService(sellerRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	sellerHandler := handler.NewSellerHandler(sellerService)

	routes := handler.InitRoutes(userHandler, productHandler, sellerHandler)

	server.StartServerWithGracefulShutdown(routes, config.Get().ServerURL)
}
