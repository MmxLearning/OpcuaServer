package router

import (
	"github.com/MmxLearning/OpcuaServer/internal/api/middlewares"
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"github.com/gin-gonic/gin"
)

func routerClient(G *gin.RouterGroup) {
	G.Use(middlewares.TokenAuth(global.Config.ClientToken))
}
