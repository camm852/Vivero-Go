package Entities

type IPlantRepository interface {
	GetPlants() []Plant
	GetPlant(id uint) Plant
	NewPlant(plant Plant) bool
	UpdatePlant(plant Plant) bool
	DeletePlant(id uint) bool
}

// implementar interfaz
func (s Plant) GetPlants() []Plant {
	plants := make([]Plant, 0)
	//plants = codigo para traer las plantas de la bd
	return plants
}

func (s Plant) GetPlant(id uint) Plant {
	var plant Plant
	//plants = codigo para traer la planta de la bd
	return plant
}

func (s Plant) NewPlant(plant Plant) bool {
	//var created = codigo para guardar la planta
	//return creada
	return false
}

func (s Plant) UpdatePlant(plant Plant) bool {
	//var  updated = codigo para actualizar la planta
	//return updated
	return false
}

func (s Plant) DeletePlant(id uint) bool {
	//var  updated = codigo para eliminar de la base de datos
	//return updated
	return false
}
