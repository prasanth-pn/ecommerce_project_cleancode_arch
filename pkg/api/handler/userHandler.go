package handler

import (
	//"net/http"
	//domain "clean/pkg/domain"
	"clean/pkg/common/response"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// 	response "clean/pkg/common/response"
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
