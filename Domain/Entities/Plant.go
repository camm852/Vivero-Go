package Entities

import "time"

type Plant struct {
	ID                      uint      `gorm:"primaryKey" json:"id"`
	Name                    string    `gorm:"column:name" json:"name"`
	DegreeSurvival          float64   `gorm:"column:degreesurvival" json:"degreeSurvival"`
	AmountWaterRequired     float64   `gorm:"column:amountwaterrequired" json:"amountWaterRequired"`
	AmountWaterSystem       float64   `gorm:"column:amountwatersystem" json:"amountWaterSystem"`
	DegreeHydration         uint      `gorm:"column:degreehydration" json:"degreeHydration"`
	AmountNutrientsRequired uint      `gorm:"column:amountnutrientsrequired" json:"amountNutrientsRequired"`
	AmountNutrientsSystem   float64   `gorm:"column:amountnutrientssystem" json:"amountNutrientsSystem"`
	DegreeNutrition         uint      `gorm:"column:degreenutrition" json:"degreeNutrition"`
	Created                 time.Time `gorm:"column:created" json:"created"`
	LastUpdate              time.Time `gorm:"column:lastupdate" json:"lastUpdate"`
}
