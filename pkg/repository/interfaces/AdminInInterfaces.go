package interfaces

import "clean/pkg/domain"

type AdminRepository interface{
	ListUsers()([]domain.UserResponse,error)
	AddProducts(product domain.Product)(error)
	AddCategory(category domain.Category)(error)

}