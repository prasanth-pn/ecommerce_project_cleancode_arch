package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	"clean/pkg/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
)

type Home struct {
	userid      string
	Name        string
	total_price int
	Amount      int
	OrderId     string
	Email       string
	Contact     string
}

func (cr *UserHandler) RazorPay(c *gin.Context) {
	//email:=c.Writer.Header().Get("email")
	email := "prasanthpn68@gmail.com"
	user, _ := cr.AuthService.FindUser(email)

	sum, err := cr.UserService.FindTheSumOfCart(int(user.ID))
	if err != nil {
		c.JSON(422, gin.H{
			"message": "we can't find the sum of total cart",
		})
	}
	client := razorpay.NewClient("rzp_test_0h8oXuKI0kORyw", "McmgVREukL239BhjpTuS4j3t")
	razorpaytotal := sum * 100
	data := map[string]interface{}{
		"amount":   razorpaytotal,
		"currency": "INR",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		c.JSON(404, gin.H{
			"fail": "error creating order",
		})
		c.Abort()
		return
	}
	value := fmt.Sprint(body["id"])
	fmt.Println("\n oderder_id", value, "\n\n ")
	user_id := fmt.Sprint(user.ID)
	Home := Home{
		userid:      user_id,
		Name:        user.First_Name,
		total_price: sum,
		Amount:      razorpaytotal,
		OrderId:     value,
		Email:       user.Email,
		Contact:     user.Password,
	}
	order := domain.Orders{
		Created_at:     time.Now(),
		User_Id:        user.ID,
		Order_Id:       value,
		Total_Amount:   uint(sum),
		PaymentMethod:  "razorpay",
		Payment_Status: "uncomplete",
		Order_Status:   "ordered",
		Address_Id:     1,
	}
	err = cr.UserService.CreateOrder(order)
	if err != nil {
		res := response.ErrorResponse("order is not created", err.Error(), nil)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	c.HTML(200, "app.html", Home)
}
func (cr *UserHandler) Payment_Success(c *gin.Context) {
	payment_id := c.Query("paymentid")
	orderid := c.Query("orderid")
	orderid = strings.Trim(orderid, " ")
	//signature := c.Query("signature")
	order, err := cr.UserService.SearchOrder(orderid)
	if err != nil {
		res := response.ErrorResponse("order is not successfull", err.Error(), "order is failed")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	err = cr.UserService.UpdateOrders(payment_id, orderid)
	if err != nil {
		res := response.ErrorResponse("order is no is no updated", err.Error(), "order updation problem")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
	}
	cart, err := cr.UserService.ViewCart(order.User_Id)
	if err != nil {
		res := response.ErrorResponse("error in list cart", err.Error(), "update orders listcart")
		utils.ResponseJSON(c, res)
		return
	}
	for _, list := range cart {

		err = cr.UserService.Insert_To_My_Order(list, orderid)
		fmt.Println(err)

	}
	//clear the cart
	err = cr.UserService.ClearCart(order.User_Id)
	fmt.Println(err)
	//fmt.Println(cart)

	c.HTML(200, "success.html", nil)
}
