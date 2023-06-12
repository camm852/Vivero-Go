package Api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"proyecto.com/Application/Services"
)

func AddNutrient(context *gin.Context) {
	//Getting Data
	id, err := strconv.Atoi(context.Param("id"))
	if (err != nil) && (id < 0) {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
	}

	amountNutrient, err := strconv.Atoi(context.Param("amountNutrient"))
	if (err != nil) && (amountNutrient < 0) {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
	}

	plant, err := Services.GetPlant(uint(id))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
	}

	err = Services.AddNutrients(plant, uint(amountNutrient))
	if err != nil {
		// Bad Request
		context.IndentedJSON(http.StatusBadRequest, nil)
	}
	context.IndentedJSON(http.StatusOK, nil)

}
