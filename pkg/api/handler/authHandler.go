package handler

import (
	response "clean/pkg/common/response"
	domain "clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	utils "clean/pkg/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	// "github.com/golang-jwt/jwt/v4"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type AuthHandler struct {
	authUseCase services.AuthUseCase
	jwtUseCase services.JWTService
}

func NewAuthHandler(usecase services.AuthUseCase,jwtUseCase services.JWTService) *AuthHandler {
	return &AuthHandler{
		authUseCase: usecase,
		jwtUseCase: jwtUseCase,
	}

}

// func (cr *UseHandler)FindAll(c *gin.Context){
// 	users,err:=cr.userUseCase.FindAll(c.Request.Context())
// 	if err!=nil{
// 		c.AbortWithStatus(http.StatusNotFound)

// 	}else{
// 		response:=[]Response{}
// 		copier.Copy(&response,&users)
// 		c.JSON(http.StatusOK,response)

// 	}
// }

// @Summary add a new item to the users
// @ID register user
// @Produce json
// @Param data body users true "users data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /user/register [post]
func (cr *AuthHandler) Register(c *gin.Context) {
	var user domain.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	_,err:=cr.authUseCase.FindUser(user.Email)
	if err!=nil{
		//errors.New("the user already exists")
		return 
	}
	

	reply := "welcome  " + user.First_Name
	user, err= cr.authUseCase.Register(c.Request.Context(), user)

	if err != nil {
		respons := response.ErrorResponse("failed to login", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("\n\n%v", respons)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "login successfully", reply)
	copier.Copy(&respons, &reply)
	c.JSON(http.StatusOK, respons)

}

// -------------------------------UserLogin----------------------------------------
func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var user domain.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := cr.authUseCase.VerifyUser(user.Email, user.Password)
	if err != nil {
		respons := response.ErrorResponse("failed to login", err.Error(), nil)
		// c.Header().Add("Content-Type","application/json")
		c.Writer.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusUnauthorized)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("\n\n%v", respons)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "Login successfully", nil)
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
		respons := response.ErrorResponse("failed to login", err.Error(), nil)
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
func (cr *AuthHandler) AdminLogin(c *gin.Context){
	var admin domain.Admins

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//fmt.Println(admin.UserName,admin.Password,"admionhfjdfnjkfdnjf")
	err:=cr.authUseCase.VerifyAdmin(admin.UserName,admin.Password)
	if err!=nil{
		respons:=response.ErrorResponse("failes to login",err.Error(),nil)
		c.Writer.Header().Set("Content-Type","application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c,respons)
		return 
	}
	user,_:=cr.authUseCase.FindAdmin(admin.UserName)
	fmt.Println(user)
	token:=cr.jwtUseCase.GenerateToken(uint(user.ID),user.UserName,"admin")
	
	fmt.Println(user.UserName,token)
	fmt.Printf("\n\ntockenygbhuy : %v\n\n",token)
	admin.Password=""
	user.Token=token

	respons:=response.SuccessResponse(true,"login successful",user.Token)
	copier.Copy(&respons,respons.Data)
	c.JSON(http.StatusOK,respons)

}




// func (cr *UserHandler)FindAll(c *gin.Context){
// 	users,err:=cr.userUseCase.FindAll(c.Request.Context())
// 	if err!=nil{
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}else{
// 		response:=response.SuccessResponse(true,"showing users",users)
// 		copier.Copy(&response,&users)
// 		c.JSON(http.StatusOK,response)
// 	}
// }
