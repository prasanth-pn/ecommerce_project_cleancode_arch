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
	authService  services.AuthUseCase
}

func NewAdminHandler(usecase services.AdminUseCase, jwtService services.JWTService, authusecase services.AuthUseCase) *AdminHandler {
	return &AdminHandler{
		adminService: usecase,
		jwtService:   jwtService,
		authService:  authusecase,
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
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}
	users, metadata, err := cr.adminService.ListUsers(pagenation)
	if err != nil {
		respons := response.ErrorResponse("cannot show users", err.Error(), nil)
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
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}
func (cr *AdminHandler) ListBlockedUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}

	users, metadata, err := cr.adminService.ListBlockedUsers(pagenation)
	result := struct {
		Users []domain.Users
		Meta  *utils.Metadata
	}{
		Users: *users,
		Meta:  metadata,
	}
	if err != nil {

		res := response.ErrorResponse("blocked users are not listed", err.Error(), "wait a minute")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	} else {
		res := response.SuccessResponse(true, "succesfully listed blocked users", result)
		c.Writer.WriteHeader(http.StatusOK)
		utils.ResponseJSON(c, res)
	}

}
func (cr *AdminHandler) BlockUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	user, _ := cr.authService.FindUserById(uint(user_id))
	fmt.Println(user, "\n\n\n\n ")
	fmt.Println(user.Block_Status)

	if user.Block_Status {
		c.JSON(200, gin.H{
			"message ": "the user is blocked",
		})
		c.Abort()
		return
	}
	err := cr.authService.BlockUnblockUser(uint(user_id), true)
	if err != nil {
		resp := response.ErrorResponse("error is happend from block user", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, resp)
	}
	//fmt.Println(err)

	resp := response.SuccessResponse(true, "user is bllocked successfully", "blocked successsfully")
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, resp)

}
func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	user, _ := cr.authService.FindUserById(uint(user_id))

	if !user.Block_Status {
		c.JSON(300, gin.H{
			"message": "user is already unblocked",
		})
		c.Abort()
		return
	} else {
		err := cr.authService.BlockUnblockUser(uint(user_id), false)
		if err == nil {
			res := response.SuccessResponse(true, "successfully unblocked", user.First_Name)
			c.Writer.WriteHeader(200)
			utils.ResponseJSON(c, res)
			return
		}
	}

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
