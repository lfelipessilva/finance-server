package entity

import (
	"time"
)

type Installment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Quantity  uint      `json:"quantity"`
	Value     uint      `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Installment) Validate() error {
	return nil
}
