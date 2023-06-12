package Services

import (
	Entities "proyecto.com/Domain/Entities"
)

func GetPlants() []Entities.Plant {
	var _plants Entities.IPlantRepository = Entities.Plant{}
	var plants []Entities.Plant
	plants = _plants.GetPlants()
	return plants
}

func GetPlant(id uint) Entities.Plant {
	//logica para encontrarlo
	var plant Entities.Plant = Entities.Plant{}

	return plant

}

func NewPlant(plant *Entities.Plant) bool {
	var _plants Entities.IPlantRepository = Entities.Plant{}

	var plantCreated = _plants.NewPlant(plant)

	if !plantCreated {
		return false
	}
	return true

}

/*func UpdatePlant(plant Entities.Plant) bool {

}

func DeletePlant(id uint) bool {

}

func AddWater(plant Entities.Plant) Entities.Plant {

}

func AddManyWater(plants []Entities.Plant) []Entities.Plant {

}

func AddNutrients(plant Entities.Plant) Entities.Plant {

}

func AddManyNutrients(plants []Entities.Plant) []Entities.Plant {

}

func DecreaseWater() {
	// Esta funcion recoge a todas las plantas y le agua baja segun lo explicado en el documento
}

func DecreaseNutrients() {
	// Esta funcion recoge a todas las plantas y le baja nutrientes segun lo explicado en el documento

}*/
