package database

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var (
	mysqlConf = &MysqlConfig{
		ServerAddr: "127.0.0.1",
		ServerPort: 3306,
		Database:   "node",
		Username:   "root",
		Password:   "nodemonitor",
	}
)

func TestQueryUser(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	userInfo, err := mysqlConf.GetUser("t211")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)
}

func TestQueryNodeInfo(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	userInfo, err := mysqlConf.GetNodeInfo(1)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	fmt.Println(userInfo)
}

func TestQueryUserNodes(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	userInfo, err := mysqlConf.GetUserNode("func")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)
}

func TestInsertNodeData(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	nodeData := &NodeData{
		NodeID:           1,
		CpuUsage:         "50",
		MemoryUsage:      "24",
		DiskUsage:        "32",
		NetworkUpSpeed:   "4",
		NetworkDownSpeed: "5",
		//PingDelay:        "222",
		Connections: "222",
	}
	err = mysqlConf.CreateNodeData(nodeData)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(nodeData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func TestCheckNodePassword(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	status, err := mysqlConf.CheckNodePassword(1, "Aa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(status)
}

func TestQueryNodeData(t *testing.T) {
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	nodeData, err := mysqlConf.GetNodeData(1, "yesterday")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nodeData)
}

func TestCreateUser(t *testing.T) {
	var newUser = &User{
		UserName: "t31",
		Password: "1234",
		Email:    "1@q.cn",
		Balance:  "3",
		GroupID:  1,
	}
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	err = mysqlConf.CreateUser(newUser)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUpdateNodeInfo(t *testing.T) {
	var nodeInfo = &NodeInfo{
		ID:             6,
		NodeSystem:     "C",
		CoreVersion:    "C",
		CpuType:        "A",
		MemorySize:     "A",
		DiskSize:       "A",
		StartupTime:    time.Now(),
		NetworkUpSum:   "A",
		NetworkDownSum: "A",
		IsInfoExpired:  0,
	}
	err := mysqlConf.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	err = mysqlConf.UpdateNodeInfo(nodeInfo)
	if err != nil {
		fmt.Println(err)
	}
}
