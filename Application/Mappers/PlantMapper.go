package Mappers

import (
	"proyecto.com/Domain/Entities"
	"time"
)

func MapPlantDtoToPlant(plantDTO Entities.PlantDTO) (*Entities.Plant, error) {

	// Obtener el tiempo actual
	now := time.Now()

	// Crear una instancia de Plant y realizar el mapeo directo
	plant := &Entities.Plant{
		Name:                    plantDTO.Name,
		DegreeSurvival:          plantDTO.DegreeSurvival,
		AmountWaterRequired:     plantDTO.AmountWaterRequired,
		AmountWaterSystem:       plantDTO.AmountWaterSystem,
		DegreeHydration:         plantDTO.DegreeHydration,
		AmountNutrientsRequired: plantDTO.AmountNutrientsRequired,
		AmountNutrientsSystem:   plantDTO.AmountNutrientsSystem,
		DegreeNutrition:         plantDTO.DegreeNutrition,
		Created:                 now,
		LastUpdate:              now,
	}

	return plant, nil
}
