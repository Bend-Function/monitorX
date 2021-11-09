package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Config = &MysqlConfig{}
)

func (mysqlConf *MysqlConfig) GetDB() error {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local", mysqlConf.Username, mysqlConf.Password, mysqlConf.ServerAddr, mysqlConf.ServerPort, mysqlConf.Database)
	db, err := gorm.Open(mysql.Open(conn))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mysqlConf.MysqlConn = db
	return nil
}

func GetConfig() *MysqlConfig {
	return Config
}

func (mysqlConf *MysqlConfig) GetNodeInfo(nodeID int) (nodeInfo *NodeInfo, err error) {
	nodeInfo = new(NodeInfo)
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return nil, err
		}
	}
	mysqlConf.MysqlConn.Find(&nodeInfo, "id=?", nodeID)
	return nodeInfo, nil
}

func (mysqlConf *MysqlConfig) GetUserNode(name string) (nodeList *[]NodeInfo, err error) {
	nodeList = new([]NodeInfo)
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return nil, err
		}
	}
	mysqlConf.MysqlConn.Table("node_info").Select("*").Joins("INNER JOIN `user` ON node_info.owned_user_id = `user`.id WHERE `user`.user_name = ?", name).Scan(&nodeList)
	// SELECT node_info.u* FROM node_info INNER JOIN `user` ON node_info.owned_user_id = `user`.id WHERE `user`.user_name = name

	return nodeList, nil
}

func (mysqlConf *MysqlConfig) GetUser(name string) (userInfo *User, err error) {
	userInfo = new(User)
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return nil, err
		}
	}
	mysqlConf.MysqlConn.Find(&userInfo, "user_name=?", name)

	return userInfo, nil
}

func (mysqlConf *MysqlConfig) CheckNodePassword(nodeID int, password string) (status bool, err error) {
	nodeInfo, err := mysqlConf.GetNodeInfo(nodeID)
	if err != nil {
		return false, err
	}

	if nodeInfo.Password == password {
		return true, nil
	} else {
		return false, nil
	}

}

func (mysqlConf *MysqlConfig) InsertNodeData(nodeData *NodeData) (err error) {
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return err
		}
	}

	mysqlConf.MysqlConn.Create(&nodeData)
	return nil
}
