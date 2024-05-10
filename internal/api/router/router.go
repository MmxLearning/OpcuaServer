package router

import "github.com/gin-gonic/gin"

func Router(G *gin.RouterGroup) {
	routerUser(G.Group("user"))
	routerClient(G.Group("client"))
}
