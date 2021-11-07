package database

import (
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
	userInfo, err := mysqlConf.GetUserByName("t2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)
}
