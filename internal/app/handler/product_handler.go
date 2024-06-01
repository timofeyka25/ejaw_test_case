package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ejaw_test_case/internal/domain"
)

type ProductService interface {
	CreateProduct(name, description string, price float64, sellerID int) (*domain.Product, error)
	GetProduct(id int) (*domain.Product, error)
	UpdateProduct(id int, name, description string, price float64, sellerID int) (*domain.Product, error)
	DeleteProduct(id int) error
	GetAllProducts() ([]domain.Product, error)
}

type ProductHandler struct {
	productService ProductService
}

func NewProductHandler(productService ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

type productDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerID    int     `json:"seller_id"`
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDTO
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productService.CreateProduct(product.Name, product.Description, product.Price, product.SellerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(productToProductDTO(createdProduct)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	product, err := h.productService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(productToProductDTO(product)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	var product productDTO
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	product.ID = id

	updatedProduct, err := h.productService.UpdateProduct(product.ID, product.Name, product.Description, product.Price, product.SellerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(productToProductDTO(updatedProduct)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	err = h.productService.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(productsResponseDTO(products)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
