package sqlstore

import (
	"database/sql"
	"strconv"
	"strings"

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

//GetIDsByDate ...
func (r *DatesRepository) GetIDsByDate(date int) (interface{}, error) {

	type response struct {
		ID         int   `json:"id"`
		Date       int   `json:"date_created"`
		ProductIDs []int `json:"product_ids"`
	}

	stringIDs := ""
	resp := &response{}

	if err := r.store.db.QueryRow(
		"SELECT * FROM dates WHERE date_created = $1",
		date,
	).Scan(
		&resp.ID,
		&stringIDs,
		&resp.Date,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	resp.ProductIDs = getSliceFromString(stringIDs)

	return resp, nil
}

func getSliceFromString(s string) []int {
	sliceIDs := []int{}

	stringSliceIDs := strings.Split(s, ",")
	for _, stringID := range stringSliceIDs {
		id, err := strconv.Atoi(stringID)
		if err != nil {
			sliceIDs = append(sliceIDs, -1)
		} else {
			sliceIDs = append(sliceIDs, id)
		}
	}

	return sliceIDs
}
