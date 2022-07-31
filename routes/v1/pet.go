package routes

import (
	"github.com/chagspace/petserver/controller"
	"github.com/gin-gonic/gin"
)

func InitPetRouter(pet_router *gin.RouterGroup) {
	pet_router.GET("/pets", controller.GetPets)
	pet_router.GET("/pets/:id", controller.GetPet)
	pet_router.POST("/pets", controller.CreatePet)
	pet_router.PUT("/pets/:id", controller.UpdatePet)
	pet_router.DELETE("/pets/:id", controller.DeletePet)
}
