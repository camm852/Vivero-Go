package Entities

import "time"

type Plant struct {
	Id                      uint
	Name                    string
	DegreeSurvival          float64
	AmountWaterRequired     float64
	AmountWaterSystem       float64
	DegreeHydration         uint
	AmountNutrientsRequired float64
	AmountNutrientsSystem   float64
	DegreeNutrition         int
	LastUpdate              time.Time
}
