package main

import (
	"monitorX/internal/database"
	"monitorX/internal/web"
)

var (
	mysqlConf = &database.MysqlConfig{
		ServerAddr: "127.0.0.1",
		ServerPort: 3306,
		Database:   "node",
		Username:   "root",
		Password:   "nodemonitor",
	}
)

func main() {
	database.Config = mysqlConf
	web.Start("0.0.0.0", 80, 120)
}
