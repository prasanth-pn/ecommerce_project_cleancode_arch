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

	//swagger docs
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//request jwt

	userapi := engine.Group("user")
	{
		userapi.POST("/register", AuthHandler.Register)
		userapi.GET("/list/products", UserHandler.ListProducts)
		userapi.POST("/login", AuthHandler.UserLogin)
	}

	//------------------------------------user middleware-------------
	userapi.Use(middleware.AuthorizeJWT)
	userapi.POST("/add/cart", UserHandler.AddToCart)
	userapi.GET("/list/cart",UserHandler.ListCart)
	userapi.GET("/checkout",UserHandler.Checkout)
	userapi.POST("/checkout/add-address",UserHandler.AddAddress)


	//------------------------------admin----------------
	adminapi := engine.Group("admin")
	adminapi.POST("/register", AuthHandler.AdminRegister)
	adminapi.POST("/login", AuthHandler.AdminLogin)

	//---------------------------middleware checking------------------
	adminapi.Use(middleware.AuthorizeJWT)
	adminapi.GET("/list/users", AdminHandler.ListUsers)
	adminapi.POST("add/category", AdminHandler.AddCategory)
	adminapi.POST("/add/products", AdminHandler.AddProducts)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":8080")

}
