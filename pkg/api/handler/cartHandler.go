package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"

	"github.com/gin-gonic/gin"
)

// @Summary AddToCart for User
// @ID UserAddCart for user
// @Tags USERCARTMANAGEMENT
// @Security BearerAuth
// @Param Cart body domain.ProductDetails{} true "productDetails"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/add/cart [post]
func (cr *UserHandler) AddToCart(c *gin.Context) {
	var ResponseCart domain.Cart
	var totalPrice float32
	id := c.Writer.Header().Get("id")
	user, _ := strconv.Atoi(id)
	user_id := uint(user)
	var ProductDetails domain.ProductDetails
	if err := c.BindJSON(&ProductDetails); err != nil {
		res := response.ErrorResponse("failed to fetch data from user", err.Error(), "failed to fetch product details ")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
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
		Image:       products.Image,
		Product_Id:  ProductDetails.Product_id,
		Quantity:    ProductDetails.Quantity,
		Total_Price: total,
	}
	Cart, err := cr.UserService.ListViewCart(user_id)
	if err != nil {
		res := response.ErrorResponse("error in the list cart", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	count := 0
	for _, l := range Cart {
		if l.Product_id == product_id {
			count++
			quantity, err := cr.UserService.QuantityCart(product_id, user_id)
			if err != nil {
				res := response.ErrorResponse("failed to check quantity ", err.Error(), nil)
				c.Writer.WriteHeader(http.StatusBadRequest)
				utils.ResponseJSON(c, res)
				return
			}
			totalPrice = float32(product_quantity+quantity.Quantity) * products.Price
			product_quantity += quantity.Quantity
			ResponseCart, err = cr.UserService.UpdateCart(totalPrice, product_quantity, product_id, user_id)
			if err != nil {
				res := response.ErrorResponse("failed to update cart cart update", err.Error(), nil)
				c.Writer.WriteHeader(http.StatusBadRequest)
				utils.ResponseJSON(c, res)
				return
			}
		}
	}
	if count == 0 {
		ResponseCart, err = cr.UserService.CreateCart(cart)
		if err != nil {
			res := response.ErrorResponse("failed to put product into cart", err.Error(), nil)
			c.Writer.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(c, res)
			return
		}

	}

	data := struct {
		Product_Details domain.Cart
		Product         domain.Product
	}{
		Product_Details: ResponseCart,
		Product:         products,
	}
	res := response.SuccessResponse(true, "successfully product aded into cart", "enjoy the shopping", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

//--------------------------------------------------------listCart-------------------------

// @Summary ListCart for User
// @ID UserListCart for user
// @Tags USERCARTMANAGEMENT
// @Security BearerAuth
// @Param page query string true "page"
// @Param pagesize  query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/list/cart [get]
func (cr *UserHandler) ListCart(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	fmt.Println(email)
	user, err := cr.AuthService.FindUser(email)
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	if err != nil {
		respons := response.ErrorResponse("oops user not found", err.Error(), nil)
		c.Writer.WriteHeader(502)
		utils.ResponseJSON(c, respons)
		return
	}
	cart, metadata, err := cr.UserService.ListCart(pagenation, user.ID)
	if cart == nil {
		res := response.ErrorResponse("empty cart", "cart is empty", nil)
		c.Writer.WriteHeader(205)
		utils.ResponseJSON(c, res)
		return
	}
	fmt.Println(err)
	if err != nil {
		respons := response.ErrorResponse("oops carts not fetched ", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, respons)
		return
	}
	var totalPrice float32
	totalPrice, err = cr.UserService.TotalCartPrice(user.ID)
	if err != nil {
		respo := response.ErrorResponse("cannot  calculate total amount", err.Error(), nil)
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, respo)
		return
	}
	data := struct {
		Cart             []domain.CartListResponse
		Total_Cart_Value int
		Metadata         utils.Metadata
	}{
		Cart:             cart,
		Total_Cart_Value: int(totalPrice),
		Metadata:         metadata,
	}
	respons := response.SuccessResponse(true, "successfully listed cart", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, respons)
}

// @Sumary UpdateCart for User
// @ID UserUpdateCart for user
// @Tags USERCARTMANAGEMENT
// @Security BearerAuth
// @Param product_id query string true "page"
// @Param quantity query string true "quantity"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/update/cart [patch]
func (cr *UserHandler) UpdateCart(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	quantity, _ := strconv.Atoi(c.Query("quantity"))
	product, err := cr.UserService.FindProduct(uint(product_id))
	if err != nil {
		res := response.ErrorResponse("failed to find the product", err.Error(), "product is not available")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	totalprice := (quantity) * int(product.Price)
	fmt.Println(totalprice, "totalprice")

	//err = cr.UserService.UpdateCart(totalprice, product_quantity, product_id, user_id)
	ResponseCart, err := cr.UserService.UpdateCart(float32(totalprice), uint(quantity), uint(product_id), uint(user_id))
	if err != nil {
		res := response.ErrorResponse("it is not updated something error ", err.Error(), nil)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "succefully updated the cart", ResponseCart)
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, res)
}

// @Summary DeleteCart for user
// @ID UserDeleteCart for user
// @Tags USERCARTMANAGEMENT
// @Security BearerAuth
// @Param product_id query string true "product_id"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/delete-cart [delete]
func (cr *UserHandler) DeleteCart(c *gin.Context) {
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	cartitem, err := cr.UserService.FindProduct(uint(product_id))
	if err != nil {
		res := response.ErrorResponse("failed to find  the product ", err.Error(), "failed to find product")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	err = cr.UserService.DeleteCart(product_id, user_id)
	if err != nil {
		res := response.ErrorResponse("failed to delete the data", err.Error(), "failed  to remove from  cart please try again later")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "success fully cart item delted ", cartitem)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary Appy_Coupon for User
// @ID UserApplyCoupon for user
// @Tags APPLY_COUPON
// @Security BearerAuth
// @Param address_id query string true "address_id"
// @Param coupon query string true "coupon"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/apply-coupon [post]
func (cr *UserHandler) Apply_Coupon(c *gin.Context) {
	var cpn string
	var total int
	var msg string
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	address_id, _ := strconv.Atoi(c.Query("address_id"))
	coupon := c.Query("coupon")
	var Coupons domain.Coupon
	//find CartValue
	cartvalue, err := cr.UserService.FindTheSumOfCart(user_id)
	if err != nil {
		res := response.ErrorResponse("error when finding sum of the cart", err.Error(), total)
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	if coupon == "" {
		res := response.ErrorResponse("please enter a coupon if you have", "coupon field is  empty", "Ener a coupon to get offer")
		c.Writer.WriteHeader(300)
		utils.ResponseJSON(c, res)
	} else {
		Coupons, err = cr.UserService.FindCoupon(coupon)
		if err != nil {
			res := response.ErrorResponse("This coupon is not available right now", err.Error(), coupon)
			c.Writer.WriteHeader(422)
			utils.ResponseJSON(c, res)
			return
		}
		var time = time.Now().Unix()
		if Coupons.Validity > time {
			msg = fmt.Sprintf("coupen %d rupess applied ", Coupons.Discount)
			total = cartvalue - Coupons.Discount
			v := Coupons.Validity - time
			v = v / 3600
			cpn = fmt.Sprintf("coupon is valid user before %d hours", v)

		} else {
			cpn = "coupon is expired check your"
			res := response.ErrorResponse("coupon is expired", "enter another coupon", "or place order")
			c.Writer.WriteHeader(300)
			utils.ResponseJSON(c, res)
			msg = "coupon is not applied"
		}
	}
	data := struct {
		Coupon          domain.Coupon
		Validity_alert  string
		Address_Id      int
		User_Id         int
		Total_CartValue int
		Order_Total     int
	}{
		Coupon:          Coupons,
		Validity_alert:  cpn,
		Address_Id:      address_id,
		User_Id:         user_id,
		Total_CartValue: cartvalue,
		Order_Total:     total,
	}
	res := response.SuccessResponse(true, msg, data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary UserCheckout for User
// @ID UserCheckout for user
// @Tags USERCARTMANAGEMENT
// @Param address_id query string true "address_id"
// @Param payment_method query string true "payment_method"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /user/cart/checkout [post]
func (cr *UserHandler) Checkout(c *gin.Context) {
	//email := c.Writer.Header().Get("email")
	address_id, _ := strconv.Atoi(c.Query("address_id"))

	// user, _ := cr.AuthService.FindUser(email)
	// if err != nil {
	// 	res := response.ErrorResponse("failed to fetch the total value", err.Error(), "failedt to get the total price")
	// 	c.Writer.WriteHeader(400)
	// 	utils.ResponseJSON(c, res)
	// 	return
	// }
	payment_method := c.Query("payment_method")
	cr.Apply_Coupon(c)

	// cod := "cod"
	// razorpay := "razorpay"

	data := struct {
		Payment_Method string
		Address_Id     int
	}{
		Payment_Method: payment_method,
		Address_Id:     address_id,
	}
	res := response.SuccessResponse(true, "sucees fully checkout", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}
