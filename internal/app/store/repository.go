package store

import "github.com/Frezyx/calory-calc-server/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
