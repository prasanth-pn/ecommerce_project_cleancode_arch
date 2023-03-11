package usecase

import (
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
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
func (c *adminUseCase) ListUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	//var user domain.UserResponse
	user, metadata, err := c.adminRepo.ListUsers(pagenation)
	//fmt.Println(&metadata, "metadata", &user)

	return &user, &metadata, err

}
func (c *adminUseCase) ListBlockedUsers(pagenation utils.Filter) (*[]domain.Users, *utils.Metadata, error) {
	user, metadata, err := c.adminRepo.ListBlockedUsers(pagenation)
	return &user, &metadata, err

}
func (c *adminUseCase) AddProducts(ctx context.Context, product domain.Product) (int, error) {
	product_id, err := c.adminRepo.AddProducts(product)
	return product_id, err
}
func (c *adminUseCase) FindProduct(product_id int) (domain.ProductResponse, error) {
	product, err := c.adminRepo.FindProduct(product_id)
	return product, err
}
func (c *adminUseCase) ListProductByCategories(pagenation utils.Filter, cate_id int) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := c.adminRepo.ListProductByCategories(pagenation, cate_id)
	return product, metadata, err
}
func (c *adminUseCase) DeleteProduct(product_id int) error {
	err := c.adminRepo.DeleteProduct(product_id)
	return err
}
func (c *adminUseCase) UpdateProduct(product domain.Product) error {
	err := c.adminRepo.UpdateProduct(product)
	return err
}

// ---------------------addCategory----------------------
func (c *adminUseCase) AddCategory(category domain.Category) error {
	err := c.adminRepo.AddCategory(category)
	return err
}
func (c *adminUseCase) ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error) {
	category, metadata, err := c.adminRepo.ListCategories(pagenation)
	return category, metadata, err
}

func (c *adminUseCase) AddBrand(ctx context.Context, brand domain.Brand) error {

	err := c.adminRepo.AddBrand(brand)
	return err
}
func (c *adminUseCase) AddModel(ctx context.Context, model domain.Model) error {
	err := c.adminRepo.AddModel(model)
	return err
}
func (c *adminUseCase) ImageUpload(image []string, product_id int) error {
	err := c.adminRepo.ImageUpload(image, product_id)
	return err
}
