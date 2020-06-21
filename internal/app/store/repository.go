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
	Search(textRequest string) ([]model.Product, error)
}

//UserProductRepository ...
type UserProductRepository interface {
	Create(*model.UserProduct) error
	Get(int) (*model.UserProduct, error)
	Edit(u *model.UserProduct) error
	Delete(int) (bool, error)
}

//DatesRepository ...
type DatesRepository interface {
	Create(*model.Date) error
	GetIfSet(int) (interface{}, error)
	GetIDsByDate(int) (interface{}, error)
}
