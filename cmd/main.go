package main

import (
	"ejaw_test_case/internal/app/handler"
	"ejaw_test_case/internal/repository"
	"ejaw_test_case/internal/service"
	"ejaw_test_case/pkg/config"
	"ejaw_test_case/pkg/database"
	"ejaw_test_case/pkg/utils"
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

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	routes := handler.InitRoutes(userHandler, productHandler)

	utils.StartServerWithGracefulShutdown(routes, config.Get().ServerURL)
}
