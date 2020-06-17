package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID                int     `json:"id"`
	Email             string  `json:"email"`
	Password          string  `json:"password,omitempty"`
	EncryptedPassword string  `json:"-"`
	Name              string  `json:"name"`
	Surname           string  `json:"surname"`
	Weight            float64 `json:"weight"`
	Height            float64 `json:"height"`
	Age               int     `json:"age"`
	WorkModel         float64 `json:"workmodel"`
	WorkFutureModel   float64 `json:"workfuturemodel"`
	Gender            bool    `json:"gender"`
}

//Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

//BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return nil
		}
		u.EncryptedPassword = enc
	}
	return nil
}

//Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

//ComparePassword ...
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
