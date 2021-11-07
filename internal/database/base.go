package database

import "database/sql"

type MysqlConfig struct {
	ServerAddr string `json:"server_addr"`
	ServerPort int    `json:"server_port"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	MysqlConn  *sql.DB
}

type User struct {
	ID       uint
	Username string
	Password string
	Email    string
	balance  float32
	groupID  uint
}
