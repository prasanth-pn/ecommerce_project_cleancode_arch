package handler

import (
	"clean/pkg/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cr *UserHandler) AddToCart(c *gin.Context) {
	//th := c.GetString("id")
	//fmt.Println(th, "this is ths get string validation")
	email := c.Writer.Header().Get("email")

	var user domain.UserResponse

	user, _ = cr.AuthService.FindUser(email)
	fmt.Println(user.ID, user.Email, "user_idavailaljndsflk")

	fmt.Println(user.First_Name, user)
	// c.Writer.Header().Get("id")
	var ProductDetails struct {
		Product_id uint
		Quantity   uint
	}
	//var products domain.Product
	if err := c.BindJSON(&ProductDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	products, err := cr.UserService.FindProduct(ProductDetails.Product_id)
	if err != nil {
		log.Println("some error happend in finding product")
		return
	}
	fmt.Println(err, "error printing ")
	fmt.Println(products.Price, products.Quantity)
	//total price of products
	total := products.Price * float32(ProductDetails.Quantity)
	fmt.Println(total)
	product_id := ProductDetails.Product_id     // value from above struct
	product_quantity := ProductDetails.Quantity //value from above struct
	cart := domain.Cart{

		User_Id:     user.ID,
		ProductID:   ProductDetails.Product_id,
		Quantity:    ProductDetails.Quantity,
		Total_Price: total,
	}
	Cart, err := cr.UserService.ListCart(user.ID)
	fmt.Println(err)
	for _, l := range Cart {
		fmt.Println(l.ProductID, "productid new", product_id)
		if l.ProductID == product_id {
			quantity, err := cr.UserService.QuantityCart(product_id, user.ID)
			fmt.Println(product_quantity)
			totalprice := float32(product_quantity+quantity.Quantity) * products.Price
			fmt.Println(quantity, err, "atlast found", total)
			product_quantity += quantity.Quantity
			err = cr.UserService.UpdateCart(totalprice, product_quantity, product_id, user.ID)
			fmt.Println(err)
			c.JSON(200, gin.H{
				"msg": "quantity updated",
			})
			c.Abort()
			return
		}
	}
	err = cr.UserService.CreateCart(cart)
	fmt.Println(err)         
}
func (cr *UserHandler)ListCart(c *gin.Context){
	// var products domain.Product
	// var userid domain.Users
	// var cart domain.CartListResponse

	email := c.Writer.Header().Get("email")
	fmt.Println(email)



}
func (cr *UserHandler)Checkout(c *gin.Context){
	email:=c.Writer.Header().Get("email")
	fmt.Println(email)
}