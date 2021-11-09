package database

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
