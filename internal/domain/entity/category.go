package entity

import (
	"errors"
)

type Category struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Name  string
	Color string
	Icon  string
}

func (e *Category) Validate() error {
	if e.Color == "" {
		return errors.New("field 'color' is required")
	}
	if e.Name == "" {
		return errors.New("field 'name' is required")
	}
	return nil
}
