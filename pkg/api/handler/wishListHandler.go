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

func (cr *UserHandler) AddWishList(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	user, err := cr.AuthService.FindUser(email)
	if err != nil {
		fmt.Println(err, "printing error")
	}
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	//checking the product already wishlisted or not
	count := cr.UserService.Count_WishListed_Product(user.ID, uint(product_id))
	if count >= 1 {
		c.JSON(400, gin.H{
			"msg": "product already added to wishlist",
		})
		c.Abort()
		return
	}
	wishlist := domain.WishList{
		UserID:     user.ID,
		Product_Id: uint(product_id),
	}
	err = cr.UserService.AddTo_WishList(wishlist)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to add to wish list query mistake ",
		})
		c.Abort()
		return

	} else {
		c.JSON(200, gin.H{
			"message": "wish list is added",
		})
		c.Abort()
		return
	}

}

func (cr *UserHandler) ViewWishList(c *gin.Context) {
	// var wishlist domain.WishListResponse
	email := c.Writer.Header().Get("email")
	user, _ := cr.AuthService.FindUser(email)

	wishlist := cr.UserService.ViewWishList(user.ID)
	respons := response.SuccessResponse(true, "SUCCESS", wishlist)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)
}
func (cr *UserHandler) RemoveFromWishlist(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	user, _ := cr.AuthService.FindUser(email)
	fmt.Println(product_id, user.ID, "product id")
	cr.UserService.RemoveFromWishlist(int(user.ID),product_id)
}
