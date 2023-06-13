package Entities

import (
	"proyecto.com/Infraestructure/Database"
)

type IPlantRepository interface {
	GetPlants() []Plant
	GetPlant(id uint) (Plant, error)
	NewPlant(plant *Plant) bool
	UpdatePlant(plant *Plant) bool
	DeletePlant(id uint) bool
}

// implementar interfaz
func (p Plant) GetPlants() []Plant {
	var plants []Plant
	_ = Database.DB.Find(&plants)

	return plants
}

func (p Plant) GetPlant(id uint) (Plant, error) {
	var plant Plant

	result := Database.DB.First(&plant, id)

	return plant, result.Error
}

func (p Plant) NewPlant(plant *Plant) bool {
	createdPlant := Database.DB.Create(plant)

	err := createdPlant.Error

	if err != nil {
		return false
	}
	return true
}

func (p Plant) UpdatePlant(plant *Plant) bool {
	plantFlag, err := p.GetPlant(plant.ID)

	if err != nil {
		return false
	}

	plantFlag.Name = plant.Name
	plantFlag.DegreeSurvival = plant.DegreeSurvival
	plantFlag.AmountWaterRequired = plant.AmountWaterRequired
	plantFlag.AmountWaterSystem = plant.AmountWaterSystem
	plantFlag.DegreeHydration = plant.DegreeHydration
	plantFlag.AmountNutrientsRequired = plant.AmountNutrientsRequired
	plantFlag.AmountNutrientsSystem = plant.AmountNutrientsSystem
	plantFlag.DegreeNutrition = plant.DegreeNutrition
	plantFlag.LastUpdate = plant.LastUpdate

	result := Database.DB.Save(&plantFlag)

	if result.Error != nil {
		return false
	}

	return true
}

func (p Plant) DeletePlant(id uint) bool {
	plant, err := p.GetPlant(id)
	if err != nil {
		return false
	}
	result := Database.DB.Delete(&plant)
	if result.Error != nil {
		return false
	}
	return true
}
