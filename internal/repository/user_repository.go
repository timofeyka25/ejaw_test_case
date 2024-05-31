package repository

import (
	"database/sql"
	"ejaw_test_case/internal/domain"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) AddUser(user domain.User) error {
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUser(id int) (*domain.User, error) {
	query := "SELECT id, username, password, role FROM users WHERE id = $1"
	user := &domain.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	query := "SELECT id, username, password, role FROM users WHERE username = $1"
	user := &domain.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by username: %w", err)
	}
	return user, nil
}
