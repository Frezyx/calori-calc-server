package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//UserProductRepository ...
type UserProductRepository struct {
	store *Store
}

//Create ...
func (r *UserProductRepository) Create(uP *model.UserProduct) (*model.UserProduct, error) {
	return nil, nil
}

// Get User Product by ID...
func (r *UserProductRepository) Get(int) (*model.UserProduct, error) {
	return nil, nil
}

//GetAllByUserID ...
func (r *UserProductRepository) GetAllByUserID(id int) ([]model.UserProduct, error) {
	return nil, nil
}

// Edit User Product by ID...
func (r *UserProductRepository) Edit(u *model.UserProduct) error {
	return nil
}

//Delete ...
func (r *UserProductRepository) Delete(ID int) (bool, error) {
	return false, nil
}

//DeleteAll ...
func (r *UserProductRepository) DeleteAll(ID int) (bool, error) {
	return false, nil
}

//JoinUser ...
func (r *UserProductRepository) JoinUser(uP *model.UserProduct) error {
	return nil
}

//DeleteInGorutine ...
func (r *UserProductRepository) DeleteInGorutine(ID int, UserID int) (bool, error) {
	return false, nil
}
