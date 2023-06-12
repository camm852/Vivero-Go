package Entities

import "time"

type Plant struct {
	ID                      uint      `gorm:"primaryKey" json:"id"`
	Name                    string    `gorm:"column:name" json:"name"`
	DegreeSurvival          float64   `gorm:"column:degreeSurvival" json:"degreeSurvival"`
	AmountWaterRequired     float64   `gorm:"column:amountWaterRequired" json:"amountWaterRequired"`
	AmountWaterSystem       float64   `gorm:"column:amountWaterSystem" json:"amountWaterSystem"`
	DegreeHydration         int       `gorm:"column:degreeHydration" json:"degreeHydration"`
	AmountNutrientsRequired uint   `gorm:"column:amountNutrientsRequired" json:"amountNutrientsRequired"`
	AmountNutrientsSystem   float64   `gorm:"column:amountNutrientsSystem" json:"amountNutrientsSystem"`
	DegreeNutrition         int       `gorm:"column:degreeNutrition" json:"degreeNutrition"`
	Created                 time.Time `gorm:"column:created" json:"created"`
	LastUpdate              time.Time `gorm:"column:lastUpdate" json:"lastUpdate"`
}
