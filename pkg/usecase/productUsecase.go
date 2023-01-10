package usecase

import (
	"clean/pkg/domain"
	"fmt"
)



func (c *userUseCase)FindProduct(product_id uint)(domain.Product,error){
	//var Products domain.Product
	Products,err:=c.userRepo.FindProduct(product_id)
	return Products,err
}
func (c *userUseCase)ListCart(User_ID uint)([]domain.Cart,error){
	var cart []domain.Cart
	cart,_=c.userRepo.ListCart(User_ID)
	fmt.Println(cart)
	return cart ,nil

}