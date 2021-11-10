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
		return returnServerError(err)
	}
	if passwordStatus {
		err = mysqlConf.InsertNodeData(data)
		if err != nil {
			return returnServerError(err)
		} else {
			responseBody.Msg = "Update success!"
		}
	}

	return &responseBody
}

func QueryNodeData(nodeID int, userName string, timePeriod string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	checkStatus, err := mysqlConf.CheckNodeOwner(nodeID, userName)
	if err != nil {
		return returnServerError(err)
	}
	if checkStatus == false {
		return returnServerError(err)
	} else {
		nodeData, err := mysqlConf.GetNodeData(nodeID, timePeriod)
		if err != nil {
			return returnServerError(err)
		}
		responseBody.Data = nodeData
	}
	return &responseBody
}
