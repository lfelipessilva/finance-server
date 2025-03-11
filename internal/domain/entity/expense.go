package entity

import (
	"errors"
	"time"
)

type Expense struct {
	ID        uint
	Name      string
	Category  string
	Timestamp time.Time
	Value     float64
}

func (e *Expense) Validate() error {
	if e.Value == 0 {
		return errors.New("field 'value' is required")
	}
	if e.Name == "" {
		return errors.New("field 'name' is required")
	}
	if e.Category == "" {
		return errors.New("field 'category' is required")
	}
	if e.Timestamp.IsZero() {
		return errors.New("field 'timestamp' must be a valid time")
	}
	return nil
}
