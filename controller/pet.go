package controller

import (
	"net/http"
	"strconv"

	"github.com/chagspace/petserver/common"
	"github.com/chagspace/petserver/model"
	"github.com/chagspace/petserver/service"
	"github.com/gin-gonic/gin"
)

func CreatePet(c *gin.Context) {
	var pet model.PetModel
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, common.StatusBadRequestMessage("invalid pet data"))
		return
	}
	service.CreatePet(pet)
	c.JSON(http.StatusOK, common.StatusOKMessage(gin.H{"pet": pet}, "create pet success"))
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
	pet_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.StatusBadRequestMessage("invalid pet id"))
		return
	}
	pet, is_record := service.GetPet(pet_id)
	if !is_record {
		c.JSON(http.StatusOK, common.StatusOKMessage(gin.H{"pet": nil}, "pet not found"))
		return
	}
	c.JSON(http.StatusOK, common.StatusOKMessage(gin.H{"pet": pet}, "get pet success"))
}

func GetPets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get_pets",
	})
}
