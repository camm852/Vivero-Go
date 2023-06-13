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

	return plantCreated
}

func UpdatePlant(plant *Entities.Plant) bool {
	var _plant Entities.IPlantRepository = Entities.Plant{}
	isUpdated := _plant.UpdatePlant(plant)

	return isUpdated
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

func DecreaseWater() {
	plants := GetPlants()

	for i := range plants {
		plant := &plants[i]

		if plant.AmountWaterSystem == 0 {
			//logica para que la planta muera
			continue
		} else if plant.AmountWaterSystem > (1 + plant.DegreeSurvival) {
			excessWater := plant.AmountWaterSystem - (plant.AmountWaterRequired + plant.DegreeSurvival)
			if excessWater > plant.DegreeSurvival {
				//logica para obtener el tiempo en el que la planta muere por exceso de agua
				plant.AmountWaterSystem -= 2 * (1 / float64(plant.DegreeHydration))
			} else {
				plant.AmountWaterSystem -= 2 * (1 / float64(plant.DegreeHydration))
			}
		} else if plant.AmountWaterSystem < plant.DegreeSurvival {
			plant.AmountWaterSystem -= 2 * (1 / float64(plant.DegreeHydration))
		} else {
			plant.AmountWaterSystem -= 1 / float64(plant.DegreeHydration)
		}

		if plant.AmountWaterSystem < 0 {
			plant.AmountWaterSystem = 0
		}

		UpdatePlant(plant)
	}
}

func DecreaseNutrients() {
	plants := GetPlants()

	for i := range plants {
		plant := &plants[i]

		if plant.AmountNutrientsSystem == 0 {
			//logica para que la planta muera
			continue
		} else if plant.AmountNutrientsSystem > (1 + plant.DegreeSurvival) {
			excessNutrients := plant.AmountNutrientsSystem - (plant.AmountWaterRequired + plant.DegreeSurvival)
			if excessNutrients > plant.DegreeSurvival {
				//logica para obtener el tiempo en el que la planta muere por exceso de agua
				plant.AmountNutrientsSystem -= 2 * (1 / float64(plant.DegreeHydration))
			} else {
				plant.AmountNutrientsSystem -= 2 * (1 / float64(plant.DegreeHydration))
			}
		} else if plant.AmountNutrientsSystem < plant.DegreeSurvival {
			plant.AmountNutrientsSystem -= 2 * (1 / float64(plant.DegreeHydration))
		} else {
			plant.AmountNutrientsSystem -= 1 / float64(plant.DegreeHydration)
		}

		if plant.AmountNutrientsSystem < 0 {
			plant.AmountNutrientsSystem = 0
		}

		UpdatePlant(plant)
	}
}
