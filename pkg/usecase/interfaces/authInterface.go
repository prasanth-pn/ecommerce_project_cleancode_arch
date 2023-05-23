package interfaces

import (
	domain "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
)

type AuthUseCase interface {
	Register(user domain.Users) (domain.Users, error)
	AdminRegister(admin domain.Admins) (domain.Admins, error)
	VerifyUser(email, password string) error
	FindUser(email string) (domain.UserResponse, error)
	VerifyAdmin(username, password string) error
	FindAdmin(username string) (*domain.AdminResponse, error)
	SendVerificationEmail(email string) error
	VerifyUserOtp(email string, code int) error
	UpdateUserStatus(email string) error
	FindUserById(user_id uint) (domain.Users, error)
	BlockUnblockUser(user_id uint, val bool) error
}
