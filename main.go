package main

import (
	"github.com/chagspace/petserver/database"
	"github.com/chagspace/petserver/model"
	"github.com/chagspace/petserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// setup database
	database.SetupDatabase()

	// register multiple auto migrate
	model.RegisterMultipleAutoMigrate()

	// init routes
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes()
}
