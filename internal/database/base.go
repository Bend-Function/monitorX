package database

import (
	"database/sql"
	"time"
)

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
	Balance  float32
	GroupID  uint
}

type NodeInfo struct {
	ID              uint      `db:"id" json:"id"`
	NodeName        string    `db:"node_name" json:"node_name"`
	Password        string    `db:"passwd" json:"password"`
	GroupID         uint      `db:"group_id" json:"group_id"`
	OwnedUserID     uint      `db:"owned_user_id" json:"owned_user_id"`
	UpdateFrequency uint      `db:"update_frequency" json:"update_frequency"`
	NodeSystem      string    `db:"node_system" json:"node_system"`
	CoreVersion     string    `db:"core_version" json:"core_version"`
	StartupTime     time.Time `db:"startup_time" json:"startup_time"`
	CpuType         string    `db:"cpu_type" json:"cpu_type"`
	MemorySize      float32
	DiskSize        float32
	NetUpSum        int
	NetDownSum      int
	IsExpired       bool
	CreationTime    time.Time
	UpdateTime      time.Time
}

type NodeData struct {
	ID           uint
	NodeID       uint
	CpuUsage     float32
	MemoryUsage  float32
	DiskUsage    float32
	NetUpSpeed   float32
	NetDownSpeed float32
	PingDelay    float32
	Conn         uint
}
