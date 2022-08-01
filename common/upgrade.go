package common

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WebsocketUpgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
