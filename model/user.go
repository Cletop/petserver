package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (UserModel) TableName() string {
	return "users"
}
