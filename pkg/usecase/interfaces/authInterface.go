package interfaces


import(
	domain"clean/pkg/domain"
	"context"
)

type AuthUseCase interface{
	Register(ctx context.Context,user domain.Users)(domain.Users,error)
	AdminRegister(ctx context.Context,admin domain.Admins)(domain.Admins,error)
	// FindAll(ctx context.Context)([]domain.Users,error)
	VerifyUser(email,password string )(error)
	FindUser(email string)(domain.UserResponse,error)
	VerifyAdmin(username,password string)(error)
	FindAdmin(username string)(*domain.AdminResponse,error)
	//VerifyPassword(password,userpassword string)(error)


}