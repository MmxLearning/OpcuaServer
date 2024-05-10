package router

import (
	controllers "github.com/MmxLearning/OpcuaServer/internal/api/controllers/public"
	"github.com/gin-gonic/gin"
)

func routerPublic(G *gin.RouterGroup) {
	G.POST("login", controllers.Login)
}
