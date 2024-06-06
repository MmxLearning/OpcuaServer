package wsPool

import (
	pool "github.com/Mmx233/WsPool"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Rdp = pool.New(&websocket.Upgrader{
	HandshakeTimeout: time.Second * 30,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
})
