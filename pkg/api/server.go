package http

import (
	handler "clean/pkg/api/handler"

	//middleware 		"clean/pkg/api/middleware"
	//"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "clean/cmd/api/docs"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(UserHandler *handler.UserHandler,AuthHandler *handler.AuthHandler) *ServerHTTP {
	engine := gin.New()

	//engine logger from gin

	engine.Use(gin.Logger())

	//swagger docs
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//request jwt

	// engine.POST("/login", middleware.LoginHandler)
	engine.POST("/admin/register",AuthHandler.AdminRegister)
	engine.POST("/admin/login",AuthHandler.AdminLogin)
	engine.POST("/user/register", AuthHandler.Register)
	engine.POST("/user/login",AuthHandler.UserLogin)
	//api := engine.Group("/api", middleware.AuthorizationMiddleware)

	// api.GET("list/users", UserHandler.FindAll)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":8080")

}
