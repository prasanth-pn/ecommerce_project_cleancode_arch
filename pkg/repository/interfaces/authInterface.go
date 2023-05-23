package interfaces

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
)

type AuthRepository interface {
	Register(user domain.Users) (int, error)
	AdminRegister(admin domain.Admins) error
	FindAdmin(username string) (domain.AdminResponse, error)
	// FindAll(ctx context.Context)([]domain.Users,error)
	FindUser(email string) (domain.UserResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyOtp(email string, code int) error
	UpdateUserStatus(email string) error
	FindUserById(user_id uint) (domain.Users, error)
	BlockUnblockUser(user_id uint, val bool) error
}
