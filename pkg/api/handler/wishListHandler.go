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
	fmt.Println(count, "count")
	if count >= 1 {
		res := response.ErrorResponse("The product already added to the wishlist", "product already exists", "duplicate product")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	wishlist := domain.WishList{
		UserID:     user.ID,
		Product_Id: uint(product_id),
	}
	product, err := cr.UserService.FindProduct(uint(product_id))
	if err!=nil{
		res:=response.ErrorResponse("failed to find the product",err.Error(),"product not found")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c,res)
		return 
	}
	err = cr.UserService.AddTo_WishList(wishlist)
	if err != nil {
		res := response.ErrorResponse("failed to add to wish list", err.Error(), "error while adding product to wishlist")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return

	}
	data := struct {
		Product_id   int
		Product_name string
		Price        int
		Description  string
		Image        string
	}{
		Product_id:   product_id,
		Product_name: product.Product_Name,
		Price:        int(product.Price),
		Description:  product.Description,
		Image:        product.Image,
	}
	res := response.SuccessResponse(true, "successfully added", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)

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
	product, err := cr.UserService.FindProduct(uint(product_id))
	if err != nil {
		res := response.ErrorResponse("failed to identify the project", err.Error(), "failed to find project")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return

	}
	err = cr.UserService.RemoveFromWishlist(int(user.ID), product_id)
	if err != nil {
		res := response.ErrorResponse("failed to remove from wishlist", err.Error(), "failed to remove from wishlist")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	data := struct {
		Product_Name        string
		Product_Id          int
		Product_image       string
		product_description string
	}{
		Product_Name:        product.Product_Name,
		Product_Id:          product_id,
		Product_image:       product.Image,
		product_description: product.Description,
	}
	res := response.SuccessResponse(true, "successfully deleted from wish list", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}
func (cr *UserHandler) WishListTo_Cart(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	user, _ := cr.AuthService.FindUser(email)
	product, _ := cr.UserService.FindProduct(uint(product_id))
	cart, _ := cr.UserService.FindCart(user.ID, uint(product_id))
	if cart.Quantity != 0 {
		res := response.ErrorResponse("failed to  add cart ", "failed prduct already there", "failes to create cart")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return

	} else {
		cart := domain.Cart{
			Product_Id:  uint(product_id),
			Quantity:    1,
			User_Id:     user.ID,
			Total_Price: product.Price,
		}
		cart, err := cr.UserService.CreateCart(cart)
		if err != nil {
			res := response.ErrorResponse("failed to add to create", err.Error(), "failes to create cart")
			c.Writer.WriteHeader(422)
			utils.ResponseJSON(c, res)
			return
		}
		cr.UserService.RemoveFromWishlist(int(user.ID), product_id)
		res := response.SuccessResponse(true, "successfully added to Cart", cart)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
	}

}
