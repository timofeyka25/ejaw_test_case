package repository

import (
	"database/sql"
	"ejaw_test_case/internal/domain"
	"errors"
	"fmt"
)

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) AddSeller(seller domain.Seller) error {
	query := "INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, seller.Name, seller.Phone).Scan(&seller.ID)
	if err != nil {
		return fmt.Errorf("error adding seller: %w", err)
	}
	return nil
}

func (r *SellerRepository) GetSeller(id int) (*domain.Seller, error) {
	query := "SELECT id, name, phone FROM sellers WHERE id = $1"
	seller := &domain.Seller{}
	err := r.db.QueryRow(query, id).Scan(&seller.ID, &seller.Name, &seller.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, fmt.Errorf("error getting seller: %w", err)
	}
	return seller, nil
}

func (r *SellerRepository) GetSellerByName(name string) (*domain.Seller, error) {
	query := "SELECT id, name, phone FROM sellers WHERE name = $1"
	seller := &domain.Seller{}
	err := r.db.QueryRow(query, name).Scan(&seller.ID, &seller.Name, &seller.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting seller by name: %w", err)
	}
	return seller, nil
}

func (r *SellerRepository) UpdateSeller(seller domain.Seller) error {
	query := "UPDATE sellers SET name = $1, phone = $2 WHERE id = $3"
	result, err := r.db.Exec(query, seller.Name, seller.Phone, seller.ID)
	if err != nil {
		return fmt.Errorf("error updating seller: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected after update: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no seller found with id %d", seller.ID)
	}

	return nil
}

func (r *SellerRepository) DeleteSeller(id int) error {
	query := "DELETE FROM sellers WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting seller: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected after delete: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no seller found with id %d", id)
	}

	return nil
}

func (r *SellerRepository) GetAllSellers() ([]domain.Seller, error) {
	query := "SELECT id, name, phone FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting all sellers: %w", err)
	}
	defer rows.Close()

	var sellers []domain.Seller
	for rows.Next() {
		var seller domain.Seller
		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Phone); err != nil {
			return nil, fmt.Errorf("error scanning seller: %w", err)
		}
		sellers = append(sellers, seller)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return sellers, nil
}
