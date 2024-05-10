package router

import (
	api "github.com/MmxLearning/OpcuaServer/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	E := gin.Default()

	api.Router(E.Group("api"))

	serveFrontend(E)

	return E
}
