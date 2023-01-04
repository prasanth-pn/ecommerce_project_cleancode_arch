package handler

import (
	//"net/http"
	//domain "clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
// 	response "clean/pkg/common/response"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/copier"
 )



type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase)*UserHandler{
	return &UserHandler{
		userUseCase:usecase,
	}

}