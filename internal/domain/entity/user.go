package entity

import (
	"errors"
)

type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Provider       string `json:"provider"`
	ProviderUserID string `json:"prover_user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

func (e *User) Validate() error {
	if e.Name == "" {
		return errors.New("field 'name' is required")
	}
	if e.Email == "" {
		return errors.New("field 'email' is required")
	}
	return nil
}
