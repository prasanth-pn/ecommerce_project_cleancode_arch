package usecase

import (
	//"clean/pkg/common/response"

	"clean/pkg/config"
	domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
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
func (c *authUseCase) Register(user domain.Users) (domain.Users, error) {
	retun, _ := c.FindUser(user.Email)
	if retun.Email == user.Email {
		return user, errors.New("user already exists ")

	}

	_, err := c.authRepo.Register(user)
	fmt.Println(err)
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

	//messag = []byte(message)
	// send random code to user's email
	err := c.mailConfig.SendMail(c.config, email, []byte(message))
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

// -------------------------------------------user otp check-------------------
func (c *authUseCase) VerifyUserOtp(email string, code int) error {
	err := c.authRepo.VerifyOtp(email, code)
	return err
}
func (c *authUseCase) UpdateUserStatus(email string) error {
	err := c.authRepo.UpdateUserStatus(email)
	return err
}

// -----------------------------------------verifyUser-----------------------------
func (c *authUseCase) VerifyUser(email, password string) error {
	user, err := c.FindUser(email)

	if err != nil {
		return errors.New("username is incorrect")
	}
	if user.First_Name == "" {
		return errors.New("the user not registered")
	}
	IsValidPassword := VerifyPassword(password, user.Password)
	if !IsValidPassword {
		return errors.New(" password is incorrect")
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
func (c *authUseCase) AdminRegister(admin domain.Admins) (domain.Admins, error) {
	err := c.authRepo.AdminRegister(admin)
	return admin, err
}
func (c *authUseCase) VerifyAdmin(username, password string) error {
	//fmt.Println(username, "ifsdkohgndsiujkhgbdjkbhjfvbdhjffdgbauthcase")
	admin, err := c.authRepo.FindAdmin(username)
	if admin.UserName == "" {
		return errors.New("the admin not registered")
	}
	if err != nil {
		return errors.New("username is incorrect")
	}
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
func (c *authUseCase) FindUserById(user_id uint) (domain.Users, error) {
	user, err := c.authRepo.FindUserById(user_id)
	return user, err

}
func (c *authUseCase) BlockUnblockUser(user_id uint, val bool) error {
	err := c.authRepo.BlockUnblockUser(user_id, val)
	return err
}
