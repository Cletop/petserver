package database

import (
	"fmt"

	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func SetupDatabase() {
	db, err := ConnectPostgres()
	fmt.Print("Connecting to Postgres...", db)
	if err != nil {
		panic(err)
	}
	GlobalDB = db
}
