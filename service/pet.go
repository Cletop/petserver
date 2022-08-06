package service

import (
	"github.com/chagspace/petserver/database"
	"github.com/chagspace/petserver/model"
	"gorm.io/gorm"
)

func CreatePet(pet model.PetModel) {
	db := database.GlobalDB
	db.Create(&pet)
}

func GetPet(uid uint64) (model.PetModel, bool) {
	db := database.GlobalDB
	var pet model.PetModel
	err := db.Where("uuid = ?", uid).First(&pet).Error
	if err == gorm.ErrRecordNotFound {
		return model.PetModel{}, false
	}
	return pet, true
}

func UpdatePet() {
}

func DeletePet() {
}

func FollowPet() {
}

func UnFollowPet() {
}
