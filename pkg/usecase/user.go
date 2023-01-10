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

// func (c *userUseCase)FindAll(ctx context.Context)([]domain.Users,error){
// 	users,err:=c.userRepo.FindAll(ctx)
// 	return users,err
// }
