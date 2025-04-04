package entity

import (
	"errors"
	"time"
)

type Expense struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	OriginalName string    `json:"original_name"`
	Description  string    `json:"description"`
	Bank         string    `json:"bank"`
	Card         string    `json:"card"`
	Timestamp    time.Time `json:"timestamp"`
	Value        float64   `json:"value"`
	CategoryID   *uint     `json:"category_id,omitempty"`
	Category     Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;" json:"category"`
	Tags         []Tag     `gorm:"many2many:expense_tags;" json:"tags"`
}

func (e *Expense) Validate() error {
	if e.Value == 0 {
		return errors.New("field 'value' is required")
	}
	if e.Name == "" {
		return errors.New("field 'name' is required")
	}
	return nil
}

type ExpenseByGroup struct {
	CategoryID    uint    `json:"category_id"`
	CategoryName  string  `json:"category_name"`
	TotalAmount   float64 `json:"total_amount"`
	CategoryColor string  `json:"category_color"`
}

type ExpenseByDate struct {
	Timestamp  string  `json:"timestamp"`
	TotalValue float64 `json:"total_value"`
}
