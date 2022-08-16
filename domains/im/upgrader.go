package im

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WebsocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	// },
	EnableCompression: true,
}

func UpgradeWebsocketProtocol(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := WebsocketUpgrade.Upgrade(w, r, nil)
}
