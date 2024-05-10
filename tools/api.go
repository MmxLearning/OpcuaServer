package tools

import (
	"github.com/MmxLearning/OpcuaServer/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	keyUserClaims = "api-user-claims"
)

func SetUserClaims(c *gin.Context, value *jwt.Claims) {
	c.Set(keyUserClaims, value)
}
func GetUserClaims(c *gin.Context) *jwt.Claims {
	return c.MustGet(keyUserClaims).(*jwt.Claims)
}
