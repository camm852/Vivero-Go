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

	userId := context.MustGet("userId").(float64)
	fmt.Println(userId)
	var plants []Entities.Plant = Services.GetPlantsByUserId(Utils.ParseFloat64ToUint(userId))

	context.IndentedJSON(http.StatusOK, plants)
}

func GetPlant(context *gin.Context) {
	var id uint = Utils.ParseUint(context.Param("id"))

	plant, error := Services.GetPlant(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Cannot find plant"})
		return
	}
	context.IndentedJSON(http.StatusOK, plant)
}

func NewPlant(context *gin.Context) {
	var newPlantDto dto.NewPlantDTO //Token de usuario en el dto
	if err := context.ShouldBindJSON(&newPlantDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	plant, err := Mappers.MapNewPlantDtoToPlant(newPlantDto)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Cannot map data"})
		return
	}
	// obtener el token y pasarlo comoo parametro
	userId := context.MustGet("userId").(float64)
	plant.UserID = Utils.ParseFloat64ToUint(userId)
	var isCreated = Services.NewPlant(plant)

	if !isCreated {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Cannot create plant"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "Plant created succesfully"})

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
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Cannot map data"})
		return
	}

	var isUpdated = Services.UpdatePlant(plant)

	if !isUpdated {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Cannot update plant"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "Plant updated succesfully"})
}

func DeletePlant(context *gin.Context) {
	var id uint = Utils.ParseUint(context.Param("id"))

	_, error := Services.GetPlant(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Cannot find plant"})
		return
	}

	isDeleted := Services.DeletePlant(id)

	if !isDeleted {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Cannot delete plant"})
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Plant deleted succesfully"})
}

func AddNutrient(context *gin.Context) {
	//Getting Data
	var plantNutrientDTO dto.PlantNutrientSupplyDTO
	if func() bool {
		if err := context.ShouldBindJSON(&plantNutrientDTO); err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return true
		}
		return false
	}() {
		fmt.Println("Required id and amount fields")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required id and amount fields"})
		return
	}
	/*if err := context.ShouldBind(&plantNutrientDTO); err != nil {
		// Bad Request
		fmt.Println("Required id<uint> and amount<uint> fields")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required id<uint> and amount<uint> fields"})
		return
	}*/

	id := plantNutrientDTO.Id
	if id < 0 {
		// Bad Request

		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The id cannot be negative"})
		return
	}

	amountNutrient := plantNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The amount cannot be negative"})
		return
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Plant not found"})
		return
	}

	err = Services.AddNutrient(plant, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Added successfully"})
}

func AddManyNutrient(context *gin.Context) {
	//Getting Data
	var plantsNutrientDTO dto.PlantsNutrientSupplyDTO

	if err := context.ShouldBindJSON(&plantsNutrientDTO); err != nil {
		// Bad Request
		fmt.Println(err.Error())
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println(plantsNutrientDTO)
	ids := plantsNutrientDTO.Ids
	if len(ids) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required ids and amount fields"})
		return
	}

	amountNutrient := plantsNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The amount field cannot be negative"})
		return
	}

	plants := func() (plants []Entities.Plant) {
		for _, id := range plantsNutrientDTO.Ids {
			plantStored, err := Services.GetPlant(uint(id))
			if err != nil {
				continue
			}
			plants = append(plants, plantStored)
		}
		return plants
	}()
	fmt.Println(plants)
	if len(plants) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Plants not found"})
		return
	}

	err := Services.AddManyNutrient(plants, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Added successfully"})
}

func AddWater(context *gin.Context) {
	//Getting Data
	var plantWatertDTO dto.PlantWaterSupplyDTO
	if func() bool {
		if err := context.ShouldBindJSON(&plantWatertDTO); err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return true
		}
		return false
	}() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required id and amount fields"})
		return
	}

	/*if err := context.ShouldBind(&plantWatertDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required id<uint> and amount<uint> fields"})
		return
	}*/

	id := plantWatertDTO.Id
	if id < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The id cannot be negative"})
		return
	}

	amountWater := plantWatertDTO.Amount
	if amountWater < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The amount cannot be negative"})
		return
	}
	if amountWater > 1 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The amount cannot be greater than one"})
		return
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Plant not found"})
		return
	}

	err = Services.AddWater(plant, float64(amountWater))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Added successfully"})
}

func AddManyWater(context *gin.Context) {
	//Getting Data
	plantsWaterDTO := dto.PlantsWaterSupplyDTO{}

	if err := context.ShouldBindJSON(&plantsWaterDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required ids<[]uint> and amount<uint> fields"})
		return
	}

	ids := plantsWaterDTO.Ids
	if len(ids) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Required ids and amount fields"})
		return
	}

	amountWater := plantsWaterDTO.Amount
	if amountWater < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "The amount field cannot be negative"})
		return
	}

	plants := func() (plants []Entities.Plant) {
		var plantsStored []Entities.Plant
		for _, id := range plantsWaterDTO.Ids {
			plantStored, err := Services.GetPlant(uint(id))
			if err != nil {
				continue
			}
			plantsStored = append(plantsStored, plantStored)
		}
		return plantsStored
	}()

	if len(plants) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Plants not found"})
		return
	}

	err := Services.AddManyWater(plants, float64(amountWater))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": "Added successfully"})
}
