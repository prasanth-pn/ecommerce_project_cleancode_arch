package handler

import (
	//"net/http"
	//domain "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	services "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/usecase/interfaces"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// 	response "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	// 	"github.com/gin-gonic/gin"
	// 	"github.com/jinzhu/copier"
)

type UserHandler struct {
	UserService services.UserUseCase
	AuthService services.AuthUseCase
}

func NewUserHandler(usecase services.UserUseCase, AuthService services.AuthUseCase) *UserHandler {
	return &UserHandler{
		UserService: usecase,
		AuthService: AuthService,
	}

}

// @Summary ListProducts for user
// @ID UserListProducts for user
// @Tags USERVIEW
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/list/products [get]
func (cr *UserHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	products, metadata, err := cr.UserService.ListProducts(pagenation)
	fmt.Println(metadata)
	result := struct {
		Products *[]domain.ProductResponse
		Meta     *utils.Metadata
	}{
		Products: &products,
		Meta:     &metadata,
	}
	if err != nil {
		respons := response.ErrorResponse("cannot show users", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}

// @Summary ListCategories for user
// @ID UserListCategories for user
// @Tags USERVIEW
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/list/categories [get]
func (cr *UserHandler) ListCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	catego, metadata, err := cr.UserService.ListCategories(pagenation)
	result := struct {
		Category *[]domain.Category
		Meta     *utils.Metadata
	}{
		Category: &catego,
		Meta:     &metadata,
	}
	if err != nil {
		res := response.ErrorResponse("canot list categories", err.Error(), nil)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	} else {
		res := response.SuccessResponse(true, "successfully category listed", result)
		c.Writer.WriteHeader(200)
		utils.ResponseJSON(c, res)
		return

	}

}

// @Summary ListProductsByCategories for user
// @ID UserListProductsByCategories for user
// @Tags USERVIEW
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Param category_id query string true "category_id"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/list/products_bycategories [get]
func (cr *UserHandler) ListProductsByCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	category_id, _ := strconv.Atoi(c.Query("category_id"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	product, metadata, err := cr.UserService.ListProductsByCategories(category_id, pagenation)
	result := struct {
		Product []domain.ProductResponse
		Meta    utils.Metadata
	}{
		Product: product,
		Meta:    metadata,
	}
	if err != nil {
		res := response.ErrorResponse("data is  not fetched", err.Error(), nil)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	} else {
		res := response.SuccessResponse(true, "successfully fetched the data", result)
		c.Writer.WriteHeader(200)
		utils.ResponseJSON(c, res)
	}
}

// @Sumamry UserProfile for user
// @ID UserProfile for user
// @Tags USERCREDENTIALS
// @Security BearerAuth
// @success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/profile [get]
func (cr *UserHandler) Profile(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	user, err := cr.AuthService.FindUserById(uint(user_id))
	if err != nil {
		res := response.ErrorResponse("failed to fetch the user details ", err.Error(), "failed to get  user details")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "successfully fetched the data", user)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary UserProfileEdit for user
// @ID UserProfileEdit for user
// @Tags USERCREDENTIALS
// @Security BearerAuth
// @Param profile formData file true "select a image"
// @Param user formData string true "usersdetails"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/profile-edit [patch]
func (cr *UserHandler) UserEdit(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	profile, _ := c.FormFile("profile")
	file := c.PostForm("user")
	var user domain.Users
	if err := json.Unmarshal([]byte(file), &user); err != nil {
		res := response.ErrorResponse("data is not fetched ", err.Error(), "data is not fetched from bindjson user")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	extention := filepath.Ext(profile.Filename)
	profimage := "profile" + uuid.New().String() + extention
	user.Password = ""
	user.User_Id = uint(user_id)
	user.Profile_Pic = profimage
	userResponse, err := cr.UserService.UpdateUser(user)
	if err != nil {
		res := response.ErrorResponse("failed update your profile update later", err.Error(), "try again later")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	c.SaveUploadedFile(profile, "./public/"+profimage)
	res := response.SuccessResponse(true, "success", userResponse)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}
