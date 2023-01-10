package handler

import (
	response "clean/pkg/common/response"
	domain "clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	utils "clean/pkg/utils"
	"fmt"
	"net/http"

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

func (cr *AuthHandler) Register(c *gin.Context) {
	var user domain.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user, err := cr.authUseCase.Register(c.Request.Context(), user)
	reply := "welcome  " + user.First_Name

	if err != nil {
		//fmt.Println("user already exists")
		respons := response.ErrorResponse("failed register", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("\n\n%v", respons)
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
// @Produce json
// @Param        Email   path      string  true  "Email : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var user domain.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//fmt.Println(user.Email)
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

	token := cr.jwtService.GenerateToken(uint(user.User_Id), user.First_Name, user.Email)

	fmt.Println(user.First_Name, token)
	fmt.Printf("\n\ntockenygbhuy : %v\n\n", token)
	users.Password = ""
	users.Token = token
	respons := response.SuccessResponse(true, "Login successfully", users.Token)
	copier.Copy(&respons, nil)
	c.JSON(http.StatusOK, respons)
}

// ----------------------------------adminRegister-----------------------------------
func (cr *AuthHandler) AdminRegister(c *gin.Context) {
	var admin domain.Admins
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	admin, err := cr.authUseCase.AdminRegister(c.Request.Context(), admin)
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
func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	var admin domain.Admins
	
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := cr.authUseCase.VerifyAdmin(admin.UserName, admin.Password)
	if err != nil {
		respons := response.ErrorResponse("failes to login", err.Error(), nil)
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
