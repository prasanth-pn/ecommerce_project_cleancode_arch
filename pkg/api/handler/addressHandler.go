package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (cr *UserHandler) AddAddress(c *gin.Context) {
	email := c.Writer.Header().Get("email")
	fmt.Println(email)
}
