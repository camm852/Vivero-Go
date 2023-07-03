package Entities

import "time"

type Plant struct {
	ID                      uint      `gorm:"primaryKey" json:"id"`
	Name                    string    `gorm:"column:name" json:"name"`
	DegreeSurvival          float64   `gorm:"column:degree_survival" json:"degreeSurvival"`
	AmountWaterRequired     float64   `gorm:"column:amount_water_required" json:"amountWaterRequired"`
	AmountWaterSystem       float64   `gorm:"column:amount_water_system" json:"amountWaterSystem"`
	DegreeHydration         uint      `gorm:"column:degree_hydration" json:"degreeHydration"`
	AmountNutrientsRequired uint      `gorm:"column:amount_nutrients_required" json:"amountNutrientsRequired"`
	AmountNutrientsSystem   float64   `gorm:"column:amount_nutrients_system" json:"amountNutrientsSystem"`
	IsDead                  bool      `gorm:"column:is_dead" json:"isDead"`
	DegreeNutrition         uint      `gorm:"column:degree_nutrition" json:"degreeNutrition"`
	Created                 time.Time `gorm:"column:created" json:"created"`
	LastUpdate              time.Time `gorm:"column:last_update" json:"lastUpdate"`
}
