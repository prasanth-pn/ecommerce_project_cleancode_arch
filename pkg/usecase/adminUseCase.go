package usecase

import (
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"context"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: adminRepo,
	}
}
func (c *adminUseCase) ListUsers() ([]domain.UserResponse, error) {
	//var user domain.UserResponse
	user, err := c.adminRepo.ListUsers()

	return user, err

}
func (c *adminUseCase) AddProducts(ctx context.Context, product domain.Product) error {
	err := c.adminRepo.AddProducts(product)
	return err
}

// ---------------------addCategory----------------------
func (c *adminUseCase) AddCategory(ctx context.Context, category domain.Category) error {
	err := c.adminRepo.AddCategory(category)
	return err
}

func (c *adminUseCase)AddBrand(ctx context.Context, brand domain.Brand)error  {

	err:=c.adminRepo.AddBrand(brand)
	return err
}
func ( c *adminUseCase)AddModel(ctx context.Context,model domain.Model)error{
	err:=c.adminRepo.AddModel(model)
	return err
}
