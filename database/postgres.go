package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=postgres dbname=petappdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		// NamingStrategy: gorm.NamingStrategy{ SingularTable: true }, // 单数表名 在v2版本中使用
	})
}
