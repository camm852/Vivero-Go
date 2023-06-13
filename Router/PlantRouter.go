package Router

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Api"
)

func PlantRoutes(router *gin.Engine) {
	plantRouter := router.Group("/api/plant")
	{
		plantRouter.GET("/", Api.GetPlants)
		plantRouter.POST("/", Api.NewPlant)
		plantRouter.PUT("/", Api.UpdatePlant)
		plantRouter.GET("/:id", Api.GetPlant)
		plantRouter.DELETE("/:id", Api.DeletePlant)
		plantRouter.POST("/nutrients/add", Api.AddNutrient)
		plantRouter.POST("/nutrients/add-many", Api.AddManyNutrient)
		plantRouter.POST("/water/add", Api.AddWater)
		plantRouter.POST("/water/add-many", Api.AddManyWater)
	}
}
