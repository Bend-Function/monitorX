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
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysql.Username, mysql.Password, mysql.ServerAddr, mysql.ServerPort, mysql.Database)
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

func queryUser(db *sql.DB, sql string) (*User, error) {
	userInfo := &User{}
	row := db.QueryRow(sql)
	if err := row.Scan(&userInfo.ID, &userInfo.Username, &userInfo.Password, &userInfo.Email, &userInfo.balance, &userInfo.groupID); err != nil {
		return nil, err
	}
	return userInfo, nil
}
func (mysql *MysqlConfig) GetUserByName(name string) (*User, error) {
	if mysql.MysqlConn == nil {
		err := mysql.GetDB()
		if err != nil {
			return nil, err
		}
	}
	user, err := queryUser(mysql.MysqlConn, fmt.Sprintf("SELECT id, user_name, passwd, email, balance, group_id FROM user WHERE BINARY user_name='%s'", name))
	if err != nil {
		return nil, err
	}
	return user, nil
}
