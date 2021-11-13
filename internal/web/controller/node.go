package controller

import (
	"monitorX/internal/database"
	"net/http"
)

func CreateData(data *database.NodeData, password string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	passwordStatus, err := mysqlConf.CheckNodePassword(data.NodeID, password)
	if err != nil {
		return returnServerError(err)
	}
	if passwordStatus {
		err = mysqlConf.CreateNodeData(data)
		if err != nil {
			return returnServerError(err)
		} else {
			responseBody.Msg = "Update success!"
		}
	} else {
		responseBody.Msg = "Password error!"
	}

	return &responseBody
}

func GetNodeData(nodeID int, userName string, timePeriod string) *ResponseBody {
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

func CreateNodeInfo(newNode *database.NodeInfo, userName string) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()
	if newNode.OwnedUserID == 0 {
		userInfo, err := mysqlConf.GetUser(userName)
		if err != nil {
			return returnServerError(err)
		}

		newNode.OwnedUserID = userInfo.ID
	}

	err := mysqlConf.CreateNodeInfo(newNode)
	if err != nil {
		return returnServerError(err)
	} else {
		responseBody.Msg = "Node created successfully"
	}

	return &responseBody
}

func UpdateNodeInfo(node *database.NodeInfo) *ResponseBody {
	responseBody := ResponseBody{code: http.StatusOK}
	mysqlConf := database.GetConfig()

	err := mysqlConf.UpdateNodeInfo(node)
	if err != nil {
		return returnServerError(err)
	} else {
		responseBody.Msg = "Node update successfully"
	}

	return &responseBody
}
