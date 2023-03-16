package handler

import (
	response "clean/pkg/common/response"
	domain "clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	utils "clean/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AuthHandler struct {
	authUseCase services.AuthUseCase
	jwtService  services.JWTService
}

func NewAuthHandler(usecase services.AuthUseCase, jwtUseCase services.JWTService) *AuthHandler {
	return &AuthHandler{
		authUseCase: usecase,
		jwtService:  jwtUseCase,
	}

}

// @Summary Register for user
// @ID Register for user
// @Tags USERTAUTHENTICATION
// @Produce json
// @Param Register body domain.Users{} true "Register"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /user/register [post]
func (cr *AuthHandler) Register(c *gin.Context) {
	var user domain.Users
	if err := c.BindJSON(&user); err != nil {
		res := response.ErrorResponse("errror while getting the data from user side", err.Error(), "error while fetching data")
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	user, err := cr.authUseCase.Register(user)
	reply := "welcome  " + user.First_Name

	if err != nil {
		respons := response.ErrorResponse("failed register", err.Error(), "register failed")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return

	}
	respons := response.SuccessResponse(true, "user registration completed  successfully", reply)
	copier.Copy(&respons, &reply)
	c.JSON(http.StatusOK, respons)

}

// -------------------------------UserLogin----------------------------------------
// @Summary Login for user
// @ID user login authentication
// @Tags USERTAUTHENTICATION
// @Produce json
// @Param userLogin body domain.Login{} true "userLogin"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /user/login [post]
func (cr *AuthHandler) UserLogin(c *gin.Context) {
	fmt.Println(time.Now())

	var user domain.Login
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := cr.authUseCase.VerifyUser(user.Email, user.Password)

	if err != nil {
		respons := response.ErrorResponse("failed to login", err.Error(), user)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("\n\n%v", respons)
		utils.ResponseJSON(c, respons)
		return
	}
	fmt.Println(user)
	users, _ := cr.authUseCase.FindUser(user.Email)

	token := cr.jwtService.GenerateToken(uint(users.ID), users.First_Name, user.Email)

	fmt.Println(users.First_Name, token)
	fmt.Printf("\n\ntockenygbhuy : %v\n\n", token)
	users.Password = ""
	users.Token = token
	respons := response.SuccessResponse(true, "Login successfully", users.Token)
	copier.Copy(&respons, nil)
	c.JSON(http.StatusOK, respons)
}

// ----------------------------------adminRegister-----------------------------------
// @Summary Register for Admin
// @ID AdminRegister for Admin
// @Tags Admin
// @Produce json
// @Param AdminRegister body domain.Admins{} true "AdminRegister"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /admin/register [post]
func (cr *AuthHandler) AdminRegister(c *gin.Context) {
	var admin domain.Admins
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	admin, err := cr.authUseCase.AdminRegister(admin)
	if err != nil {
		respons := response.ErrorResponse("failed to register", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("\n\n%v", respons)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "login successfully", admin)
	copier.Copy(&respons, &admin)
	c.JSON(http.StatusOK, respons)

}

// @Summary Admin Login for Admin
// @ID AdminLogin for Admin
// @Tags ADMINLOGIN
// @Produce json
// @Param AdminLogin body domain.Admins{} true "AdminLogin"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /admin/login [post]
func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	var admin domain.Admins

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(admin.UserName)
	err := cr.authUseCase.VerifyAdmin(admin.UserName, admin.Password)
	if err != nil {
		respons := response.ErrorResponse("failes to login", err.Error(), admin)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	user, _ := cr.authUseCase.FindAdmin(admin.UserName)
	token := cr.jwtService.GenerateToken(uint(user.ID), user.UserName, "admin")

	admin.Password = ""
	user.Token = token

	respons := response.SuccessResponse(true, "login successful", user.Token)
	copier.Copy(&respons, respons.Data)
	c.JSON(http.StatusOK, respons)

}

// @Summary Send EmailOtp for User
// @ID SendUserMail
// @Tags USERVERIFICATION
// @Produce json
// @Param  email query  string true "email"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /user/send/verificationmail [post]
func (cr *AuthHandler) SendUserMail(c *gin.Context) {
	email := c.Query("email")
	user, err := cr.authUseCase.FindUser(email)
	if err != nil {
		return
	}

	err = cr.authUseCase.SendVerificationEmail(email)
	fmt.Println(err, "errot in handler find the error please find it ")

	if err != nil {
		respons := response.ErrorResponse("Error whilw sending verification email", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(c, respons)
		return
	}
	user.Password = ""
	respons := response.SuccessResponse(true, "SUCCESS", "otp  successfully send")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}

// @Summary VerifyUserOtp for User
// @ID VerifyUserOtp for authentication
// @Tags USERVERIFICATION
// @Produce json
// @Param  email query  string true "email"
// @Param  code  query  int true "code"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /user/verify/otp [get]
func (cr *AuthHandler) VerifyUserOtp(c *gin.Context) {
	email := c.Query("email")
	code, _ := strconv.Atoi(c.Query("code"))
	err := cr.authUseCase.VerifyUserOtp(email, code)
	if err != nil {
		res := response.ErrorResponse("user not valid otp incorrect", err.Error(), nil)
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return

	}
	err = cr.authUseCase.UpdateUserStatus(email)
	if err != nil {
		respo := response.SuccessResponse(true, "user verified successfully", " login the page using username and password ")
		c.Writer.WriteHeader(http.StatusOK)
		utils.ResponseJSON(c, respo)
	}

}
