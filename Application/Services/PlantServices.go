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

func UpdatePlant(plant *Entities.Plant) bool {
	var _plant Entities.IPlantRepository = Entities.Plant{}
	isUpdated := _plant.UpdatePlant(plant)
	if !isUpdated {
		return false
	}
	return true
}

func AddNutrient(plant Entities.Plant, amountNutrient uint) error {
	if plant.AmountNutrientsRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountNutrientsSystem += float64(amountNutrient) / float64(plant.AmountNutrientsRequired)
	plant.UpdatePlant(&plant)
	return nil
}

func AddWater(plant Entities.Plant, amountWater float64) error {
	if plant.AmountWaterRequired == 0 {
		return errors.New("Division by zero")
	}
	plant.AmountWaterSystem += float64(amountWater) / float64(plant.AmountWaterRequired)
	return nil
}

func DeletePlant(id uint) bool {
	var _plant Entities.IPlantRepository = Entities.Plant{}
	Plant2Delete, err := _plant.GetPlant(id)

	if err != nil {
		return false
	} else {
		_plant.DeletePlant(Plant2Delete.ID)
		return true
	}
}

func AddManyNutrient(plants []Entities.Plant, amountNutrient uint) error {
	var err error
	for _, plant := range plants {
		err = AddNutrient(plant, amountNutrient)
	}
	return err
}

func AddManyWater(plants []Entities.Plant, amountWater float64) error {

	var err error
	for _, plant := range plants {
		err = AddWater(plant, amountWater)
	}
	return err

}

/*
func DecreaseWater() {
	// Esta funcion recoge a todas las plantas y le agua baja segun lo explicado en el documento
}

func DecreaseNutrients() {
	// Esta funcion recoge a todas las plantas y le baja nutrientes segun lo explicado en el documento

}*/
