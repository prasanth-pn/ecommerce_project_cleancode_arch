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

	userapi := engine.Group("userapi")
	
	userapi.POST("/user/register", AuthHandler.Register)
	userapi.POST("/user/login",AuthHandler.UserLogin)


	userapi.Use(middleware.AuthorizeJWT)


	adminapi := engine.Group("adminapi")


	adminapi.POST("/admin/register",AuthHandler.AdminRegister)
	adminapi.POST("/admin/login",AuthHandler.AdminLogin)

	adminapi.Use(middleware.AuthorizeJWT)
	adminapi.GET("admin/list/users", AdminHandler.ListUsers)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":8080")

}
