package service

import (
	"github.com/chagspace/petserver/database"
	"github.com/chagspace/petserver/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.UserModel) {
	db := database.GlobalDB
	db.Create(user)
}

func UpdateUser(user *model.UserModel) {
	db := database.GlobalDB

	// 条件更新 username
	// db.Model(&model.UserModel{}).Where("UpdateAt = ?", true).Update("Username", user.Username)

	// 更新选定字段
	// db.Model(&user).Select("Email").Update("Email", user.Email)

	// 批量更新
	db.Model(&user).Updates(user)
}

func DeleteUser() {
	db := database.GlobalDB
	db.Delete(&model.UserModel{}, 20)             // 删除 id 为 20 的用户
	db.Where("id", 20).Delete(&model.UserModel{}) // 删除 id 为 20 的用户
}

func GetUser(username string) (*model.UserModel, bool) {
	db := database.GlobalDB
	user := &model.UserModel{}
	err := db.Where("username = ?", username).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}
	return user, true
}

func GetUserByUID(uid uint) (model.UserModel, bool) {
	db := database.GlobalDB
	var user model.UserModel
	err := db.Where("uuid = ?", uid).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return model.UserModel{}, false
	}

	return user, true
}
