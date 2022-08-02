package controller

import (
	"net/http"

	"github.com/chagspace/petserver/common"
	"github.com/chagspace/petserver/model"
	"github.com/chagspace/petserver/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {}
func GetUser(c *gin.Context)  {}

func CreateUser(c *gin.Context) {
	user := &model.UserModel{}
	c.BindJSON(&user)

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	service.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"msg":      "success",
		"username": user.Username,
		"userId":   user.ID,
		"email":    user.Email,
		"uid":      user.UID,
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"message": "update_user",
	})
}
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"message": "delete_user",
	})
}

// subscribe a user
func SubscribeUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"message": "subscribe_user",
	})
}
func UnsubscribeUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"message": "unsubscribe_user",
	})
}

// notify  a user
func NotifyUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"message": "notify_user",
	})
}

// Login user
func Login(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	database_user, allowed_user := service.GetUser(user.Username)
	password_error := bcrypt.CompareHashAndPassword([]byte(database_user.Password), []byte(user.Password))
	if password_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "password error"})
		return
	}
	if !allowed_user {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "unauthorized"})
		return
	}
	token, err := common.CreateToken(database_user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"msg":      "success",
		"token":    token,
		"username": user.Username,
		"userId":   database_user.ID,
		"uid":      database_user.UID,
	})
}

func Logout(c *gin.Context) {}
