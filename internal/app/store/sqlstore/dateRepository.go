package sqlstore

import (
	"database/sql"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
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

//GetIfSet ...
func (r *DatesRepository) GetIfSet(date int) (interface{}, error) {
	type response struct {
		ID int `json:"id"`
	}

	resp := &response{}

	if err := r.store.db.QueryRow(
		"SELECT id FROM dates WHERE date_created = $1",
		date,
	).Scan(
		&resp.ID,
	); err != nil {
		if err == sql.ErrNoRows {
			return response{ID: -1}, store.ErrRecordNotFound
		}
		return response{ID: -1}, err
	}
	return resp, nil
}
