package handler

import (
	"ejaw_test_case/internal/app/middleware"
	"net/http"
)

func InitRoutes(userHandler *UserHandler, productHandler *ProductHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/signup", userHandler.SignUp)
	mux.HandleFunc("POST /auth/signin", userHandler.SignIn)

	productsMux := http.NewServeMux()
	productsMux.HandleFunc("POST /products/", productHandler.CreateProduct)
	productsMux.HandleFunc("GET /products/{id}", productHandler.GetProduct)
	productsMux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	productsMux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)
	productsMux.HandleFunc("GET /products/all", productHandler.GetAllProducts)

	mux.Handle("/products/", middleware.AuthMiddleware(productsMux))

	return mux
}
