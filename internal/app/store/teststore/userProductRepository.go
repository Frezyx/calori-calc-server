package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//UserProductRepository ...
type UserProductRepository struct {
	store *Store
}

//Create ...
func (r *UserProductRepository) Create(uP *model.UserProduct) error {
	if err := uP.Validate(); err != nil {
		return err
	}
	return nil
}

// Get User Product by ID...
func (r *UserProductRepository) Get(int) ([]model.UserProduct, error) {
	return nil, nil
}
