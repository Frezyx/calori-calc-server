package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//DatesRepository ...
type DatesRepository struct {
	store *Store
}

//Create ...
func (r *DatesRepository) Create(d *model.Date) error {
	return nil
}

//GetIfSet ...
func (r *DatesRepository) GetIfSet(d *model.Date) (interface{}, error) {
	return -1, nil
}

//GetIDsByDate ...
func (r *DatesRepository) GetIDsByDate(d *model.Date) (interface{}, error) {
	return nil, nil
}

//UpdateDate ...
func (r *DatesRepository) UpdateDate(d *model.Date) error {
	return nil
}

//DeleteAll ...
func (r *DatesRepository) DeleteAll() error {
	return nil
}
