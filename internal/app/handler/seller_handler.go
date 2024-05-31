package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ejaw_test_case/internal/domain"
)

type SellerService interface {
	AddSeller(name, phone string) (*domain.Seller, error)
	GetSeller(id int) (*domain.Seller, error)
	GetSellerByName(name string) (*domain.Seller, error)
	UpdateSeller(id int, name, phone string) (*domain.Seller, error)
	DeleteSeller(id int) error
	GetAllSellers() ([]domain.Seller, error)
}

type SellerHandler struct {
	sellerService SellerService
}

func NewSellerHandler(sellerService SellerService) *SellerHandler {
	return &SellerHandler{
		sellerService: sellerService,
	}
}

func (h *SellerHandler) AddSeller(w http.ResponseWriter, r *http.Request) {
	var seller domain.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	createdSeller, err := h.sellerService.AddSeller(seller.Name, seller.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSeller)
}

func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	seller, err := h.sellerService.GetSeller(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if seller == nil {
		http.Error(w, "seller not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(seller)
}

func (h *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	var seller domain.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	seller.ID = id

	updatedSeller, err := h.sellerService.UpdateSeller(seller.ID, seller.Name, seller.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSeller)
}

func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	err = h.sellerService.DeleteSeller(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SellerHandler) GetAllSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := h.sellerService.GetAllSellers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sellers)
}