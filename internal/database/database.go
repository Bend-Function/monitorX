package database

import (
	"database/sql"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

var (
	Config = &MysqlConfig{}
)

func (mysql *MysqlConfig) GetDB() error {
	// 屏蔽mysql驱动包的日志输出
	mysqlDriver.SetLogger(log.New(ioutil.Discard, "", 0))
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", mysql.Username, mysql.Password, mysql.ServerAddr, mysql.ServerPort, mysql.Database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mysql.MysqlConn = db
	return nil
}

func GetConfig() *MysqlConfig {
	return Config
}

func (mysql *MysqlConfig) GetNodeInfo(nodeID uint) (nodeInfo *NodeInfo, err error) {
	nodeInfo = new(NodeInfo)
	if mysql.MysqlConn == nil {
		err = mysql.GetDB()
		if err != nil {
			return nil, err
		}
	}
	row := mysql.MysqlConn.QueryRow(fmt.Sprintf("SELECT * FROM node_info WHERE BINARY id='%v'", nodeID))
	if err = row.Scan(&nodeInfo.ID, &nodeInfo.NodeName, &nodeInfo.Password, &nodeInfo.GroupID, &nodeInfo.OwnedUserID, &nodeInfo.UpdateFrequency, &nodeInfo.NodeSystem, &nodeInfo.CoreVersion, &nodeInfo.StartupTime, &nodeInfo.CpuType, &nodeInfo.MemorySize, &nodeInfo.DiskSize, &nodeInfo.NetUpSum, &nodeInfo.NetDownSum, &nodeInfo.IsExpired, &nodeInfo.CreationTime, &nodeInfo.UpdateTime); err != nil {
		return nil, err
	}
	return nodeInfo, nil
}

func (mysql *MysqlConfig) GetUserByName(name string) (userInfo *User, err error) {
	userInfo = new(User)
	if mysql.MysqlConn == nil {
		err = mysql.GetDB()
		if err != nil {
			return nil, err
		}
	}
	//userInfo, err = queryUser(mysql.MysqlConn, fmt.Sprintf("SELECT id, user_name, passwd, email, balance, group_id FROM user WHERE BINARY user_name='%s'", name) )
	row := mysql.MysqlConn.QueryRow(fmt.Sprintf("SELECT id, user_name, passwd, email, balance, group_id FROM user WHERE BINARY user_name='%s'", name))
	if err = row.Scan(&userInfo.ID, &userInfo.Username, &userInfo.Password, &userInfo.Email, &userInfo.Balance, &userInfo.GroupID); err != nil {
		return nil, err
	}
	return userInfo, nil
}
