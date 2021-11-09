package controller

import "net/http"

type ResponseBody struct {
	code uint
	Data interface{}
	Msg  string
}

func returnMysqlError(err error) (responseBody *ResponseBody) {
	responseBody.code = http.StatusInternalServerError
	responseBody.Msg = err.Error()
	return responseBody
}
