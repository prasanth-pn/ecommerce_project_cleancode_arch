package interfaces

import (
	"clean/pkg/domain"
	"clean/pkg/utils"
)

type AdminRepository interface{
	ListUsers(pagenation utils.Filter)([]domain.UserResponse,utils.Metadata,error)
	ListBlockedUsers(pagenation utils.Filter)([]domain.Users,utils.Metadata,error)
	AddProducts(product domain.Product)(error)
	AddCategory(category domain.Category)(error)
	AddBrand(brand domain.Brand)error
	AddModel(model domain.Model)error

}