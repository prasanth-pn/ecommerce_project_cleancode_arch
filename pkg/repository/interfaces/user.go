package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
)

type UserRepository interface {
	ListProducts(pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error)
	ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	ListProductsByCategories(category_id int, pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error)
	FindProduct(product_id uint) (domain.Product, error)
	ListCart(User_id uint) ([]domain.Cart, error)
	FindCart(user_id, product_id uint) (domain.CartResponse, error)
	QuantityCart(product_id, user_id uint) (domain.Cart, error)
	UpdateCart(totalprice float32, quantity, product_id, user_id uint) error
	CreateCart(cart domain.Cart) error
	ViewCart(user_id uint) ([]domain.CartListResponse, error)
	FindTheSumOfCart(user_id int) (int, error)
	TotalCartPrice(user_id uint) (float32, error)
	AddAddress(address domain.Address) error
	ListAddress(user_id uint) ([]domain.Address, error)
	Count_WishListed_Product(user_id, product_id uint) int
	AddTo_WishList(wishlist domain.WishList) error
	ViewWishList(user_id uint) []domain.WishListResponse
	RemoveFromWishlist(user_id, product_id int)
	FindAddress(user_id, address_id uint) (domain.Address, error)
	UpdateAddress(add domain.Address, user_id, address_id uint) error
	CreateOrder(order domain.Orders) error
	SearchOrder(order_id string)(domain.Orders,error)
	UpdateOrders(payment_id ,order_id string)error
	Insert_To_My_Order(carts domain.CartListResponse ,order_id string)error
	ListOrder(user_id uint)([]domain.ListOrder, uint,error)
	ClearCart(user_id uint)error
	
}
