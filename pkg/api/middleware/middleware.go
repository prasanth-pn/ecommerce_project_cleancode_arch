package middleware

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	services "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/usecase/interfaces"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthorizeJWT(*gin.Context)
}
type middleware struct {
	jwtUseCase services.JWTService
}

func NewUserMiddleware(jwtUseCase services.JWTService) Middleware {
	return &middleware{
		jwtUseCase: jwtUseCase,
	}
}
func (cr *middleware) AuthorizeJWT(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	//fmt.Println(len(authHeader), authHeader)
	auth := strings.Join(authHeader, " ")
	//fmt.Println(len(auth), auth) 
	bearerToken := strings.Split(auth, " ")
	//fmt.Println(len(bearerToken), bearerToken)

	if len(bearerToken) != 2 {
		err := errors.New("request does not contain an access token")
		respons := response.ErrorResponse("failed to create user", err.Error(), nil)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		c.Abort()
		return
	}
	authtoken := bearerToken[1]
	ok, claims := cr.jwtUseCase.VerifyToken(authtoken)
	if !ok {
		err := errors.New("your token is not valid ")
		respons := response.ErrorResponse("failed to authenticate user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, respons)
		c.Abort()
		return
	}
	user_email := fmt.Sprint(claims.Email)
	id := fmt.Sprint(claims.User_Id)
	c.Writer.Header().Set("id", id)
	c.Writer.Header().Set("email", user_email)
}
