package usecase

import (
	"clean/pkg/domain"
	"fmt"
)

func (c *userUseCase) FindProduct(product_id uint) (domain.Product, error) {
	//var Products domain.Product
	Products, err := c.userRepo.FindProduct(product_id)
	return Products, err
}
func (c *userUseCase) ListCart(User_ID uint) ([]domain.Cart, error) {
	var cart []domain.Cart
	cart, _ = c.userRepo.ListCart(User_ID)
	fmt.Println(cart)
	return cart, nil

}
func (c *userUseCase) QuantityCart(product_id, user_id uint) (domain.Cart, error) {
	cart, err := c.userRepo.QuantityCart(product_id, user_id)
	return cart, err

}
func (c *userUseCase) UpdateCart(totalprice float32, quantity, product_id, user_id uint) error {
	err := c.userRepo.UpdateCart(totalprice, quantity, product_id, user_id)
	return err
}

func (c *userUseCase) CreateCart(cart domain.Cart) error {
	err := c.userRepo.CreateCart(cart)
	return err

}
