package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
)

type UserUseCase interface {
	ListProducts(pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error)
	ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	ListProductsByCategories(category_id int, pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error)
	FindProduct(product_id uint) (domain.Product, error)
	ListCart(pagention utils.Filter, User_id uint) ([]domain.CartListResponse, utils.Metadata, error)
	ListViewCart(User_id uint) ([]domain.CartListResponse, error)
	FindCart(user_id, product_id uint) (domain.CartResponse, error)
	QuantityCart(product_id, user_id uint) (domain.Cart, error)
	UpdateCart(totalprice float32, quantity, product_id, user_id uint) (domain.Cart, error)
	CreateCart(cart domain.Cart) (domain.Cart, error)
	FindTheSumOfCart(user_id int) (int, error)
	TotalCartPrice(user_id uint) (float32, error)
	AddAddress(address domain.Address) error
	ListAddress(user_id uint) ([]domain.Address, error)
	Count_WishListed_Product(user_id, product_id uint) int
	AddTo_WishList(domain.WishList) error
	ViewWishList(user_id uint) []domain.WishListResponse
	RemoveFromWishlist(user_id, product_id int)(error)
	FindAddress(user_id, address_id uint) (domain.Address, error)
	UpdateAddress(add domain.Address, user_id, address_id uint) error
	CreateOrder(order domain.Orders) error
	SearchOrder(order_id string) (domain.Orders, error)
	UpdateOrders(payment_id, order_id string) error
	Insert_To_My_Order(carts domain.CartListResponse, order_id string) error
	ClearCart(user_id uint) error
	ListOrder(user_id uint) ([]domain.ListOrder, uint, error)
	FindCoupon(coupon string) (domain.Coupon, error)
	UpdateUser(user domain.Users)(domain.Users,error)
}
