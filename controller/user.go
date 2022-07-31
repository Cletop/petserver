package controller

import (
	"github.com/chagspace/petserver/model"
	"github.com/chagspace/petserver/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {}
func GetUser(c *gin.Context)  {}
func CreateUser(c *gin.Context) {
	user := &model.UserModel{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email:    c.PostForm("email"),
	}
	service.CreateUser(user)
}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}

// subscribe a user
func SubscribeUser(c *gin.Context)   {}
func UnsubscribeUser(c *gin.Context) {}

// notify  a user
func NotifyUser(c *gin.Context) {}
