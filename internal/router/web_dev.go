//go:build dev

package router

import (
	gateway "github.com/Mmx233/Gateway/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

func frontendHandler() gin.HandlerFunc {
	return gateway.Proxy(&gateway.ApiConf{
		Addr: "localhost:5173",
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second * 30,
			}).DialContext,
		},
		ErrorHandler: func(_ http.ResponseWriter, _ *http.Request, err error) {
			log.Warnf("调试页面请求转发失败: %v", err)
		},
		AllowRequest: func(c *gin.Context) bool {
			return frontendRouterCheck(c)
		},
	})
}
