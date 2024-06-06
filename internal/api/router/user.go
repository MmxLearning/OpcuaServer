package router

import (
	controllers "github.com/MmxLearning/OpcuaServer/internal/api/controllers/user"
	"github.com/MmxLearning/OpcuaServer/internal/api/middlewares"
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"github.com/gin-gonic/gin"
)

func routerUser(G *gin.RouterGroup) {
	G.Use(middlewares.UserAuth([]byte(global.Config.JwtKey)))

	opcua := G.Group("opcua")
	opcua.GET("search", controllers.Search)

	rdp := G.Group("rdp")
	rdp.GET("/", controllers.GetRdpTable)
	rdp.GET("stream", controllers.RdpStream)
}
