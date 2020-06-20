package sqlstore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//DatesRepository ...
type DatesRepository struct {
	store *Store
}

//Create ...
func (r *DatesRepository) Create(d *model.Date) error {
	if err := d.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO dates (date_created, products_ids) VALUES ($1, $2) RETURNING id",
		d.Date,
		d.IDs,
	).Scan(&d.ID)
}
