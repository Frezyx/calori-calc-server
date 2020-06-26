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
