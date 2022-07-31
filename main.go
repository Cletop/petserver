package main

import (
	"github.com/chagspace/petserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// setup database
	// database.SetupDatabase()

	// init routes
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes()
}
