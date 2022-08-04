package routes

import (
	"fmt"

	"github.com/chagspace/petserver/middleware"
	"github.com/chagspace/petserver/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.New()

	// Cors
	router.Use(middleware.Cors())
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1_router := router.Group("/api/v1")
	// v1_router.Use(AuthRequired())
	{
		// account
		routes.InitUserRouter(v1_router)

		// business
		routes.InitPetRouter(v1_router)
		routes.InitMessageRouter(v1_router)
	}

	startApp(router)
}

func startApp(ginEngine *gin.Engine) {
	fmt.Printf("Starting server...: %s\n", "http://localhost:8080")
	ginEngine.Run() // gin default 8080 port

}
