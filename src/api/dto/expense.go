package dto

import (
	"time"
)

type CreateExpenseRequest struct {
	Title    string    `json:"title" binding:"required"`
	Amount   float64   `json:"amount" binding:"required"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}

type UpdateExpenseRequest struct {
	Title    string    `json:"title"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}
type ExpenseResponse struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
	UserID   uint      `json:"userId"`
}
