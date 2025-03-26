package entity

import (
	"time"
)

type Invoice struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Month     uint      `json:"month"`
	Year      uint      `json:"year"`
	Value     float64   `json:"value"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expenses  []Expense `json:"expenses"`
}
