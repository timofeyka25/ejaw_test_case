package repository

import (
	"database/sql"
	"ejaw_test_case/internal/domain"
	"errors"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	query := "INSERT INTO products (name, description, price, seller_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.SellerID).Scan(&product.ID)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}
	return nil
}

func (r *ProductRepository) GetProduct(id int) (*domain.Product, error) {
	query := "SELECT id, name, description, price, seller_id FROM products WHERE id = $1"
	product := &domain.Product{}
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SellerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("error getting product: %w", err)
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	query := "UPDATE products SET name = $1, description = $2, price = $3, seller_id = $4 WHERE id = $5"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.SellerID, product.ID)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}

func (r *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	query := "SELECT id, name, description, price, seller_id FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting all products: %w", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SellerID)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return products, nil
}
