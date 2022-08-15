package model

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `json:"email"`
	UID      uint   `gorm:"column:uuid" json:"uid"`
}

func (UserModel) TableName() string {
	return "user"
}

func (u *UserModel) BeforeCreate(scope *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	scope.Statement.SetColumn("UID", node.Generate().Int64())
	return nil
}
