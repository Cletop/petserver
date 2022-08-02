package controller

import (
	"log"
	"net/http"

	"github.com/chagspace/petserver/common"
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
		"userId":   user.ID,
	})
}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}

// subscribe a user
func SubscribeUser(c *gin.Context)   {}
func UnsubscribeUser(c *gin.Context) {}

// notify  a user
func NotifyUser(c *gin.Context) {}

// Login user
func Login(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	allowed_user := service.Login(user.Username, user.Password)
	if !allowed_user {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "unauthorized"})
		return
	}

	token, err := common.CreateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"msg":      "success",
		"token":    token,
		"username": user.Username,
		"userId":   user.ID,
	})
}

func Logout(c *gin.Context) {}
