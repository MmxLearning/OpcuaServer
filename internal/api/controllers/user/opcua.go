package controllers

import (
	"github.com/MmxLearning/OpcuaServer/internal/api/callback"
	"github.com/MmxLearning/OpcuaServer/internal/service"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var f struct {
		Name    string `json:"name" form:"name"`
		NodeID  string `json:"nodeID" form:"nodeID"`
		StartAt int64  `json:"startAt" form:"startAt"`
		EndAt   int64  `json:"endAt" form:"endAt"`
	}
	if err := c.ShouldBind(&f); err != nil {
		callback.Error(c, callback.ErrForm, err)
		return
	}

	data, err := service.Opcua.Search(f.Name, f.NodeID, f.StartAt, f.EndAt)
	if err != nil {
		callback.Error(c, callback.ErrDBOperation, err)
		return
	}

	callback.Success(c, data)
}
