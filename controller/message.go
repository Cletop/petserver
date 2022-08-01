package controller

import (
	"github.com/chagspace/petserver/common"
	"github.com/gin-gonic/gin"
)

func Upgrade(c *gin.Context) {
	// TODO：
	ws, err := common.WebsocketUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			return
		}

		// ping-pong message
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			return
		}
	}
}

// 从云端获取消息
func GetMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get_messages",
	})
}

// 像订阅者推送一条消息
func PublishMessage(c *gin.Context) {
	// TODO：
}
