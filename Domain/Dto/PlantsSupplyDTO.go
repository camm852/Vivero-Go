package dto

type PlantsNutrientSupplyDTO struct {
	Ids    []uint `form:"ids"`
	Amount uint   `form:"amount"`
}

type PlantsWaterSupplyDTO struct {
	Ids    []uint  `form:"ids"`
	Amount float64 `form:"amount"`
}
