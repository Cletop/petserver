package controller

import (
	"io"
	"net/http"

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

// server-sent events
func SSE(c *gin.Context) {
	chanStream := make(chan int, 2)
	go func() {
		defer close(chanStream)
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
		}
		return false
	})
}

// 从云端获取消息
func GetMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get_messages",
	})
}

// 像订阅者推送一条消息
func PublishMessage(c *gin.Context) {
	// TODO：
	c.JSON(http.StatusOK, gin.H{
		"message": "publish_message",
	})
}
