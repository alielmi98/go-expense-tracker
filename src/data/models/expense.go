package models

import "time"

type Expense struct {
	BaseModel
	Title    string    `gorm:"not null"`
	Amount   float64   `gorm:"not null"`
	Category string    `gorm:"not null"`
	Date     time.Time `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	User     User      `gorm:"foreignKey:UserID"`
}
