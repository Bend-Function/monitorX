package database

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

func (mysqlConf *MysqlConfig) GetNodeData(nodeID int) (dataList *[]NodeData, err error) {
	dataList = new([]NodeData)
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return nil, err
		}
	}
	mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ?", nodeID).Scan(&dataList)

	return dataList, nil
}

func (mysqlConf *MysqlConfig) CheckNodeOwner(nodeID int, userName string) (status bool, err error) {
	if mysqlConf.MysqlConn == nil {
		err = mysqlConf.GetDB()
		if err != nil {
			return false, err
		}
	}
	userNodeList, err := mysqlConf.GetUserNode(userName)
	if err != nil {
		return false, err
	}
	for _, node := range *userNodeList {
		if node.ID == nodeID {
			return true, nil
		}
	}
	return false, err
}
