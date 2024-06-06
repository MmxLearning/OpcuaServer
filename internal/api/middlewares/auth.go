package middlewares

import (
	"github.com/MmxLearning/OpcuaServer/internal/api/callback"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/jwt"
	"github.com/MmxLearning/OpcuaServer/tools"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func TokenAuth(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerSplit := strings.Split(header, " ")
		if len(headerSplit) == 2 && headerSplit[0] == "token" {
			tokenClient := headerSplit[1]
			if token == tokenClient {
				c.Next()
				return
			}
		}
		callback.Error(c, callback.ErrUnauthorized)
	}
}

func UserAuth(jwtKey []byte) gin.HandlerFunc {
	jwtSigner := jwt.New(jwtKey)
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			header = c.GetHeader("Sec-WebSocket-Protocol")
			if header != "" {
				var err error
				header, err = url.QueryUnescape(header)
				if err != nil {
					callback.Error(c, callback.ErrUnexpected, err)
					return
				}
			}
		}
		headerSplit := strings.Split(header, " ")
		if len(headerSplit) == 2 && headerSplit[0] == "user" {
			claims, err := jwtSigner.ParseToken(headerSplit[1])
			if err != nil {
				callback.Error(c, callback.ErrUnauthorized, err)
				return
			}
			tools.SetUserClaims(c, claims)
			c.Next()
			return
		}
		callback.Error(c, callback.ErrUnauthorized)
	}
}
