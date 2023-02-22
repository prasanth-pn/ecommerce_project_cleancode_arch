package handler

import (
	//"net/http"
	//domain "clean/pkg/domain"
	"clean/pkg/common/response"
	services "clean/pkg/usecase/interfaces"
	"clean/pkg/utils"
	"net/http"

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
	products, err := cr.UserService.ListProducts()
	if err != nil {
		respons := response.ErrorResponse("cannot show users", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", products)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)

}


