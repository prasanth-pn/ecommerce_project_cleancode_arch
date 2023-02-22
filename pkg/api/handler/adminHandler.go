package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

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

// @Summary List user for admin
// @ID ListUsers by admin
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list/users [get]
func (cr *AdminHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	fmt.Println(page)
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	fmt.Println(page, pageSize, "handler")
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}
	//fmt.Println(pagenation, "pagenation")
	users, metadata, err := cr.adminService.ListUsers(pagenation)
	fmt.Println(metadata, "metadata")

	if err != nil {
		respons := response.ErrorResponse("cannot show users", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}
	respons := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}

// @Summary List user for admin
// @ID AdminAddProducts
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param AdminAddProducts body domain.Product{} true "AdminAddProduct"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add/products [post]
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

// --------------------add category-----------------------------------------------------
// @Summary List user for admin
// @ID AddCategory by admin
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param AdminAddCategory body domain.Category{} true "AdminAddCategory"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add/category [post]
func (cr *AdminHandler) AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := cr.adminService.AddCategory(c.Request.Context(), category)
	if err != nil {
		respon := response.ErrorResponse("oops products not added ", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respon)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", nil)
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}
func (cr *AdminHandler) AddBrand(c *gin.Context) {
	var brand domain.Brand
	if err := c.BindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err := cr.adminService.AddBrand(c.Request.Context(), brand)
	if err != nil {
		respon := response.ErrorResponse("oops brand not added", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(c, respon)
		return

	}
	respons := response.SuccessResponse(true, "successfully brand added", nil)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respons)

}

func (cr *AdminHandler) AddModel(c *gin.Context) {
	var Model domain.Model
	if err := c.BindJSON(&Model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err := cr.adminService.AddModel(c.Request.Context(), Model)

	if err != nil {
		resp := response.ErrorResponse("ooops failed to add model", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadGateway)
		utils.ResponseJSON(c, resp)
		return
	}
	respons := response.SuccessResponse(true, "successfully inserted", Model)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respons)

}
