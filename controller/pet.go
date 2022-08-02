package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create_pet",
	})
}
func UpdatePet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update_pet",
	})
}
func DeletePet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete_pet",
	})
}

func GetPet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get_pet",
	})
}

func GetPets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get_pets",
	})
}
