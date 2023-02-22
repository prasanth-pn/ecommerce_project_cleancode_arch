package usecase

import (
	//domain "clean/pkg/domain"
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
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
func (c *userUseCase) ListProducts() ([]domain.ProductResponse, error) {
	product, err := c.userRepo.ListProducts()

	return product, err
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
func (c *userUseCase)RemoveFromWishlist(user_id,product_id int){
	c.userRepo.RemoveFromWishlist(user_id,product_id)
}
