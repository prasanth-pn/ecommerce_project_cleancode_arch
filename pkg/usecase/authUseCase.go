package usecase

import (
	//"clean/pkg/common/response"

	"clean/pkg/config"
	domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type authUseCase struct {
	authRepo   interfaces.AuthRepository
	mailConfig config.MailConfig
	config     config.Config
}

func NewAuthUseCase(repo interfaces.AuthRepository, mailConfig config.MailConfig, config config.Config) services.AuthUseCase {
	return &authUseCase{
		authRepo:   repo,
		mailConfig: mailConfig,
		config:     config,
	}
}

// ------------------------------------------register-------------------------
func (c *authUseCase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	retun, _ := c.FindUser(user.Email)
	fmt.Println(user.Email, retun.Email)
	//fmt.Println(retun.Email, user.Email)
	if retun.Email == user.Email {
		return user, errors.New("user already exists ")

	}

	_, err := c.authRepo.Register(ctx, user)
	if err != nil {
		return user, errors.New("email id already exists")
	}
	fmt.Println(user, "register")

	//fmt.Printf("\n\n %v ")
	return user, err
}

// ---------------------------------send mail-------------------
func (c *authUseCase) SendVerificationEmail(email string) error {
	rand.Seed(time.Now().UnixNano())
	max := 9999
	min := 1000
	code := rand.Intn((max - min) + min)
	message := fmt.Sprintf("\n the verification code is \n \n %d \n\n use to verify \n your account.\n have good shoping experience", code)
	// send random code to user's email
	err := c.mailConfig.SendMail(c.config, email, message)
	if err != nil {

		return err

	}

	err = c.authRepo.StoreVerificationDetails(email, code)
	fmt.Println(err, "repositoryS")
	if err != nil {
		return err
	}
	return nil
}


// -----------------------------------------verifyUser-----------------------------
func (c *authUseCase) VerifyUser(email, password string) error {
	user, err := c.FindUser(email)
	fmt.Println(err, user, "kijsdfghbjdbjekhsafbnfjhnjknb")
	if err != nil {
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
	user, err := c.authRepo.FindUser(email)

	return user, err
}

// --------------------------veryfypassword-------------------------------------------------------------
func VerifyPassword(password, dbpassword string) bool {
	return password == dbpassword
}

// ------------------------------------------AdminRegister-----------------------------------
func (c *authUseCase) AdminRegister(ctx context.Context, admin domain.Admins) (domain.Admins, error) {
	err := c.authRepo.AdminRegister(ctx, admin)

	return admin, err
}
func (c *authUseCase) VerifyAdmin(username, password string) error {
	//fmt.Println(username, "ifsdkohgndsiujkhgbdjkbhjfvbdhjffdgbauthcase")
	admin, err := c.authRepo.FindAdmin(username)
	fmt.Println(err, admin.UserName, "test user name")
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
