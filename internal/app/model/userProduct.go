package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//UserProduct ...
type UserProduct struct {
	ID         int     `json:"id"`
	ProductID  int     `json:"productid"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Calory     float64 `json:"calory"`
	Squi       float64 `json:"squi"`
	Fat        float64 `json:"fat"`
	Carboh     float64 `json:"carboh"`
	Grams      float64 `json:"grams"`
	DateCreate int     `json:"date_created"`
	UserID     int     `json:"user_id"`
}

//Validate ...
func (uP *UserProduct) Validate() error {
	return validation.ValidateStruct(
		uP,
		validation.Field(&uP.Name, validation.Required),
		validation.Field(&uP.Category, validation.Required),
	)
}
