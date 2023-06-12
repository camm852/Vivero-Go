package Entities

type PlantDTO struct {
	Name                    string  `json:"name"`
	DegreeSurvival          float64 `json:"degreeSurvival"`
	AmountWaterRequired     float64 `json:"amountWaterRequired"`
	AmountWaterSystem       float64 `json:"amountWaterSystem"`
	DegreeHydration         uint    `json:"degreeHydration"`
	AmountNutrientsRequired uint    `json:"amountNutrientsRequired"`
	AmountNutrientsSystem   float64 `json:"amountNutrientsSystem"`
	DegreeNutrition         int     `json:"degreeNutrition"`
}
