package dto

type UpdatePlantDTO struct {
	ID                      uint    `json:"id"`
	Name                    string  `json:"name"`
	DegreeSurvival          float64 `json:"degreeSurvival"`
	AmountWaterRequired     float64 `json:"amountWaterRequired"`
	AmountWaterSystem       float64 `json:"amountWaterSystem"`
	DegreeHydration         uint    `json:"degreeHydration"`
	IsDead                  bool    `json:"isDead"`
	AmountNutrientsRequired uint    `json:"amountNutrientsRequired"`
	AmountNutrientsSystem   float64 `json:"amountNutrientsSystem"`
	DegreeNutrition         uint    `json:"degreeNutrition"`
}
