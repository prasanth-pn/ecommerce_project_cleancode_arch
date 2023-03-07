package http

import (
	handler "clean/pkg/api/handler"
	"clean/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "clean/cmd/api/docs"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(UserHandler *handler.UserHandler,
	AuthHandler *handler.AuthHandler,
	AdminHandler *handler.AdminHandler,
	middleware middleware.Middleware,
) *ServerHTTP {
	engine := gin.New()

	//engine logger from gin

	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("templates/*.html")

	//swagger docs
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//request jwt

	userapi := engine.Group("user")
	{
		userapi.POST("/register", AuthHandler.Register)
		userapi.GET("/list/products", UserHandler.ListProducts)
		userapi.GET("/list/categories", UserHandler.ListCategories)
		userapi.GET("/list/products_bycategories", UserHandler.ListProductsByCategories)
		userapi.POST("/login", AuthHandler.UserLogin)
		userapi.POST("/send/verificationmail", AuthHandler.SendUserMail)
		userapi.GET("verify/otp", AuthHandler.VerifyUserOtp)
		userapi.GET("/order/razorpay", UserHandler.RazorPay)
		userapi.GET("/payment-success",UserHandler.Payment_Success)

	}

	//------------------------------------user middleware-------------
	userapi.Use(middleware.AuthorizeJWT)
	userapi.POST("/add/wishlist", UserHandler.AddWishList)
	userapi.GET("/view/wishlist", UserHandler.ViewWishList)
	userapi.DELETE("/delete/wishlist", UserHandler.RemoveFromWishlist)
	userapi.POST("/add/wishlist_tocart", UserHandler.WishListTo_Cart)
	userapi.POST("/add/cart", UserHandler.AddToCart)
	userapi.GET("/list/cart", UserHandler.ListCart)
	userapi.PATCH("/update/cart", UserHandler.UpdateCart)
	userapi.POST("/cart/checkout", UserHandler.Checkout)
	userapi.POST("/add/address", UserHandler.AddAddress)
	userapi.GET("/list/address", UserHandler.ListAddress)
	userapi.GET("/edit/address", UserHandler.GetAddressToEdit)
	userapi.PATCH("/update/address", UserHandler.UpdateAddress)
	userapi.GET("/list-order",UserHandler.ListOrder)

	//------------------------------admin----------------
	adminapi := engine.Group("admin")
	adminapi.POST("/register", AuthHandler.AdminRegister)
	adminapi.POST("/login", AuthHandler.AdminLogin)

	//---------------------------middleware checking------------------
	adminapi.Use(middleware.AuthorizeJWT)
	adminapi.PATCH("user/block", AdminHandler.BlockUser)
	adminapi.PATCH("/user/unblock", AdminHandler.UnblockUser)
	adminapi.GET("list/blockedusers", AdminHandler.ListBlockedUsers)
	adminapi.GET("/list/users", AdminHandler.ListUsers)
	adminapi.POST("add/category", AdminHandler.AddCategory)
	adminapi.POST("add/brands", AdminHandler.AddBrand)
	adminapi.POST("add/models", AdminHandler.AddModel)
	adminapi.POST("/add/products", AdminHandler.AddProducts)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":8080")

}
