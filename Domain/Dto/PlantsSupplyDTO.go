package dto

type PlantsSupplyDTO struct {
	Ids    []uint `form:"ids"`
	Amount uint   `form:"amount"`
}
