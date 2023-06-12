package Router

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Api"
)

func PlantRoutes(router *gin.Engine) {
	plantRounter := router.Group("/plant")
	{
		plantRounter.GET("/", Api.GetPlants)
		plantRounter.POST("/", Api.NewPlant)
		plantRounter.GET("/:id", Api.GetPlant)
		plantRounter.POST("/nutrients", Api.AddNutrient)
	}
}
