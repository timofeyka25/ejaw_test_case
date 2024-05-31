package handler

import (
	"ejaw_test_case/internal/domain"
	"ejaw_test_case/pkg/utils"
	"encoding/json"
	"net/http"
)

type UserService interface {
	CreateUser(username, password, role string) error
	AuthenticateUser(username, password string) (*domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type userDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user userDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.userService.CreateUser(user.Username, user.Password, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials userDTO
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "authentication failed", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":    user.ID,
		"token": token,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
