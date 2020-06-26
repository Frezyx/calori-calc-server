package sqlstore

import (
	"database/sql"
	"math"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
)

//DietsRepository ...
type DietsRepository struct {
	store *Store
}

//Create ...
func (r *DietsRepository) Create(u *model.User, name string, isAutoCreated bool) error {

	u, err := r.store.User().GetByID(u.ID)
	if err != nil || u == nil {
		return err
	}

	var squiPercent, fatPercent, carbohPercent float64
	var genderDelta float64

	if u.Gender {
		genderDelta = 5
	} else {
		genderDelta = -161
	}

	caloryLimit := (10*u.Weight + 6.25*u.Height - (4.92 * float64(u.Age)) + genderDelta) * u.WorkModel

	if u.WorkFutureModel == 1 {
		fatPercent = 0.30
		squiPercent = 0.30
		carbohPercent = 0.40
	} else if u.WorkFutureModel == 2 {
		fatPercent = 0.35
		squiPercent = 0.35
		carbohPercent = 0.45
	} else if u.WorkFutureModel == 3 {
		fatPercent = 0.275
		squiPercent = 0.325
		carbohPercent = 0.50
	}

	d := &model.Diet{
		UserID:        u.ID,
		Name:          name,
		Calory:        caloryLimit,
		Squi:          makeDietPart(caloryLimit*squiPercent, 4),
		Fat:           makeDietPart(caloryLimit*fatPercent, 9),
		Carboh:        makeDietPart(caloryLimit*carbohPercent, 4),
		IsAutoCreated: isAutoCreated,
	}

	return r.store.db.QueryRow("INSERT INTO diets (user_id, name, calory, squi, fat, carboh, is_auto_created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		d.UserID,
		d.Name,
		d.Calory,
		d.Squi,
		d.Fat,
		d.Carboh,
		d.IsAutoCreated,
	).Scan(&d.ID)
}

func makeDietPart(num float64, part float64) float64 {
	return roundMidle(num / part)
}

func roundMidle(num float64) float64 {
	return math.Round(num*100) / 100
}

//GetByID ...
func (r *DietsRepository) GetByID(id int) (*model.Diet, error) {
	d := &model.Diet{}
	if err := r.store.db.QueryRow(
		"SELECT id, user_id, name, calory, squi, fat, carboh, is_auto_created FROM diets WHERE id = $1",
		id,
	).Scan(
		&d.ID,
		&d.UserID,
		&d.Name,
		&d.Calory,
		&d.Squi,
		&d.Fat,
		&d.Carboh,
		&d.IsAutoCreated,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return d, nil
}

//GetAllByUserID ...
func (r *DietsRepository) GetAllByUserID(id int) ([]model.Diet, error) {
	dList := []model.Diet{}
	rows, err := r.store.db.Query("select * from diets diets WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		d := model.Diet{}
		err := rows.Scan(
			&d.ID,
			&d.UserID,
			&d.Name,
			&d.Calory,
			&d.Squi,
			&d.Fat,
			&d.Carboh,
			&d.IsAutoCreated,
		)
		if err != nil {
			return nil, err
		}
		dList = append(dList, d)
	}

	return dList, nil
}
