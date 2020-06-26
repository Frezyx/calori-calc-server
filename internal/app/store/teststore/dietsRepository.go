package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//DietsRepository ...
type DietsRepository struct {
	store *Store
}

//Create ...
func (r *DietsRepository) Create(u *model.User, name string, isAutoCreated bool) error {
	return nil
}

//GetByID ...
func (r *DietsRepository) GetByID(id int) (*model.Diet, error) {
	return nil, nil
}

// GetAllByUserID ...
func (r *DietsRepository) GetAllByUserID(id int) ([]model.Diet, error) {
	return nil, nil
}

//Delete ...
func (r *DietsRepository) Delete(ID int) (bool, error) {
	return false, nil
}

//Edit ...
func (r *DietsRepository) Edit(d *model.Diet) error {
	return nil
}
