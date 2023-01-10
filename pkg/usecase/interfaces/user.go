package interfaces

import "clean/pkg/domain"

type UserUseCase interface {
	ListProducts() ([]domain.ProductResponse, error)
	FindProduct(product_id uint) (domain.Product, error)
	ListCart(User_id uint) ([]domain.Cart, error)
}
