package handler

import (
	//"net/http"
	//domain "clean/pkg/domain"
	"clean/pkg/common/response"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	fmt.Println(product)

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
