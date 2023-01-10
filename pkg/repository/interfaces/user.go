package interfaces

import "clean/pkg/domain"

// "context"
// "clean/pkg/domain"
type UserRepository interface{
	ListProducts()([]domain.ProductResponse,error)
	FindProduct(product_id uint)(domain.Product,error)
	ListCart(User_id uint)([]domain.Cart,error)
	//Register(ctx context.Context,user domain.Users)(domain.Users,error)
	// FindAll(ctx context.Context)([]domain.Users,error)
}
