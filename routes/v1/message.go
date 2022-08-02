package routes

import (
	"github.com/chagspace/petserver/controller"
	"github.com/gin-gonic/gin"
)

func InitMessageRouter(message_router *gin.RouterGroup) {
	message_router.GET("messages/upgrade", controller.Upgrade) // 升级WebSocket
	message_router.GET("messages/sse", controller.SSE)         // Server-Sent Events
	message_router.GET("/messages", controller.GetMessages)
	message_router.POST("/messages/publish", controller.PublishMessage)
}
