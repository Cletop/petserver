package model

import "github.com/chagspace/petserver/database"

func RegisterMultipleAutoMigrate() {
	db := database.GlobalDB
	db.AutoMigrate(
		&UserModel{},
		&PetModel{},
		&MessageModel{},
		// other models...
	)
}
