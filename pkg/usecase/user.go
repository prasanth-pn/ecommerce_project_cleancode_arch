package usecase

import (
	//domain "clean/pkg/domain"
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"errors"
	//"context"
	//"fmt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
func (c *userUseCase) ListProducts(pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := c.userRepo.ListProducts(pagenation)

	return product, metadata, err
}
func (c *userUseCase) ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error) {
	catego, metadata, err := c.userRepo.ListCategories(pagenation)
	return catego, metadata, err
}
func (c *userUseCase) ListProductsByCategories(category_id int, pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error) {
	products, metadata, err := c.userRepo.ListProductsByCategories(category_id, pagenation)
	return products, metadata, err

}

func (c *userUseCase) Count_WishListed_Product(user_id, product_id uint) int {
	count := c.userRepo.Count_WishListed_Product(user_id, product_id)
	return count

}
func (c *userUseCase) AddTo_WishList(wishlist domain.WishList) error {

	err := c.userRepo.AddTo_WishList(wishlist)
	return err

}

func (c *userUseCase) ViewWishList(user_id uint) []domain.WishListResponse {

	wishlist := c.userRepo.ViewWishList(user_id)
	return wishlist

}
func (c *userUseCase) RemoveFromWishlist(user_id, product_id int) error {
	err := c.userRepo.RemoveFromWishlist(user_id, product_id)
	return err
}
func (c *userUseCase) FindCart(user_id, product_id uint) (domain.CartResponse, error) {
	cart, error := c.userRepo.FindCart(user_id, product_id)
	return cart, error
}
func (c *userUseCase) DeleteCart(product_id, user_id int) error {
	err := c.userRepo.DeleteCart(product_id, user_id)
	return err
}
func (c userUseCase) FindAddress(user_id, address_id uint) (domain.Address, error) {
	address, err := c.userRepo.FindAddress(user_id, address_id)
	return address, err

}
func (c *userUseCase) UpdateAddress(add domain.Address, user_id, Address_id uint) error {

	err := c.userRepo.UpdateAddress(add, user_id, Address_id)

	return err
}

func (c *userUseCase) FindTheSumOfCart(user_id int) (int, error) {
	sum, err := c.userRepo.FindTheSumOfCart(user_id)
	return sum, err

}
func (c *userUseCase) CreateOrder(order domain.Orders) error {

	err := c.userRepo.CreateOrder(order)
	return err
}
func (c *userUseCase) SearchOrder(order_id string) (domain.Orders, error) {
	order, err := c.userRepo.SearchOrder(order_id)
	return order, err
}
func (c *userUseCase) UpdateOrders(payment_id, order_id string) error {
	err := c.userRepo.UpdateOrders(payment_id, order_id)
	return err

}
func (c *userUseCase) Insert_To_My_Order(cart domain.CartListResponse, order_id string) error {
	err := c.userRepo.Insert_To_My_Order(cart, order_id)
	return err
}
func (c *userUseCase) ClearCart(user_id uint) error {
	err := c.userRepo.ClearCart(user_id)
	return err
}
func (c *userUseCase) ListOrder(pagenation utils.Filter, user_id uint) ([]domain.OrderResponse, utils.Metadata, error) {
	order, metadata, err := c.userRepo.ListOrder(pagenation, user_id)

	return order, metadata, err
}
func (c *userUseCase) FindCoupon(coupon string) (domain.Coupon, error) {
	Coupon, err := c.userRepo.FindCoupon(coupon)
	return Coupon, err
}
func (c *userUseCase) UpdateUser(user domain.Users) (domain.Users, error) {
	User, err := c.userRepo.UpdateUser(user)
	return User, err
}
func (c *userUseCase) UpdateProductQuantity(product_id, quantity int) error {
	product, err := c.FindProduct(uint(product_id))
	if err != nil {
		return errors.New("error finding  product")
	}
	totalQuantity := product.Quantity - uint16(quantity)
	err = c.userRepo.UpdateProductQuantity(product_id, int(totalQuantity))
	return err
}
