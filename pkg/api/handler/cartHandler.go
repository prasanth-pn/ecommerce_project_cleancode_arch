package handler

import (
	"clean/pkg/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cr *UserHandler) AddToCart(c *gin.Context) {
	//th := c.GetString("id")
	//fmt.Println(th, "this is ths get string validation")
	email := c.Writer.Header().Get("email")

	var user domain.UserResponse

	user, _ = cr.AuthService.FindUser(email)

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
	fmt.Println(err, "error printing ")
	fmt.Println(products.Price, products.Quantity)

	total := products.Price * float32(ProductDetails.Quantity)
	fmt.Println(total)
	product_id := ProductDetails.Product_id
	product_quantity := ProductDetails.Quantity
	cart := domain.Cart{
		
		User_Id:     user.ID,
		ProductID:   ProductDetails.Product_id,
		Quantity:    ProductDetails.Quantity,
		Total_Price: total,
	}
	fmt.Println(cart, product_id, product_quantity)
	Cart,err:=cr.UserService.ListCart(user.ID)
	for _,l:=range Cart{
		if l.ProductID==product_id{
			
		}
	}
	// // getting price for setting total amount
	// total:=products.Price*float32(ProductDetails.Quantity)
	// product_id:=ProductDetails.Product_id
	// product_quantity:=ProductDetails.Quantity
	// cart:=domain.Cart{
	// 	ProductID: ProductDetails.Product_id,
	// 	Quantity: ProductDetails.Quantity,
	// 	User_Id: user.ID,
	// 	Total_Price: uint(total),
	// }
	// in the cart find the product already exists

	// var Cart []domain.Cart

}
