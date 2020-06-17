package store

import "github.com/Frezyx/calory-calc-server/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	DeleteUser(int) (bool, error)
	Edit(*model.User) error
}

//ProductRepository ...
type ProductRepository interface {
	Create(*model.Product) error
}
