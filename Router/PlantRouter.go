package Router

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Api"
)

func PlantRoutes(router *gin.Engine) {
	plantRouter := router.Group("/plant")
	{
		plantRouter.GET("/", Api.GetPlants)
		plantRouter.POST("/", Api.NewPlant)
		plantRouter.GET("/:id", Api.GetPlant)
<<<<<<< HEAD
		plantRouter.POST("/nutrients/add/", Api.AddNutrient)
		plantRouter.POST("/nutrients/add-many/", Api.AddManyNutrient)
=======
		plantRouter.POST("/nutrients", Api.AddNutrient)
>>>>>>> 99e43e488abf42097312a29de2c69610f32bed12
	}
}
