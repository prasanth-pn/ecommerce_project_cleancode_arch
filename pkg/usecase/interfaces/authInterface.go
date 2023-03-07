package interfaces

import (
	domain "clean/pkg/domain"
	"context"
)

type AuthUseCase interface {
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	AdminRegister(ctx context.Context, admin domain.Admins) (domain.Admins, error)
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
