package controllers

import (
	"fmt"
	pool "github.com/Mmx233/WsPool"
	"github.com/MmxLearning/OpcuaServer/internal/api/callback"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/rdpTable"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/wsPool"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

func GetRdpTable(c *gin.Context) {
	callback.Success(c, rdpTable.RdpList())
}

func RdpStream(c *gin.Context) {
	var f struct {
		Name string `json:"name" form:"name" binding:"required"`
	}
	if err := c.ShouldBind(&f); err != nil {
		callback.Error(c, callback.ErrForm, err)
		return
	}

	clientIdentity := c.ClientIP() + "-" + fmt.Sprint(time.Now().UnixNano())
	conn, err := wsPool.Rdp.NewConn(c, clientIdentity, nil)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	unregister, ok := rdpTable.ListenRegister(f.Name, clientIdentity, func(bytes []byte) {
		_ = conn.WriteMessage(websocket.BinaryMessage, bytes)
	})
	if !ok {
		callback.Error(c, callback.ErrNotExist)
		return
	}

	conn.OnClose = func(_ *pool.Conn) {
		unregister()
	}
}
