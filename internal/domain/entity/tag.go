package entity

import (
	"errors"
)

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (e *Tag) Validate() error {
	if e.Name == "" {
		return errors.New("field 'name' is required")
	}
	if e.Color == "" {
		return errors.New("field 'color' is required")
	}
	return nil
}
