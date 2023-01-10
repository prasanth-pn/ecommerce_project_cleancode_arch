package interfaces

import (
	"clean/pkg/domain"
	"context"
)


type AdminUseCase interface{
	ListUsers()([]domain.UserResponse,error)
	AddProducts( ctx context.Context,product domain.Product)(error)
	AddCategory(ctx context.Context,category domain.Category)(error)


}