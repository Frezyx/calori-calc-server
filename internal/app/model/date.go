package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//Date ...
type Date struct {
	ID   int    `json:"id"`
	Date int    `json:"date_created"`
	IDs  string `json:"products_ids"`
}

//Validate ...
func (d *Date) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.IDs, validation.Required),
		validation.Field(&d.Date, validation.Required),
	)
}
