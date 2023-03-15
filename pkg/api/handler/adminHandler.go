package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @title Go + Gin ecommerce API
// @version 1.0
// @description This is a sample server  server. You can visit the GitHub repository at https://github.com/prasanth-pn/clean-code-architecture-ecommerce

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

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
	p("page", page, "pagesize", pagesize)

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
		res := response.SuccessResponse(true, "the user  is already blocked", user)
		c.Writer.WriteHeader(200)
		utils.ResponseJSON(c, res)
		return
	}
	err := cr.authService.BlockUnblockUser(uint(user_id), true)
	if err != nil {
		resp := response.ErrorResponse("error is happend from block user", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, resp)
	}
	resp := response.SuccessResponse(true, "user is bllocked successfully", user)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, resp)

}
func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	user, _ := cr.authService.FindUserById(uint(user_id))

	if !user.Block_Status {
		res := response.SuccessResponse(true, "user already unblocked", user)
		c.Writer.WriteHeader(200)
		utils.ResponseJSON(c, res)
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
func (cr *AdminHandler) SearchUserByName(c *gin.Context) {
	name := c.Query("name")
	name = "%" + name + "%"
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize := 5
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	users, metadata, err := cr.adminService.SearchUserByName(pagenation, name)
	if len(users) == 0 {
		res := response.ErrorResponse("can't find the user", "enter the correct name", "please enter the correct validation")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	if err != nil {
		res := response.ErrorResponse("failed to search ", err.Error(), "user data is not available")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	data := struct {
		Users    []domain.Users
		Metadata utils.Metadata
	}{
		Users:    users,
		Metadata: metadata,
	}
	res := response.SuccessResponse(true, "search completed", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// --------------------add category-----------------------------------------------------
// @Summary List user for admin
// @ID AddCategory by admin
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param category formdata domain.Category{} true "AdminAddCategory"
// @Param image formData file true "select image"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add/category [post]
func (cr *AdminHandler) AddCategory(c *gin.Context) {
	files, err := c.FormFile("image")
	if err != nil {
		res := response.ErrorResponse("please select the image to upload ", err.Error(), "no image is selected")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	extention := filepath.Ext(files.Filename)
	image := "category" + uuid.New().String() + extention
	var category domain.Category
	file := c.PostForm("category")
	err = json.Unmarshal([]byte(file), &category)
	if err != nil {
		res := response.ErrorResponse("error when unmarshal the jason file ", err.Error(), "error from parsing data")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	category.Image = image
	err = cr.adminService.AddCategory(category)
	if err != nil {
		respon := response.ErrorResponse("oops products not added ", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respon)
		return
	}
	c.SaveUploadedFile(files, "./public/"+image)
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
func (cr *AdminHandler) ListOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	listOrder, metadata, err := cr.adminService.ListOrder(pagenation)
	if err != nil {
		res := response.ErrorResponse("failed to list order", err.Error(), "failed to list order")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	data := struct {
		ListOrder []domain.Orders
		Metadata  utils.Metadata
	}{
		ListOrder: listOrder,
		Metadata:  metadata,
	}
	res := response.SuccessResponse(true, "successfully listed the order", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)

}
