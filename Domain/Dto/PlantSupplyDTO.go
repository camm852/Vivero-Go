package dto

type PlantNutrientSupplyDTO struct {
	Id     uint `form:"id"`
	Amount uint `form:"amount"`
}
type PlantWaterSupplyDTO struct {
	Id     uint    `form:"id"`
	Amount float64 `form:"amount"`
}
