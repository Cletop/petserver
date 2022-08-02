package database

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB
var GlobalRedis *redis.Client

func SetupDatabase() {
	db, err := ConnectPostgres()
	if err != nil {
		panic(err)
	}

	GlobalRedis = ConnectRedis()
	GlobalDB = db
}
