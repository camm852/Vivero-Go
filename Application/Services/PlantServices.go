package Services

import (
	"errors"

	Entities "proyecto.com/Domain/Entities"
)

func GetPlants() []Entities.Plant {
	var _plants Entities.IPlantRepository = Entities.Plant{}
	var plants []Entities.Plant
	plants = _plants.GetPlants()
	return plants
}

func GetPlant(id uint) (Entities.Plant, error) {
	//logica para encontrarlo
	var _plant Entities.IPlantRepository = Entities.Plant{}
	plant, err := _plant.GetPlant(id)
	return plant, err
}

func NewPlant(plant *Entities.Plant) bool {
	var _plant Entities.IPlantRepository = Entities.Plant{}

	var plantCreated = _plant.NewPlant(plant)

	if !plantCreated {
		return false
	}
	return true

}

func UpdatePlant(plant Entities.Plant) bool {
	plantFlag, err := GetPlant(plant.ID)

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
	plantFlag.LastUpdate = plant.Created

	// todo lo que esta aqui deberia estar en IPlantRepository en la funcion update

	return true
}

func AddNutrients(plant Entities.Plant, amountNutrients uint) error {
	if plant.AmountNutrientsRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountNutrientsSystem += float64(amountNutrients) / float64(plant.AmountNutrientsRequired)
	return nil
}

func AddWater(plant Entities.Plant, amountWater float64) error {
	if plant.AmountWaterRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountWaterSystem += float64(amountWater) / float64(plant.AmountWaterRequired)
	return nil
}

/*

func DeletePlant(id uint) bool {

}

func AddManyWater(plants []Entities.Plant) []Entities.Plant {

}

func AddManyNutrients(plants []Entities.Plant) []Entities.Plant {

}

func DecreaseWater() {
	// Esta funcion recoge a todas las plantas y le agua baja segun lo explicado en el documento
}

func DecreaseNutrients() {
	// Esta funcion recoge a todas las plantas y le baja nutrientes segun lo explicado en el documento

}*/
