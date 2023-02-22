package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
	"context"
)

type AdminUseCase interface {
	ListUsers(pagenation utils.Filter) (*[]domain.UserResponse,*utils.Metadata, error)
	AddProducts(ctx context.Context, product domain.Product) error
	AddCategory(ctx context.Context, category domain.Category) error
	AddBrand(ctx context.Context, brand domain.Brand) error
	AddModel(ctx context.Context,model domain.Model)error
}
