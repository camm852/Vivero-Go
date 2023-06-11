package Application

import (
	Entities2 "proyecto.com/Domain/Entities"
)

func GetPlants() []Entities2.Plant {
	var _plants Entities2.IPlantRepositorio
	var plants []Entities2.Plant
	plants = _plants.GetPlants()
	return plants
}

func GetPlant(id uint) Entities2.Plant {

}

func NewPlant(plant Entities2.PlantDTO) bool {

}

func UpdatePlant(plant Entities2.Plant) bool {

}

func DeletePlant(id uint) bool {

}

func AddWater(plant Entities2.Plant) Entities2.Plant {

}

func AddManyWater(plants []Entities2.Plant) []Entities2.Plant {

}

func AddNutrients(plant Entities2.Plant) Entities2.Plant {

}

func AddManyNutrients(plants []Entities2.Plant) []Entities2.Plant {

}

func DecreaseWater() {
	// Esta funcion recoge a todas las plantas y le agua baja segun lo explicado en el documento
}

func DecreaseNutrients() {
	// Esta funcion recoge a todas las plantas y le baja nutrientes segun lo explicado en el documento

}
