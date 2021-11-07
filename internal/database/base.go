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

//type User struct {
//	ID       uint
//	Username string
//	Password string
//	Email    string
//	Balance  float32
//	GroupID  uint
//}

type User struct {
	ID         int       `db:"id" json:"id"`
	UserName   string    `db:"user_name" json:"user_name"`
	Password   string    `db:"passwd" json:"password"`
	Email      string    `db:"email" json:"email"`
	Balance    string    `db:"balance" json:"balance"`
	GroupID    int       `db:"group_id" json:"group_id"` // 1 - means admin; > 1 means user
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateTime time.Time `db:"update_time" json:"update_time"`
}

type UserGroup struct {
	ID         int    `db:"id" json:"id"`
	GroupName  string `db:"group_name" json:"group_name"`
	Level      int    `db:"level" json:"level"`
	Comment    string `db:"comment" json:"comment"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateTime time.Time`db:"update_time" json:"update_time"`
}
//type NodeInfo struct {
//	ID              uint      `db:"id" json:"id"`
//	NodeName        string    `db:"node_name" json:"node_name"`
//	Password        string    `db:"passwd" json:"password"`
//	GroupID         uint      `db:"group_id" json:"group_id"`
//	OwnedUserID     uint      `db:"owned_user_id" json:"owned_user_id"`
//	UpdateFrequency uint      `db:"update_frequency" json:"update_frequency"`
//	NodeSystem      string    `db:"node_system" json:"node_system"`
//	CoreVersion     string    `db:"core_version" json:"core_version"`
//	StartupTime     time.Time `db:"startup_time" json:"startup_time"`
//	CpuType         string    `db:"cpu_type" json:"cpu_type"`
//	MemorySize      float32
//	DiskSize        float32
//	NetUpSum        int
//	NetDownSum      int
//	IsExpired       bool
//	CreationTime    time.Time
//	UpdateTime      time.Time
//}

type NodeInfo struct {
	ID              int       `db:"id" json:"id"`
	NodeName        string    `db:"node_name" json:"node_name"`
	Password          string    `db:"passwd" json:"password"`
	GroupID         int       `db:"group_id" json:"group_id"`
	OwnedUserID     int       `db:"owned_user_id" json:"owned_user_id"`
	UpdateFrequency int       `db:"update_frequency" json:"update_frequency"` // minute
	NodeSystem      string    `db:"node_system" json:"node_system"`
	CoreVersion     string    `db:"core_version" json:"core_version"`
	StartupTime     time.Time `db:"startup_time" json:"startup_time"`
	CpuType         string    `db:"cpu_type" json:"cpu_type"`
	MemorySize      string    `db:"memory_size" json:"memory_size"`
	DiskSize        string    `db:"disk_size" json:"disk_size"`
	NetworkUpSum    string    `db:"network_up_sum" json:"network_up_sum"`
	NetworkDownSum  string    `db:"network_down_sum" json:"network_down_sum"`
	IsInfoExpired   int       `db:"is_info_expired" json:"is_info_expired"` // 1: need to update;
	CreatTime       time.Time `db:"creat_time" json:"creat_time"`
	UpdateTime      time.Time `db:"update_time" json:"update_time"`
}

//type NodeData struct {
//	ID           uint
//	NodeID       uint
//	CpuUsage     float32
//	MemoryUsage  float32
//	DiskUsage    float32
//	NetUpSpeed   float32
//	NetDownSpeed float32
//	PingDelay    float32
//	Conn         uint
//}

type NodeData struct {
	ID               int       `db:"id" json:"id"`
	NodeID           int       `db:"node_id" json:"node_id"`
	CpuUsage         string    `db:"cpu_usage" json:"cpu_usage"`
	MemoryUsage      string    `db:"memory_usage" json:"memory_usage"`
	DiskUsage        string    `db:"disk_usage" json:"disk_usage"`
	NetworkUpSpeed   string    `db:"network_up_speed" json:"network_up_speed"`
	NetworkDownSpeed string    `db:"network_down_speed" json:"network_down_speed"`
	PingDelay        string    `db:"ping_delay" json:"ping_delay"`
	Connections      string    `db:"connections" json:"connections"`
	UpdateTime       time.Time `db:"update_time" json:"update_time"`
}

type NodeGroup struct {
	ID              int       `db:"id" json:"id"`
	NodeGroupName   string    `db:"node_group_name" json:"node_group_name"`
	NodeGroupPassword string    `db:"node_group_passwd" json:"node_group_password"`
	Level           string    `db:"level" json:"level"`
	OwnedUserID     int       `db:"owned_user_id" json:"owned_user_id"`
	Comment         string    `db:"comment" json:"comment"`
	ExpiredTime     time.Time `db:"expired_time" json:"expired_time"`
	CreatTime       time.Time `db:"creat_time" json:"creat_time"`
	UpdateTime      time.Time `db:"update_time" json:"update_time"`
}


