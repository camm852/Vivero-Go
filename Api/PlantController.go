package Api

import (
	"fmt"
	"net/http"
	dto "proyecto.com/Domain/Dto"

	"github.com/gin-gonic/gin"
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

	plant, _ := Services.GetPlant(id)

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

func AddNutrient(context *gin.Context) {
	//Getting Data
	if func() bool {
		id := context.PostForm("id")
		amount := context.PostForm("amount")
		return (id == "" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	plantNutrientDTO := dto.PlantSupplyDTO{}

	if err := context.ShouldBind(&plantNutrientDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	id := plantNutrientDTO.Id
	if id < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	amountNutrient := plantNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = Services.AddNutrients(plant, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}
