package service

import (
	"github.com/MmxLearning/OpcuaServer/internal/db/dao"
	"gorm.io/gorm"
)

var Opcua = OpcuaSrv{DB: dao.DB}

type OpcuaSrv struct {
	*gorm.DB
}

func (a OpcuaSrv) Begin() (OpcuaSrv, error) {
	a.DB = a.DB.Begin()
	return a, a.Error
}

func (a OpcuaSrv) Store(name, nodeID, data string, timestamp uint32) (*dao.Opcua, error) {
	model := dao.Opcua{
		Name:      name,
		NodeID:    nodeID,
		Data:      data,
		Timestamp: timestamp,
	}
	return &model, model.Insert(a.DB)
}
