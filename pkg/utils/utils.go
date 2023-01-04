package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(c.Writer).Encode(data)
	if err != nil {
		return
	}

}

// respons:=response.ErrorResponse("failed to login",err.Error(),nil)
// // c.Header().Add("Content-Type","application/json")
// c.Writer.Header().Set("Content-Type","application/json")
// //w.WriteHeader(http.StatusUnauthorized)
// c.Writer.WriteHeader(http.StatusUnauthorized)
// fmt.Printf("\n\n%v",respons)
// utils.ResponseJSON(c,respons)
