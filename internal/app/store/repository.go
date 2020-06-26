package store

import "github.com/Frezyx/calory-calc-server/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	DeleteUser(int) (bool, error)
	Edit(*model.User) error
	GetByID(id int) (*model.User, error)
}

//ProductRepository ...
type ProductRepository interface {
	Create(*model.Product) error
	Search(textRequest string) ([]model.Product, error)
}

//UserProductRepository ...
type UserProductRepository interface {
	Create(*model.UserProduct) (*model.UserProduct, error)
	Get(int) (*model.UserProduct, error)
	Edit(u *model.UserProduct) error
	Delete(int) (bool, error)
	// Paste an user id
	DeleteAll(int) (bool, error)
	JoinUser(*model.UserProduct) error
	DeleteInGorutine(int, int) (bool, error)
}

//DatesRepository ...
type DatesRepository interface {
	Create(*model.Date) error
	GetIfSet(d *model.Date) (interface{}, error)
	GetIDsByDate(d *model.Date) (interface{}, error)
	UpdateDate(d *model.Date) error
	DeleteAll() error
}

//DietsRepository ...
type DietsRepository interface {
	Create(*model.User, string, bool) error
	GetByID(id int) (*model.Diet, error)
	GetAllByUserID(id int) ([]model.Diet, error)
}
