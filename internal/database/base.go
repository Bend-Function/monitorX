package database

import (
	"gorm.io/gorm"
	"time"
)

type MysqlConfig struct {
	ServerAddr string `json:"server_addr"`
	ServerPort int    `json:"server_port"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	MysqlConn  *gorm.DB
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
	ID         int       `gorm:"column:id" json:"id"`
	UserName   string    `gorm:"column:user_name" json:"user_name" form:"user_name"`
	Password   string    `gorm:"column:passwd" json:"password" form:"password"`
	Email      string    `gorm:"column:email" json:"email" form:"email"`
	Balance    string    `gorm:"column:balance" json:"balance" from:"balance"`
	GroupID    int       `gorm:"column:group_id" json:"group_id" form:"group_id"` // 1 - means admin; > 1 means user
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

type UserGroup struct {
	ID         int       `gorm:"column:id" json:"id"`
	GroupName  string    `gorm:"column:group_name" json:"group_name"`
	Level      int       `gorm:"column:level" json:"level"`
	Comment    string    `gorm:"column:comment" json:"comment"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

type NodeInfo struct {
	ID              int       `gorm:"column:id" json:"id" form:"id"`
	NodeName        string    `gorm:"column:node_name" json:"node_name" form:"node_name" binding:"required"`
	Password        string    `gorm:"column:passwd" json:"password" form:"password"`
	GroupID         int       `gorm:"column:group_id" json:"group_id" form:"group_id"`
	OwnedUserID     int       `gorm:"column:owned_user_id" json:"owned_user_id" form:"owned_user_id"`
	UpdateFrequency int       `gorm:"column:update_frequency" json:"update_frequency" form:"update_frequency"` // minute
	NodeSystem      string    `gorm:"column:node_system" json:"node_system" form:"node_system"`
	CoreVersion     string    `gorm:"column:core_version" json:"core_version" form:"core_version"`
	StartupTime     time.Time `gorm:"column:startup_time" json:"startup_time" form:"startup_time"`
	CpuType         string    `gorm:"column:cpu_type" json:"cpu_type" form:"cpu_type"`
	MemorySize      string    `gorm:"column:memory_size" json:"memory_size" form:"memory_size"`
	DiskSize        string    `gorm:"column:disk_size" json:"disk_size" form:"disk_size"`
	NetworkUpSum    string    `gorm:"column:network_up_sum" json:"network_up_sum" form:"network_up_sum"`
	NetworkDownSum  string    `gorm:"column:network_down_sum" json:"network_down_sum" form:"network_down_sum"`
	IsInfoExpired   int       `gorm:"column:is_info_expired" json:"is_info_expired" form:"is_info_expired"` // 1: need to update;
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
}

type NodeData struct {
	ID               int       `gorm:"column:id" json:"id"`
	NodeID           int       `gorm:"column:node_id" json:"node_id" form:"node_id" binding:"required"`
	CpuUsage         string    `gorm:"column:cpu_usage" json:"cpu_usage" form:"cpu_usage"`
	MemoryUsage      string    `gorm:"column:memory_usage" json:"memory_usage" form:"memory_usage"`
	DiskUsage        string    `gorm:"column:disk_usage" json:"disk_usage" form:"disk_usage"`
	NetworkUpSpeed   string    `gorm:"column:network_up_speed" json:"network_up_speed" form:"network_up_speed"`
	NetworkDownSpeed string    `gorm:"column:network_down_speed" json:"network_down_speed" form:"network_down_speed"`
	PingDelay        string    `gorm:"column:ping_delay" json:"ping_delay" form:"ping_delay"`
	Connections      string    `gorm:"column:connections" json:"connections" form:"connections"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
}

type NodeGroup struct {
	ID                int       `gorm:"column:id" json:"id"`
	NodeGroupName     string    `gorm:"column:node_group_name" json:"node_group_name"`
	NodeGroupPassword string    `gorm:"column:node_group_passwd" json:"node_group_password"`
	Level             string    `gorm:"column:level" json:"level"`
	OwnedUserID       int       `gorm:"column:owned_user_id" json:"owned_user_id"`
	Comment           string    `gorm:"column:comment" json:"comment"`
	ExpiredTime       time.Time `gorm:"column:expired_time" json:"expired_time"`
	CreatTime         time.Time `gorm:"column:creat_time" json:"creat_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
}

func (User) TableName() string {
	return "user"
}
func (UserGroup) TableName() string {
	return "user_group"
}
func (NodeInfo) TableName() string {
	return "node_info"
}
func (NodeData) TableName() string {
	return "node_data"
}
func (NodeGroup) TableName() string {
	return "node_group"
}

func (mysqlConf *MysqlConfig) checkMysqlConnection() (err error) {
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return err
		}
	}
	return
}
