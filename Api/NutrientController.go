package Api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"proyecto.com/Application/Services"
	dto "proyecto.com/Domain/Dto"
	"proyecto.com/Domain/Entities"
)

func AddManyNutrient(context *gin.Context) {
	//Getting Data
	if func() bool {
		id := context.PostForm("id")
		amount := context.PostForm("amount")
		return (id == "" || amount == "")
	}() {
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	plantsNutrientDTO := dto.PlantsSupplyDTO{}

	if err := context.ShouldBind(&plantsNutrientDTO); err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ids := plantsNutrientDTO.Ids
	if len(ids) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	amountNutrient := plantsNutrientDTO.Amount
	if amountNutrient < 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	plants := func() (plants []Entities.Plant) {
		for _, plant := range plants {
			plantFinded, err := Services.GetPlant(plant.Id)
			if err != nil {
				continue
			}
			plants = append(plants, plantFinded)
		}
		return plants
	}()

	if len(plants) == 0 {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := Services.AddManyNutrient(plants, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
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

	err = Services.AddNutrient(plant, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}
