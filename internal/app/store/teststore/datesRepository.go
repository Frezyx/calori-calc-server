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
