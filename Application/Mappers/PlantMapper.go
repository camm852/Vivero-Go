package Mappers

import (
	"proyecto.com/Domain/Dto"
	"proyecto.com/Domain/Entities"
	"time"
)

func MapNewPlantDtoToPlant(newPlantDto dto.NewPlantDTO) (*Entities.Plant, error) {

	// Obtener el tiempo actual
	now := time.Now()

	// Crear una instancia de Plant y realizar el mapeo directo
	plant := &Entities.Plant{
		Name:                    newPlantDto.Name,
		DegreeSurvival:          newPlantDto.DegreeSurvival,
		AmountWaterRequired:     newPlantDto.AmountWaterRequired,
		AmountWaterSystem:       1.0 + newPlantDto.DegreeSurvival,
		DegreeHydration:         newPlantDto.DegreeHydration,
		AmountNutrientsRequired: newPlantDto.AmountNutrientsRequired,
		AmountNutrientsSystem:   1.0 + newPlantDto.DegreeSurvival,
		IsDead:                  false,
		DegreeNutrition:         newPlantDto.DegreeNutrition,
		Created:                 now,
		LastUpdate:              now,
	}

	return plant, nil
}

func MapUpdatePlantDtoToPlant(updatePlantDto dto.UpdatePlantDTO) (*Entities.Plant, error) {

	// Obtener el tiempo actual
	now := time.Now()

	// Crear una instancia de Plant y realizar el mapeo directo
	plant := &Entities.Plant{
		ID:                      updatePlantDto.ID,
		Name:                    updatePlantDto.Name,
		DegreeSurvival:          updatePlantDto.DegreeSurvival,
		AmountWaterRequired:     updatePlantDto.AmountWaterRequired,
		AmountWaterSystem:       updatePlantDto.AmountWaterSystem,
		DegreeHydration:         updatePlantDto.DegreeHydration,
		AmountNutrientsRequired: updatePlantDto.AmountNutrientsRequired,
		IsDead:                  updatePlantDto.IsDead,
		AmountNutrientsSystem:   updatePlantDto.AmountNutrientsSystem,
		DegreeNutrition:         updatePlantDto.DegreeNutrition,
		LastUpdate:              now,
	}

	return plant, nil
}
