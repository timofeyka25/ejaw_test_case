package handler

import (
	"ejaw_test_case/internal/app/middleware"
	"net/http"
)

func InitRoutes(
	userHandler *UserHandler,
	productHandler *ProductHandler,
	sellerHandler *SellerHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/signup", userHandler.SignUp)
	mux.HandleFunc("POST /auth/signin", userHandler.SignIn)

	sellersMux := http.NewServeMux()
	sellersMux.HandleFunc("POST /sellers/", sellerHandler.AddSeller)
	sellersMux.HandleFunc("GET /sellers/{id}", sellerHandler.GetSeller)
	sellersMux.HandleFunc("PUT /sellers/{id}", sellerHandler.UpdateSeller)
	sellersMux.HandleFunc("DELETE /sellers/{id}", sellerHandler.DeleteSeller)
	sellersMux.HandleFunc("GET /sellers/all", sellerHandler.GetAllSellers)

	productsMux := http.NewServeMux()
	productsMux.HandleFunc("POST /products/", productHandler.CreateProduct)
	productsMux.HandleFunc("GET /products/{id}", productHandler.GetProduct)
	productsMux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	productsMux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)
	productsMux.HandleFunc("GET /products/all", productHandler.GetAllProducts)

	mux.Handle("/sellers/", middleware.AuthMiddleware(sellersMux))
	mux.Handle("/products/", middleware.AuthMiddleware(productsMux))

	return mux
}
