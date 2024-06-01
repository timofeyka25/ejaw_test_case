package service

import (
	"ejaw_test_case/internal/domain"
	"ejaw_test_case/internal/service/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewProductService(mockRepo)

	testProduct := &domain.Product{
		Name:        "New Product",
		Description: "A new product description",
		Price:       19.99,
		SellerID:    1,
	}

	mockRepo.EXPECT().
		CreateProduct(gomock.Any()).
		Return(nil).
		Times(1)

	resultProduct, err := service.CreateProduct("New Product", "A new product description", 19.99, 1)

	assert.NoError(t, err)
	assert.Equal(t, testProduct, resultProduct)
}
