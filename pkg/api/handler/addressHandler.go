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

// @Summary AddAddress for user
// @ID AddAddress for user
// @Security BearerAuth
// @Tags User
// @Produce json
// @Param Address body domain.Address{} true "Address"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /user/checkout/add/address [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	fmt.Println(email)
	user, _ := cr.AuthService.FindUser(email)
	var address domain.Address
	if err := c.BindJSON(&address); err != nil {
		res := response.ErrorResponse("failed to get the data from front end", err.Error(), "failed to fetch data")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
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
	resp := response.SuccessResponse(true, "successfully added the address", address)
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
func (cr *UserHandler) GetAddressToEdit(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	address_id, _ := strconv.Atoi(c.Query("address_id"))
	user, _ := cr.AuthService.FindUser(email)
	address, err := cr.UserService.FindAddress(user.ID, uint(address_id))
	if err != nil {
		res := response.ErrorResponse("failed find the address", err.Error(), "failed to fetch data from server")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	respo := response.SuccessResponse(true, "success fully fetched the data", address)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respo)
}

func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	user, _ := cr.AuthService.FindUser(email)
	address_id, _ := strconv.Atoi(c.Query("address_id"))
	var add domain.Address
	err := c.BindJSON(&add)
	if err != nil {
		res := response.ErrorResponse("failed to getting data from front end", err.Error(), "failed to fetch data from front end")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	if add.Area == "" || add.House == "" || add.City == "" || add.FName == "" || add.Phone_Number == 0 || add.Landmark == "" || add.Pincode == 0 {
		res := response.ErrorResponse("please enter the data to all fields", "the fields are empty", "please enter the data to all fields")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	err = cr.UserService.UpdateAddress(add, (user.ID), uint(address_id))
	add.Address_id = uint(address_id)
	add.User_Id = user.ID
	fmt.Println(err)
	if err != nil {
		res := response.ErrorResponse("failed to update uddress", err.Error(), "failed to update address")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "succefully updated", add)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}
