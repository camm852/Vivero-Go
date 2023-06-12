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
	var plant Entities.Plant = Entities.Plant{}

	return plant, nil

}

func NewPlant(plant *Entities.Plant) bool {
	return true
}

func AddNutrient(plant Entities.Plant, amountNutrient uint) error {
	if plant.AmountNutrientsRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountNutrientsSystem += float64(amountNutrient) / float64(plant.AmountNutrientsRequired)
	return nil
}

func AddManyNutrient(plants []Entities.Plant, amountNutrient uint) error {
	var err error
	for _, plant := range plants {
		err = AddNutrient(plant, amountNutrient)
	}
	return err
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
