package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (db *gorm.DB, err error) {
	dsn := "postgresql://localhost:5432/petserver?sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
