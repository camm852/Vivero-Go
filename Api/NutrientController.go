package Api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"proyecto.com/Application/Services"
	dto "proyecto.com/Domain/Dto"
)

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
