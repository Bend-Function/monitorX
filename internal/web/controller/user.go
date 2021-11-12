package controller

import (
	"monitorX/internal/database"
	"net/http"
)

func UserInfo(requestUserName string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	userInfo, err := mysqlConf.GetUser(requestUserName)
	if err != nil {
		responseBody.code = http.StatusBadRequest
		responseBody.Msg = err.Error()
		return &responseBody
	}
	userInfo.Password = ""
	responseBody.Data = userInfo
	return &responseBody
}

func UserNode(requestUserName string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	nodeList, err := mysqlConf.GetUserNode(requestUserName)
	if err != nil {
		responseBody.code = http.StatusBadRequest
		responseBody.Msg = err.Error()
		return &responseBody
	}
	responseBody.Data = nodeList
	return &responseBody
}

func CreateUser(newUser *database.User) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()

	checkUserName, err := mysqlConf.GetUser(newUser.UserName)
	if err != nil {
		return returnServerError(err)
	}
	if checkUserName.ID == 0 {
		responseBody.code = http.StatusBadRequest
		responseBody.Msg = "Already has same user name"
		return &responseBody
	}
	err = mysqlConf.CreateUser(newUser)
	if err != nil {
		return returnServerError(err)
	}

	return &responseBody
}
