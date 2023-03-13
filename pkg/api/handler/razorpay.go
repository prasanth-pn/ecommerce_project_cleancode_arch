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
		res := response.ErrorResponse("error in find the sum of cart ", err.Error(), "error in findinsumCArt")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	client := razorpay.NewClient("rzp_test_0h8oXuKI0kORyw", "McmgVREukL239BhjpTuS4j3t")
	razorpaytotal := sum * 100
	data := map[string]interface{}{
		"amount":   razorpaytotal,
		"currency": "INR",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		res := response.ErrorResponse("erro in crearte datat in  client order", err.Error(), "order.create")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	value := fmt.Sprint(body["id"])
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
	cart, err := cr.UserService.ListViewCart(order.User_Id)
	if err != nil {
		res := response.ErrorResponse("error in list cart", err.Error(), "update orders listcart")
		utils.ResponseJSON(c, res)
		return
	}
	for _, list := range cart {
		err = cr.UserService.Insert_To_My_Order(list, orderid)
		if err != nil {
			res := response.ErrorResponse("error in insert into order", err.Error(), "insert into myorder")
			c.Writer.WriteHeader(422)
			utils.ResponseJSON(c, res)
			return
		}
	}
	//clear the cart
	err = cr.UserService.ClearCart(order.User_Id)
	if err != nil {
		res := response.ErrorResponse("error in clearCArt", err.Error(), "error in clear cart")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	//c.HTML(200, "success.html", "success")
	res := response.SuccessResponse(true, "payment success", "payment success")
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)

}
