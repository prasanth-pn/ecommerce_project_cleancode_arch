package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
)

type AdminRepository interface {
	ListUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error)
	ListBlockedUsers(pagenation utils.Filter) ([]domain.Users, utils.Metadata, error)
	AddProducts(product domain.Product) (int, error)
	DeleteProduct(product_id int) error
	UpdateProduct(product domain.Product)error
	FindProduct(product_id int) (domain.ProductResponse, error)
	ListProductByCategories(pagenation utils.Filter,cat_id int)([]domain.ProductResponse,utils.Metadata,error)
	AddCategory(category domain.Category) error
	ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	AddBrand(brand domain.Brand) error
	AddModel(model domain.Model) error
	ImageUpload(image []string, product_id int) error
}
