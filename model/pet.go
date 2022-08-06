package model

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type PetModel struct {
	ID string `json:"id"`

	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Weight   int    `json:"weight"`
	Height   int    `json:"height"`
	Breed    string `json:"breed"`
	Category string `json:"category"`

	Contactor string `json:"contactor"`
	User      string `json:"user"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Desc      string `json:"desc"`
	Status    string `json:"status"`

	UUID string `json:"uuid"`
}

func (PetModel) TableName() string {
	return "pet"
}

func (u *PetModel) BeforeCreate(scope *gorm.DB) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	scope.Statement.SetColumn("UID", node.Generate().Int64())
	return nil
}
