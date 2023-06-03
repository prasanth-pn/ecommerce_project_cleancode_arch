package interfaces

import (
	"context"

	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"
)

type AdminUseCase interface {
	SearchUserByName(pagenation utils.Filter, name string) ([]domain.Users, utils.Metadata, error)
	ListUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ListBlockedUsers(pagenation utils.Filter) (*[]domain.Users, *utils.Metadata, error)
	AddProducts(product domain.Product) (int, error)
	FindProduct(product_id int) (domain.ProductResponse, error)
	ListProductByCategories(pagenation utils.Filter, cate_id int) ([]domain.ProductResponse, utils.Metadata, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(product_id int) error
	AddCategory(category domain.Category) error
	ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	AddBrand(ctx context.Context, brand domain.Brand) error
	AddModel(ctx context.Context, model domain.Model) error
	ImageUpload(image []string, product_id int) error
	DeleteImage(product_id int, imagename string) error
	GenerateCoupon(coupon domain.Coupon) error
	FindCoupon(coupon string) (domain.Coupon, error)
	ListOrder(pagention utils.Filter) ([]domain.Orders, utils.Metadata, error)
    
}
