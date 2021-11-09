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
		return returnMysqlError(err)
	}
	if passwordStatus {
		err = mysqlConf.InsertNodeData(data)
		if err != nil {
			return returnMysqlError(err)
		} else {
			responseBody.Msg = "Update success!"
		}
	}

	return &responseBody
}

func QueryNodeData(nodeID int, userName string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	checkStatus, err := mysqlConf.CheckNodeOwner(nodeID, userName)
	if err != nil {
		return returnMysqlError(err)
	}
	if checkStatus == false {
		return returnMysqlError(err)
	} else {
		nodeData, err := mysqlConf.GetNodeData(nodeID)
		if err != nil {
			return returnMysqlError(err)
		}
		responseBody.Data = nodeData
	}
	return &responseBody
}
