package router

import (
	"github.com/MmxLearning/OpcuaServer/internal/api/middlewares"
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"github.com/gin-gonic/gin"
)

func routerUser(G *gin.RouterGroup) {
	G.Use(middlewares.UserAuth([]byte(global.Config.JwtKey)))
}
