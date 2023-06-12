package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto.com/Application/Mappers"
	"proyecto.com/Application/Services"
	Entities "proyecto.com/Domain/Entities"
	"proyecto.com/Utils"
)

func GetPlants(context *gin.Context) {
	fmt.Println("HOla mundo")
	var plants []Entities.Plant = Services.GetPlants()

	context.IndentedJSON(http.StatusOK, plants)
}

func GetPlant(context *gin.Context) {
	var id uint = Utils.ParseUint(context.Param("id"))

	var plant Entities.Plant = Services.GetPlant(id)

	context.IndentedJSON(http.StatusOK, plant)
}

func NewPlant(context *gin.Context) {
	var plantDto Entities.PlantDTO
	if err := context.ShouldBindJSON(&plantDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plant, err := Mappers.MapPlantDtoToPlant(plantDto)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al mapear los datos"})
		return
	}

	var isCreated = Services.NewPlant(plant)

	if !isCreated {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la planta"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "Planta creada correctamente"})

}
