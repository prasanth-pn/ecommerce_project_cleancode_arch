package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	"clean/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cr *UserHandler) AddAddress(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	fmt.Println(email)
	user, _ := cr.AuthService.FindUser(email)
	var address domain.Address

	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	address.User_Id = user.ID
	err := cr.UserService.AddAddress(address)
	if err != nil {
		resp := response.ErrorResponse("user address is not added", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(c, resp)
		return
	}
	resp := response.SuccessResponse(true, "successfully added the address", nil)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, resp)
}

// -----------------------------list Address-------------------------------------
func (cr *UserHandler) ListAddress(c *gin.Context) {
	s := c.Writer.Header().Get("id")
	user_id, _ := strconv.ParseUint(s, 10, 64)
	address, err := cr.UserService.ListAddress(uint(user_id))
	if err != nil {
		respo := response.ErrorResponse("listing address is failed", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadGateway)
		utils.ResponseJSON(c, respo)
		return
	}
	respo := response.SuccessResponse(true, "successfully addres listed", address)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respo)

}
