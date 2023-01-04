package response

import "strings"



type Response struct{
	Status bool  `json:"status"`
	Message string `json:"message"`
	Errors   interface{} `json:"errors,omitempty"`
	Data interface{}  `json:"data,omitempty"`
}


func SuccessResponse(status bool,message string,data interface{})Response{
	res:=Response{
		Status:status,
		Message:message,
		Errors:nil,
		Data:data,
	}
	return res
}

func ErrorResponse(message string,err string,data interface{})Response{
splittedError:=strings.Split(err,"\n")
res:=Response{
	Status:false,
	Message:message,
	Errors:splittedError,
	Data:data,
}
return res
}


