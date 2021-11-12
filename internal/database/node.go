package database

import "errors"

func (mysqlConf *MysqlConfig) GetNodeInfo(nodeID int) (nodeInfo *NodeInfo, err error) {
	nodeInfo = new(NodeInfo)
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}
	mysqlConf.MysqlConn.Find(nodeInfo, "id=?", nodeID)
	if nodeInfo.ID != nodeID {
		return nil, errors.New("can't find node")
	}
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

func (mysqlConf *MysqlConfig) CreateNodeData(nodeData *NodeData) (err error) {
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}

	//mysqlConf.MysqlConn.Create(&nodeData)
	mysqlConf.MysqlConn.Omit("ID", "UpdateTime").Create(&nodeData)
	return nil

}

func (mysqlConf *MysqlConfig) CreateNodeInfo(newNode *NodeInfo) (err error) {
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}
	newNode.IsInfoExpired = 1
	mysqlConf.MysqlConn.Omit("ID", "CreateTime", "UpdateTime").Create(&newNode)
	return
}

func (mysqlConf *MysqlConfig) UpdateNodeInfo(newNode *NodeInfo) (err error) {
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}
	//db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
	mysqlConf.MysqlConn.Omit("ID", "Password", "GroupID", "UpdateFrequency", "CreateTime", "UpdateTime").Updates(&newNode)
	return
}

func (mysqlConf *MysqlConfig) GetNodeGroup(nodeGroupID int) (nodeGroup *NodeGroup, err error) {
	nodeGroup = new(NodeGroup)
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}

	mysqlConf.MysqlConn.Find(&nodeGroup, "id=?", nodeGroupID)

	return

}

func (mysqlConf *MysqlConfig) GetNodeData(nodeID int, timePeriod string) (dataList *[]NodeData, err error) {
	dataList = new([]NodeData)
	if mysqlConf.checkMysqlConnection() != nil {
		err = errors.New("mysql connection error")
		return
	}
	switch timePeriod {
	case "today":
		mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ? AND to_days(`node_data`.update_time) = to_days(now())", nodeID).Scan(&dataList)
	case "yesterday":
		mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ? AND to_days(NOW()) - TO_DAYS(`node_data`.update_time) = 1", nodeID).Scan(&dataList)
	case "weekly":
		mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ? AND DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(`node_data`.update_time)", nodeID).Scan(&dataList)
	case "monthly":
		mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ? AND DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(`node_data`.update_time)", nodeID).Scan(&dataList)
	default:
		// Today
		mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ? AND to_days(`node_data`.update_time) = to_days(now())", nodeID).Scan(&dataList)
	}
	//mysqlConf.MysqlConn.Table("node_data").Select("*").Joins("INNER JOIN `node_info` ON node_info.id = `node_data`.node_id WHERE `node_data`.node_id = ?", nodeID).Scan(&dataList)

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
