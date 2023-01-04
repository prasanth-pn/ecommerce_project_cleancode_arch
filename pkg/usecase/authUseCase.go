package usecase

import (
	//"clean/pkg/common/response"

	domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"context"
	"errors"
	"fmt"
)

type authUseCase struct {
	authRepo interfaces.AuthRepository
}

func NewAuthUseCase(repo interfaces.AuthRepository) services.AuthUseCase {
	return &authUseCase{
		authRepo: repo,
	}
}

// ------------------------------------------register-------------------------
func (c *authUseCase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	_,err:=c.FindUser(user.Email)

	if err!=nil{
		return user,errors.New("user already exists")
		
	}


_,	err = c.authRepo.Register(ctx, user)
if err!=nil{
	return user,errors.New("failed to register")
}
	fmt.Println(user, "register")

	//fmt.Printf("\n\n %v ")
	return user, err
}

// -----------------------------------------verifyUser-----------------------------
func (c *authUseCase) VerifyUser(email, password string) error {
	user, _ := c.FindUser(email)
	if user.ID ==0 {
		return errors.New("username or password is incorrect")
	}
	IsValidPassword := VerifyPassword(password, user.Password)
	if !IsValidPassword {
		return errors.New(" not valid passwoerd")
	}
	return nil
}

// ---------------------------------find user-----------------
func (c *authUseCase) FindUser(email string) (domain.UserResponse, error) {
	var user domain.UserResponse
	user, err := c.authRepo.FindUser(email)
	if user.ID>0&&err == nil {
		return user, errors.New("user alresady exists")
	}
	return user, nil
}

// --------------------------veryfypassword-------------------------------------------------------------
func VerifyPassword(password, dbpassword string) bool {
	return password == dbpassword
}

// ------------------------------------------AdminRegister-----------------------------------
func (c *authUseCase) AdminRegister(ctx context.Context, admin domain.Admins) (domain.Admins, error) {
	c.authRepo.AdminRegister(ctx, admin)

	return admin, nil
}
func (c *authUseCase) VerifyAdmin(username, password string) error {
	//fmt.Println(username, "ifsdkohgndsiujkhgbdjkbhjfvbdhjffdgbauthcase")
	admin, err := c.authRepo.FindAdmin(username)
	if err != nil {
		return errors.New("username is incorrect")
	}
	fmt.Println("\n\n\n", admin.Password, password, "admin,password auth usecase")
	IsValidPassword := VerifyPassword(password, admin.Password)

	if !IsValidPassword {
		return errors.New("password is not valid ")
	}
	return nil
}
func (c *authUseCase) FindAdmin(username string) (*domain.AdminResponse, error) {
	user, err := c.authRepo.FindAdmin(username)
	return &user, err
}
