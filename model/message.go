package model

import "gorm.io/gorm"

type MessageModel struct {
	gorm.Model
}

func (m *MessageModel) TableName() string {
	return "message"
}
