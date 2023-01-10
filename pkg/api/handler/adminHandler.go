package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService services.AdminUseCase
	jwtService   services.JWTService
}

func NewAdminHandler(usecase services.AdminUseCase, jwtService services.JWTService) *AdminHandler {
	return &AdminHandler{
		adminService: usecase,
		jwtService:   jwtService,
	}
}
func (cr *AdminHandler) ListUsers(c *gin.Context) {
	users, err := cr.adminService.ListUsers()
	if err != nil {
		respons := response.ErrorResponse("cannot show users", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", users)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}

func (cr *AdminHandler) AddProducts(c *gin.Context) {
	var products domain.Product
	if err := c.BindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	 err := cr.adminService.AddProducts(c.Request.Context(), products)
	if err != nil {
		respons := response.ErrorResponse("oops products not added", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", nil)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}
//--------------------add category-----------------------------------------------------
func (cr *AdminHandler)AddCategory(c *gin.Context){
	var category domain.Category
	if err:=c.BindJSON(&category);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	err:=cr.adminService.AddCategory(c.Request.Context(),category)
	if err!=nil{
		respon:=response.ErrorResponse("oops products not added ",err.Error(),nil)
		c.Writer.Header().Set("Content_Type","application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c,respon)
		return
	}
	respons:=response.SuccessResponse(true,"SUCCESS",nil)
	c.Writer.Header().Set("Content-Type","application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c,respons)

}
