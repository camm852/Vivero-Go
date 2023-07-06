package Entities

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Plants   []Plant
}
