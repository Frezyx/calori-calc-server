package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
)

//UserProductRepository ...
type UserProductRepository struct {
	store *Store
}

//Create ...
func (r *UserProductRepository) Create(uP *model.UserProduct) (*model.UserProduct, error) {
	if err := uP.Validate(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow("INSERT INTO user_products (productid, name, category, calory, squi, fat, carboh, grams, date_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		uP.ProductID,
		uP.Name,
		uP.Category,
		uP.Calory,
		uP.Squi,
		uP.Fat,
		uP.Carboh,
		uP.Grams,
		uP.DateCreate,
	).Scan(&uP.ID)
	if err != nil {
		return nil, err
	}
	return uP, err
}

// Get User Product by ID...
func (r *UserProductRepository) Get(ID int) (*model.UserProduct, error) {
	if &ID == nil {
		return nil, errors.New("empty request id")
	}

	uP := &model.UserProduct{}

	if err := r.store.db.QueryRow(
		"SELECT * FROM user_products WHERE id = $1", ID).Scan(
		&uP.ID,
		&uP.ProductID,
		&uP.Name,
		&uP.Category,
		&uP.Calory,
		&uP.Squi,
		&uP.Fat,
		&uP.Carboh,
		&uP.Grams,
		&uP.DateCreate,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return uP, nil
}

//Edit ...
func (r *UserProductRepository) Edit(uP *model.UserProduct) error {

	return r.store.db.QueryRow("UPDATE user_products SET grams = $1, calory = $2, squi = $3, fat = $4, carboh = $5 WHERE id = $6 RETURNING id",
		uP.Grams,
		uP.Calory,
		uP.Squi,
		uP.Fat,
		uP.Carboh,
		uP.ID,
	).Scan(&uP.ID)
}

//Delete ...
func (r *UserProductRepository) Delete(ID int) (bool, error) {
	res, err := r.store.db.Exec("DELETE FROM user_products WHERE id=$1", ID)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil && count != 1 {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
	}

	return count == 1, nil
}

//DeleteAll ...
func (r *UserProductRepository) DeleteAll(UserID int) (bool, error) {
	rows, err := r.store.db.Query("SELECT product_id FROM user_products_join WHERE user_id = $1", UserID)
	for rows.Next() {

		p := model.Product{}

		err := rows.Scan(
			&p.ID,
		)

		if err != nil {
			continue
		}
		go r.store.UserProduct().DeleteInGorutine(p.ID, UserID)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

//DeleteInGorutine ...
func (r *UserProductRepository) DeleteInGorutine(ID int, UserID int) (bool, error) {
	res, err := r.store.db.Exec("DELETE FROM user_products WHERE id = $1", ID)

	count, err := res.RowsAffected()
	if err != nil && count != 1 {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
	}

	res1, err := r.store.db.Exec("DELETE FROM user_products_join WHERE user_id = $1", UserID)

	count1, err := res1.RowsAffected()
	if err != nil && count1 != 1 {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
	}

	return count1 == 1 && count == 0, nil
}

//JoinUser ...
func (r *UserProductRepository) JoinUser(uP *model.UserProduct) error {
	if err := uP.Validate(); err != nil {
		return err
	}
	return r.store.db.QueryRow("INSERT INTO user_products_join (product_id, user_id) VALUES ($1, $2) RETURNING id",
		uP.ID,
		uP.UserID,
	).Scan(&uP.ID)
}

//GetAllByUserID ...
func (r *UserProductRepository) GetAllByUserID(id int) ([]model.UserProduct, error) {
	uPList := []model.UserProduct{}
	rows, err := r.store.db.Query("select product_id from user_products_join WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productID int
		err := rows.Scan(
			&productID,
		)
		if err != nil {
			return nil, err
		}
		p, err := r.store.UserProduct().Get(productID)
		if err != nil {
			return nil, err
		}
		uPList = append(uPList, *p)
	}

	return uPList, nil
}
