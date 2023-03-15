package usecase

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
)

func (c *userUseCase) FindProduct(product_id uint) (domain.Product, error) {
	//var Products domain.Product
	Products, err := c.userRepo.FindProduct(product_id)
	return Products, err
}
func (c *userUseCase) ListCart(pagenation utils.Filter, User_ID uint) ([]domain.CartListResponse, utils.Metadata, error) {
	cart, metadata, err := c.userRepo.ListCart(pagenation, User_ID)
	return cart, metadata, err
}
func (c *userUseCase) QuantityCart(product_id, user_id uint) (domain.Cart, error) {
	cart, err := c.userRepo.QuantityCart(product_id, user_id)
	return cart, err

}
func (c *userUseCase) UpdateCart(totalprice float32, quantity, product_id, user_id uint) (domain.Cart, error) {
	cart, err := c.userRepo.UpdateCart(totalprice, quantity, product_id, user_id)
	return cart, err
}

func (c *userUseCase) CreateCart(cart domain.Cart) (domain.Cart, error) {
	Cart, err := c.userRepo.CreateCart(cart)
	return Cart, err

}

func (c *userUseCase) ListViewCart(user_id uint) ([]domain.CartListResponse, error) {
	cart, err := c.userRepo.ListViewCart(user_id)
	return cart, err
}
func (c *userUseCase) TotalCartPrice(user_id uint) (float32, error) {
	sum, err := c.userRepo.TotalCartPrice(user_id)
	return sum, err
}
