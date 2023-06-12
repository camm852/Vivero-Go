package Entities

import (
	"fmt"
	"proyecto.com/Infraestructure/Database"
)

type IPlantRepository interface {
	GetPlants() []Plant
	GetPlant(id uint) (Plant, error)
	NewPlant(plant *Plant) bool
	UpdatePlant(plant Plant) bool
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

	fmt.Println(err)

	if err != nil {
		return false
	}
	return true
}

func (p Plant) UpdatePlant(plant Plant) bool {
	//var  updated = codigo para actualizar la planta
	//return updated
	return false
}

func (p Plant) DeletePlant(id uint) bool {
	//var  updated = codigo para eliminar de la base de datos
	//return updated
	return false
}
