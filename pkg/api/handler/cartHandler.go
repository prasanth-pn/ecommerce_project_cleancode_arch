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

func (cr *UserHandler) AddToCart(c *gin.Context) {
	//th := c.GetString("id")
	//fmt.Println(th, "this is ths get string validation")
	id := c.Writer.Header().Get("id")
	user, _ := strconv.ParseUint(id, 0, 0)
	user_id := uint(user)

	var ProductDetails struct {
		Product_id uint
		Quantity   uint
	}
	if err := c.BindJSON(&ProductDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return

	}

	products, err := cr.UserService.FindProduct(ProductDetails.Product_id)

	if err != nil {
		respons := response.ErrorResponse("error finding ", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(c, respons)
		return
	}
	fmt.Println(products.Price, products.Quantity)
	//total price of products
	total := products.Price * float32(ProductDetails.Quantity)
	product_id := ProductDetails.Product_id     // value from above struct
	product_quantity := ProductDetails.Quantity //value from above struct
	cart := domain.Cart{

		User_Id:     user_id,
		Product_Id:  ProductDetails.Product_id,
		Quantity:    ProductDetails.Quantity,
		Total_Price: total,
	}
	Cart, err := cr.UserService.ListCart(user_id)
	if err != nil {
		res := response.ErrorResponse("error in the list cart", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
	}
	for _, l := range Cart {
		if l.Product_Id == product_id {
			quantity, err := cr.UserService.QuantityCart(product_id, user_id)
			if err != nil {
				res := response.ErrorResponse("failed to check quantity ", err.Error(), nil)
				c.Writer.WriteHeader(http.StatusBadRequest)
				utils.ResponseJSON(c, res)
				return
			}
			totalprice := float32(product_quantity+quantity.Quantity) * products.Price
			product_quantity += quantity.Quantity
			err = cr.UserService.UpdateCart(totalprice, product_quantity, product_id, user_id)

			if err != nil {
				res := response.ErrorResponse("failed to update cart cart update", err.Error(), nil)
				c.Writer.WriteHeader(http.StatusBadRequest)
				utils.ResponseJSON(c, res)
				return
			}
			c.JSON(200, gin.H{
				"msg": "quantity updated",
			})
			c.Abort()
			return

		}
	}
	err = cr.UserService.CreateCart(cart)
	if err != nil {
		res := response.ErrorResponse("failed to put product into cart", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "successfully product aded into cart", "enjoy the shopping")
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

//--------------------------------------------------------listCart-------------------------

func (cr *UserHandler) ListCart(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	user, err := cr.AuthService.FindUser(email)
	if err != nil {
		respons := response.ErrorResponse("oops user not found", err.Error(), nil)
		c.Writer.WriteHeader(502)
		utils.ResponseJSON(c, respons)
		return
	}
	var cart []domain.CartListResponse

	cart, err = cr.UserService.ViewCart(user.ID)
	if err != nil {
		respons := response.ErrorResponse("oops carts not fetched ", err.Error(), nil)
		c.Writer.WriteHeader(502)
		utils.ResponseJSON(c, respons)
		return

	}
	var totalPrice float32
	totalPrice, err = cr.UserService.TotalCartPrice(user.ID)
	if err != nil {
		respo := response.ErrorResponse("cannot  calculate total amount", err.Error(), nil)
		c.Writer.WriteHeader(502)
		utils.ResponseJSON(c, respo)
	}

	respons := response.SuccessResponse(true, "successfully listed cart", cart, totalPrice)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respons)

	//cart:=cr.UserService.ListCart(user.ID)

}
func (cr *UserHandler) UpdateCart(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	quantity, _ := strconv.Atoi(c.Query("quantity"))
	product, err := cr.UserService.FindProduct(uint(product_id))
	if err != nil {
		return
	}
	totalprice := (quantity) * int(product.Price)
	fmt.Println(totalprice, "totalprice")

	fmt.Println(user_id, "   ", product_id)
	//err = cr.UserService.UpdateCart(totalprice, product_quantity, product_id, user_id)
	err = cr.UserService.UpdateCart(float32(totalprice), uint(quantity), uint(product_id), uint(user_id))
	if err != nil {
		res := response.ErrorResponse("it is not updated something error ", err.Error(), nil)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	} else {
		res := response.SuccessResponse(true, "succefully updated the cart", "success wish you a happy shopping")
		c.Writer.WriteHeader(http.StatusOK)
		utils.ResponseJSON(c, res)
		return
	}

}
func (cr *UserHandler) Checkout(c *gin.Context) {
	// email := c.Writer.Header().Get("email")
	// address_id, _ := strconv.Atoi(c.Query("address_id"))

	// user, _ := cr.AuthService.FindUser(email)
	// cart, _ := cr.UserService.ViewCart(uint(user.ID))
	// payment_method := c.Query("payment_method")
	// //fmt.Println(cart)
	// for _, l := range cart {
	// 	err := cr.UserService.CreateOrder(l, payment_method, address_id, int(user.ID))
	// 	fmt.Println(err)

	// }

}
