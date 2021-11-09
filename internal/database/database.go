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
