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

	plant, error := Services.GetPlant(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "No se encontro la planta"})
		return
	}
	context.IndentedJSON(http.StatusOK, plant)
}

func NewPlant(context *gin.Context) {
	var newPlantDto dto.NewPlantDTO
	if err := context.ShouldBindJSON(&newPlantDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plant, err := Mappers.MapNewPlantDtoToPlant(newPlantDto)
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

func UpdatePlant(context *gin.Context) {
	var updatePlantDto dto.UpdatePlantDTO
	if err := context.ShouldBindJSON(&updatePlantDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plant, err := Mappers.MapUpdatePlantDtoToPlant(updatePlantDto)
	fmt.Println(plant)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al mapear los datos"})
		return
	}

	var isUpdated = Services.UpdatePlant(plant)

	if !isUpdated {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la planta"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "Planta actualizada correctamente"})
}

func DeletePlant(context *gin.Context) {
	var id uint = Utils.ParseUint(context.Param("id"))

	_, error := Services.GetPlant(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "No se encontro la planta"})
		return
	}

	isDeleted := Services.DeletePlant(id)

	if !isDeleted {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "No se pudo eliminar la planta"})
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Planta eliminada correctamente"})
}

func AddManyNutrient(context *gin.Context) {
	//Getting Data
	if func() bool {
		ids := context.PostForm("id")
		amount := context.PostForm("amount")
		return (ids == "[]" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids and amount fields"})
		return
	}

	fmt.Println(context.PostForm("ids"))
	plantsNutrientDTO := dto.PlantsSupplyDTO{}

	if err := context.ShouldBind(&plantsNutrientDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids<[]uint> and amount<uint> fields"})
		return
	}

	ids := plantsNutrientDTO.Ids
	if len(ids) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids and amount fields"})
		return
	}

	amountNutrient := plantsNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The amount field cannot be negative"})
		return
	}

	plants := func() (plants []Entities.Plant) {
		for _, plant := range plants {
			plantFinded, err := Services.GetPlant(plant.ID)
			if err != nil {
				continue
			}
			plants = append(plants, plantFinded)
		}
		return plants
	}()

	if len(plants) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Plants not found"})
		return
	}

	err := Services.AddManyNutrient(plants, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}

func AddNutrient(context *gin.Context) {
	//Getting Data
	if func() bool {
		id := context.PostForm("id")
		amount := context.PostForm("amount")
		return (id == "" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required id and amount fields"})
		return
	}

	plantNutrientDTO := dto.PlantSupplyDTO{}

	if err := context.ShouldBind(&plantNutrientDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required id<uint> and amount<uint> fields"})
		return
	}

	id := plantNutrientDTO.Id
	if id < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The id cannot be negative"})
		return
	}

	amountNutrient := plantNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The amount cannot be negative"})
		return
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Plant not found"})
		return
	}

	err = Services.AddNutrient(plant, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}

func AddManyWater(context *gin.Context) {
	//Getting Data
	if func() bool {
		ids := context.PostForm("id")
		amount := context.PostForm("amount")
		return (ids == "[]" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids and amount fields"})
		return
	}

	plantsWaterDTO := dto.PlantsSupplyDTO{}

	if err := context.ShouldBind(&plantsWaterDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids<[]uint> and amount<uint> fields"})
		return
	}

	ids := plantsWaterDTO.Ids
	if len(ids) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required ids and amount fields"})
		return
	}

	amountWater := plantsWaterDTO.Amount
	if amountWater < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The amount field cannot be negative"})
		return
	}

	plants := func() (plants []Entities.Plant) {
		for _, plant := range plants {
			plantFinded, err := Services.GetPlant(plant.ID)
			if err != nil {
				continue
			}
			plants = append(plants, plantFinded)
		}
		return plants
	}()

	if len(plants) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Plants not found"})
		return
	}

	err := Services.AddManyWater(plants, float64(amountWater))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}

func AddWater(context *gin.Context) {
	//Getting Data
	if func() bool {
		id := context.PostForm("id")
		amount := context.PostForm("amount")
		return (id == "" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required id and amount fields"})
		return
	}

	plantWatertDTO := dto.PlantSupplyDTO{}

	if err := context.ShouldBind(&plantWatertDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Required id<uint> and amount<uint> fields"})
		return
	}

	id := plantWatertDTO.Id
	if id < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The id cannot be negative"})
		return
	}

	amountWater := plantWatertDTO.Amount
	if amountWater < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "The amount cannot be negative"})
		return
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Plant not found"})
		return
	}

	err = Services.AddWater(plant, float64(amountWater))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}
