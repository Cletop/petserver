package controller

import (
	"log"
	"net/http"

	"github.com/chagspace/petserver/model"
	"github.com/chagspace/petserver/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {}
func GetUser(c *gin.Context)  {}

func CreateUser(c *gin.Context) {
	user := &model.UserModel{}
	c.BindJSON(&user)
	log.Printf("%v", user)

	service.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"msg":      "success",
		"username": user.Username,
	})
}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}

// subscribe a user
func SubscribeUser(c *gin.Context)   {}
func UnsubscribeUser(c *gin.Context) {}

// notify  a user
func NotifyUser(c *gin.Context) {}
