package routes

import (
	"fmt"

	"github.com/chagspace/petserver/middleware"
	"github.com/chagspace/petserver/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.New()

	// Filter Backlist Ip
	router.Use(middleware.Backlist())
	// Cors
	router.Use(middleware.Cors())
	// Load TLS
	router.Use(middleware.LoadTLS())

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1_router := router.Group("/api/v1")
	v1_router.Use(middleware.JWTAuth())
	{
		// account
		routes.InitUserRouter(v1_router)

		// business
		routes.InitPetRouter(v1_router)
		routes.InitMessageRouter(v1_router)
		routes.InitFileRouter(v1_router)
	}

	startApp(router)
}

func startApp(ginEngine *gin.Engine) {
	fmt.Printf("Starting server...: %s\n", "https://127.0.0.1:8080")
	// ginEngine.Run() // gin default 8080 port
	err := ginEngine.RunTLS(":8080", "./config/certificates/ca.crt", "./config/certificates/ca.key")
	panic(err)
}
