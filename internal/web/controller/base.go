package controller

import (
	"fmt"
	"net/http"
)

type ResponseBody struct {
	code uint
	Data interface{}
	Msg  string
}

func returnServerError(err error) (responseBody *ResponseBody) {
	responseBody.code = http.StatusInternalServerError
	responseBody.Msg = fmt.Sprintf("mysql server error %v", err.Error())
	return responseBody
}
