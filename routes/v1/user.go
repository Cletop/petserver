package routes

import (
	"github.com/chagspace/petserver/controller"
	"github.com/chagspace/petserver/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(user_router *gin.RouterGroup) {
	user_router.GET("/users", controller.GetUsers)
	user_router.GET("/users/:id", middleware.RequestedSelfGet("uid"), controller.GetUser)

	user_router.POST("/users", controller.CreateUser)
	user_router.PUT("/users/:id", middleware.RequestedSelfGet("uid"), controller.UpdateUser)
	user_router.DELETE("/users/:id", middleware.RequestedSelfGet("uid"), controller.DeleteUser)

	user_router.POST("/users/subscribe/:id", middleware.RequestedSelfGet("uid"), controller.SubscribeUser)
	user_router.POST("/users/unsubscribe/:id", middleware.RequestedSelfGet("uid"), controller.UnsubscribeUser)

	user_router.POST("/users/notify", controller.NotifyUser)

	user_router.POST("/users/login", controller.Login)
	user_router.POST("/users/logout", controller.Logout)
}
