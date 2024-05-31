package service

import (
	"ejaw_test_case/internal/domain"
	"fmt"
)

type ProductRepository interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id int) (*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id int) error
	GetAllProducts() ([]domain.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name, description string, price float64, sellerID int) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
		SellerID:    sellerID,
	}

	err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProduct(id int) (*domain.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) UpdateProduct(id int, name, description string, price float64, sellerID int) (*domain.Product, error) {
	product, err := s.repo.GetProduct(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.SellerID = sellerID

	err = s.repo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.repo.GetAllProducts()
}
