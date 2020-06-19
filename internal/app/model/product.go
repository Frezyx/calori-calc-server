package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//Product ...
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Calory   float64 `json:"calory"`
	Squi     float64 `json:"squi"`
	Fat      float64 `json:"fat"`
	Carboh   float64 `json:"carboh"`
}

//Validate ...
func (p *Product) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Category, validation.Required),
	)
}
