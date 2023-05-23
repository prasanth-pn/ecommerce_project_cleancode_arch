package handler

import (
	"strconv"

	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"

	"github.com/gin-gonic/gin"
)

// @Summary UserListOrder for user
// @ID UserListOrder for User
// @Tags ORDERMANAGEMENT
// @Security BearerAuth
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/list-order [get]
func (cr *UserHandler) ListOrder(c *gin.Context) {
	user_id := c.Writer.Header().Get("id")
	userid, _ := strconv.Atoi(user_id)
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	order, metadata, err := cr.UserService.ListOrder(pagenation, uint(userid))
	if err != nil {
		res := response.ErrorResponse("error while listing order", err.Error(), "error in userservice .listorder")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	// add, err := cr.UserService.FindAddress(uint(userid), add_id)
	// if err != nil {
	// 	res := response.ErrorResponse("error while find address", err.Error(), "error in find address")
	// 	c.Writer.WriteHeader(422)
	// 	utils.ResponseJSON(c, res)
	// 	return
	// }
	// result := struct {
	// 	Order   []domain.ListOrder
	// 	Address domain.Address
	// }{
	// 	Order:   order,
	// 	Address: add,
	// }
	data := struct {
		Order    []domain.OrderResponse
		Metadata utils.Metadata
	}{
		Order:    order,
		Metadata: metadata,
	}

	res := response.SuccessResponse(true, "successfully listed", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)

}
