package controller

import (
	"monitorX/internal/database"
	"net/http"
)

func UpdateData(data *database.NodeData, password string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	passwordStatus, err := mysqlConf.CheckNodePassword(data.NodeID, password)
	if err != nil {
		responseBody.code = http.StatusBadRequest
		responseBody.Msg = err.Error()
		return &responseBody
	}
	if passwordStatus {
		err = mysqlConf.InsertNodeData(data)
		if err != nil {
			responseBody.code = http.StatusBadRequest
			responseBody.Msg = err.Error()
			return &responseBody
		} else {
			responseBody.Msg = "Update success!"
		}
	}

	return &responseBody
}
