package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
	"context"
)

type AdminUseCase interface {
	ListUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ListBlockedUsers(pagenation utils.Filter) (*[]domain.Users, *utils.Metadata, error)
	AddProducts(ctx context.Context, product domain.Product) (int, error)
	FindProduct(product_id int) (domain.ProductResponse, error)
	ListProductByCategories(pagenation utils.Filter, cate_id int) ([]domain.ProductResponse, utils.Metadata, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(product_id int) error
	AddCategory(category domain.Category) error
	ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	AddBrand(ctx context.Context, brand domain.Brand) error
	AddModel(ctx context.Context, model domain.Model) error
	ImageUpload(image []string, product_id int) error
}
