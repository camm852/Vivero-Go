package Entities

import (
	"fmt"

	"proyecto.com/Infraestructure/Database"
)

type IPlantRepository interface {
	GetPlantsByUserId(userId uint) []Plant
	GetPlants() []Plant
	GetPlant(plantId uint) (Plant, error)
	NewPlant(plant *Plant) bool
	UpdatePlant(plant *Plant) bool
	DeletePlant(userId uint) bool
}

// implementar interfaz
func (p Plant) GetPlantsByUserId(userId uint) []Plant {
	var plants []Plant
	// Realizar consulta con GORM para obtener las plantas del usuario con el ID userId
	err := Database.DB.Where("user_id = ?", userId).Find(&plants).Error
	if err != nil {
		// Manejo del error
		// Por ejemplo, puedes devolver un slice vacío o registrar el error
		return []Plant{}
	}

	return plants
}

func (p Plant) GetPlants() []Plant {
	var plants []Plant
	// Realizar consulta con GORM para obtener las plantas del usuario con el ID userId
	err := Database.DB.Find(&plants).Error
	if err != nil {
		// Manejo del error
		// Por ejemplo, puedes devolver un slice vacío o registrar el error
		return []Plant{}
	}

	return plants
}

func (p Plant) GetPlant(plantId uint) (Plant, error) {
	var plant Plant

	result := Database.DB.Where("id = ?", plantId).First(&plant)
	fmt.Println(plantId, plant.ID)
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

	fmt.Println("update")

	plantFlag, err := p.GetPlant(plant.ID)

	if err != nil {
		fmt.Println("primer falso")
		return false
	}

	plantStored, errGet := p.GetPlant(plant.ID)

	if errGet != nil {
		fmt.Println("segundo falso")

		return false
	}

	if plantStored.IsDead {
		fmt.Println("muerta falso")
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
	plantFlag.IsDead = plant.IsDead

	result := Database.DB.Save(&plantFlag)

	if result.Error != nil {
		fmt.Println("tercer falso")

		return false
	}
	fmt.Println("actualizo")
	return true
}

func (p Plant) DeletePlant(userId uint) bool {
	plant, err := p.GetPlant(userId)
	if err != nil {
		return false
	}
	result := Database.DB.Where("user_id = ?", userId).Delete(&plant)
	if result.Error != nil {
		return false
	}
	return true
}
