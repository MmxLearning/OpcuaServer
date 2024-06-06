package controllers

import (
	"github.com/MmxLearning/OpcuaServer/internal/api/callback"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/rdpTable"
	"github.com/gin-gonic/gin"
)

func GetRdpTable(c *gin.Context) {
	callback.Success(c, rdpTable.RdpList())
}
