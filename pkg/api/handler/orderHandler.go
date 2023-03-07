package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	"clean/pkg/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cr *UserHandler) ListOrder(c *gin.Context) {
	user_id := c.Writer.Header().Get("id")
	userid, _ := strconv.Atoi(user_id)
	order, add_id, err := cr.UserService.ListOrder(uint(userid))
	fmt.Println(order)
	if err != nil {
		res := response.ErrorResponse("error while listing order", err.Error(), "error in userservice .listorder")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	add, err := cr.UserService.FindAddress(uint(userid), add_id)
	if err != nil {
		res := response.ErrorResponse("error while find address", err.Error(), "error in find address")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	result := struct {
		Order   []domain.ListOrder
		Address domain.Address
	}{
		Order:   order,
		Address: add,
	}

	res := response.SuccessResponse(true, "successfully listed", result)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)

}
