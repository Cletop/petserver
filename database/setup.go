package database

import (
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func SetupDatabase() {
	db, err := ConnectPostgres()
	if err != nil {
		panic(err)
	}
	GlobalDB = db
}
