package service

import (
	"ejaw_test_case/internal/domain"
	"ejaw_test_case/pkg/utils"
	"fmt"
)

type UserRepository interface {
	AddUser(user domain.User) error
	GetUser(id int) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, password, role string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	user := domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	err = s.repo.AddUser(user)
	if err != nil {
		return fmt.Errorf("error adding user: %v", err)
	}

	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (*domain.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("incorrect password")
	}

	return user, nil
}
