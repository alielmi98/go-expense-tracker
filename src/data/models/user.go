package models

type User struct {
	BaseModel
	Username string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Expenses []Expense
}
