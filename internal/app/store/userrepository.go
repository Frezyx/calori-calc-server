package store

import "github.com/Frezyx/calory-calc-server/model"

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
