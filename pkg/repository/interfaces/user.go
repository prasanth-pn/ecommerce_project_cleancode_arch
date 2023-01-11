package interfaces

import "clean/pkg/domain"

type UserRepository interface {
	ListProducts() ([]domain.ProductResponse, error)
	FindProduct(product_id uint) (domain.Product, error)
	ListCart(User_id uint) ([]domain.Cart, error)
	QuantityCart(product_id, user_id uint) (domain.Cart, error)
	UpdateCart(totalprice float32,quantity,product_id,user_id uint)error
	CreateCart(cart domain.Cart)error
	ViewCart(user_id uint)([]domain.CartListResponse,error)
	TotalCartPrice(user_id uint)(float32,error)
	
}
