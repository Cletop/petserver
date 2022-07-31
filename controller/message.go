package controller

import "github.com/gin-gonic/gin"

func GetMessages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get_messages",
	})
}

// 像订阅者推送一条消息
func PublishMessage(c *gin.Context) {
	// TODO：
}
