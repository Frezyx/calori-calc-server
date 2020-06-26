package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//Diet ...
type Diet struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	Name          string  `json:"name"`
	Calory        float64 `json:"calory"`
	Squi          float64 `json:"squi"`
	Fat           float64 `json:"fat"`
	Carboh        float64 `json:"carboh"`
	DateCreate    int     `json:"date_created"`
	IsAutoCreated bool    `json:"is_auto_created"`
}

//Validate ...
func (d *Diet) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.Name, validation.Required),
	)
}
