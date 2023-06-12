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
	plant, err = _plant.GetPlant(id)

	return plant, err

}

func NewPlant(plant *Entities.Plant) bool {
	var _plants Entities.IPlantRepository = Entities.Plant{}

	var plantCreated = _plants.NewPlant(plant)

	if !plantCreated {
		return false
	}
	return true

}

func AddNutrients(plant Entities.Plant, amountNutrients uint) error {
	if plant.AmountNutrientsRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountNutrientsSystem += float64(amountNutrients) / float64(plant.AmountNutrientsRequired)
	return nil
}

/*func UpdatePlant(plant Entities.Plant) bool {

}

func DeletePlant(id uint) bool {

}

func AddWater(plant Entities.Plant) Entities.Plant {

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
