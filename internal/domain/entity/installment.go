package entity

import (
	"errors"
	"time"
)

type Installment struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Value          float64   `json:"value"`
	Quantity       int       `json:"quantity"`
	Expenses       []Expense `gorm:"foreignKey:InstallmentID;constraint:OnDelete:CASCADE;" json:"expenses"`
	TimestampStart time.Time `json:"timestamp_start"`
	TimestampEnd   time.Time `json:"timestamp_end"`
}

func (e *Installment) Validate() error {
	if e.Value == 0 {
		return errors.New("field 'value' is required")
	}
	if e.Quantity == 0 {
		return errors.New("field 'quantity' is required")
	}
	if e.TimestampStart.IsZero() {
		return errors.New("field 'timestamp_start' is required")
	}
	if e.TimestampEnd.IsZero() {
		return errors.New("field 'timestamp_end' is required")
	}
	return nil
}
