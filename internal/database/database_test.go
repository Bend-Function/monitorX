package database

import (
	"encoding/json"
	"fmt"
	"testing"
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
	userInfo, err := mysqlConf.GetUser("t2")
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
	userInfo, err := mysqlConf.GetNodeInfo(2)
	if err != nil {
		fmt.Println(err)
	}
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
		PingDelay:        "222",
		Connections:      "222",
	}
	err = mysqlConf.InsertNodeData(nodeData)
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
