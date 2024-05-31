package service

import (
	"ejaw_test_case/internal/domain"
	"fmt"
)

type SellerRepository interface {
	AddSeller(seller domain.Seller) error
	GetSeller(id int) (*domain.Seller, error)
	GetSellerByName(name string) (*domain.Seller, error)
	UpdateSeller(seller domain.Seller) error
	DeleteSeller(id int) error
	GetAllSellers() ([]domain.Seller, error)
}

type SellerService struct {
	repo SellerRepository
}

func NewSellerService(repo SellerRepository) *SellerService {
	return &SellerService{repo: repo}
}

func (s *SellerService) AddSeller(name, phone string) (*domain.Seller, error) {
	seller := &domain.Seller{
		Name:  name,
		Phone: phone,
	}

	err := s.repo.AddSeller(*seller)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (s *SellerService) GetSeller(id int) (*domain.Seller, error) {
	return s.repo.GetSeller(id)
}

func (s *SellerService) GetSellerByName(name string) (*domain.Seller, error) {
	return s.repo.GetSellerByName(name)
}

func (s *SellerService) UpdateSeller(id int, name, phone string) (*domain.Seller, error) {
	seller, err := s.repo.GetSeller(id)
	if err != nil {
		return nil, err
	}
	if seller == nil {
		return nil, fmt.Errorf("seller not found")
	}

	seller.Name = name
	seller.Phone = phone

	err = s.repo.UpdateSeller(*seller)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (s *SellerService) DeleteSeller(id int) error {
	return s.repo.DeleteSeller(id)
}

func (s *SellerService) GetAllSellers() ([]domain.Seller, error) {
	return s.repo.GetAllSellers()
}
