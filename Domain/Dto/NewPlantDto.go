package dto

type NewPlantDTO struct {
	Name                    string  `json:"name"`
	DegreeSurvival          float64 `json:"degreeSurvival"`
	AmountWaterRequired     float64 `json:"amountWaterRequired"`
	DegreeHydration         uint    `json:"degreeHydration"`
	AmountNutrientsRequired uint    `json:"amountNutrientsRequired"`
	DegreeNutrition         uint    `json:"degreeNutrition"`
}
