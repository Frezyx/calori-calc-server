package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//ProductRepository ...
type ProductRepository struct {
	store *Store
}

//Create ...
func (r *ProductRepository) Create(p *model.Product) error {
	return nil
}

//Search ...
func (r *ProductRepository) Search(textRequest string) ([]model.Product, error) {
	return nil, nil
}
