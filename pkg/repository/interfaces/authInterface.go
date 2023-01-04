package interfaces


import(
	"context"
	"clean/pkg/domain"
)
type AuthRepository interface{
	Register(ctx context.Context,user domain.Users)(error)
	AdminRegister(ctx context.Context,admin domain.Admins)(error)
	FindAdmin(username string)(domain.AdminResponse,error)
	// FindAll(ctx context.Context)([]domain.Users,error)
     FindUser(email string)(domain.UserResponse,error)
}