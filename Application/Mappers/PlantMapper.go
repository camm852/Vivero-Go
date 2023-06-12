package Mappers

import (
	"math/rand"
	"proyecto.com/Domain/Entities"
	"time"
)

func MapPlantDtoToPlant(plantDTO Entities.PlantDTO) (*Entities.Plant, error) {
	// Generar un ID aleatorio
	id := uint(rand.Intn(1000))

	// Obtener el tiempo actual
	now := time.Now()

	// Crear una instancia de Plant y realizar el mapeo directo
	plant := &Entities.Plant{
		Id:                      id,
		Name:                    plantDTO.Name,
		DegreeSurvival:          plantDTO.DegreeSurvival,
		AmountWaterRequired:     plantDTO.AmountWaterRequired,
		AmountWaterSystem:       plantDTO.AmountWaterSystem,
		DegreeHydration:         plantDTO.DegreeHydration,
		AmountNutrientsRequired: plantDTO.AmountNutrientsRequired,
		AmountNutrientsSystem:   plantDTO.AmountNutrientsSystem,
		DegreeNutrition:         plantDTO.DegreeNutrition,
		LastUpdate:              now,
	}

	return plant, nil
}
