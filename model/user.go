package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
}

func (UserModel) TableName() string {
	return "users"
}
