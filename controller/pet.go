package controller

import "github.com/gin-gonic/gin"

func CreatePet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create_pet",
	})
}
func UpdatePet(c *gin.Context) {
}
func DeletePet(c *gin.Context) {
}

func GetPet(c *gin.Context) {
}

func GetPets(c *gin.Context) {
}
