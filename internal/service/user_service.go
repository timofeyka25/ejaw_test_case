package service

import (
	"ejaw_test_case/internal/domain"
	"ejaw_test_case/pkg/hash"
	"ejaw_test_case/pkg/jwt"
	"fmt"
)

//go:generate mockgen -destination=mocks/mock_userrepository.go -package=mocks . UserRepository

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
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user := domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	err = s.repo.AddUser(user)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}

	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (*domain.User, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, "", fmt.Errorf("error getting user: %w", err)
	}

	if user == nil {
		return nil, "", fmt.Errorf("user not found")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, "", fmt.Errorf("incorrect password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil
}
