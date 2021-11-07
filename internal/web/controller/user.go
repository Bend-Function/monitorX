package controller

import (
	"monitorX/internal/database"
	"net/http"
)

func UserInfo(requestUser string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	userInfo, err := mysqlConf.GetUserByName(requestUser)
	if err != nil {
		responseBody.code = http.StatusBadRequest
		responseBody.Msg = err.Error()
		return &responseBody
	}
	userInfo.Password = ""
	responseBody.Data = userInfo
	return &responseBody
}
