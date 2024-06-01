package handler

import "ejaw_test_case/internal/domain"

func productToProductDTO(product *domain.Product) productDTO {
	return productDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		SellerID:    product.SellerID,
	}
}

func productsResponseDTO(products []domain.Product) []productDTO {
	var productsDTO []productDTO
	for _, product := range products {
		productsDTO = append(productsDTO, productToProductDTO(&product))
	}
	return productsDTO
}

func sellerToSellerDTO(seller *domain.Seller) sellerDTO {
	return sellerDTO{
		ID:    seller.ID,
		Name:  seller.Name,
		Phone: seller.Phone,
	}
}

func sellersResponseDTO(sellers []domain.Seller) []sellerDTO {
	var sellersDTO []sellerDTO
	for _, seller := range sellers {
		sellersDTO = append(sellersDTO, sellerToSellerDTO(&seller))
	}
	return sellersDTO
}
