package sqlstore

import (
	"errors"
	"log"

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

	return r.store.db.QueryRow("INSERT INTO user_products (productid, name, category, calory, squi, fat, carboh, grams, date_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
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
}

//TODO: Change to one object return

// Get User Product by ID...
func (r *UserProductRepository) Get(ID int) ([]model.UserProduct, error) {
	if &ID == nil {
		return nil, errors.New("empty request id")
	}

	products := []model.UserProduct{}
	// textRequest
	rows, err := r.store.db.Query("SELECT * FROM user_products WHERE id = $1", ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := model.UserProduct{}

		err := rows.Scan(
			&p.ID,
			&p.ProductID,
			&p.Name,
			&p.Category,
			&p.Calory,
			&p.Squi,
			&p.Fat,
			&p.Carboh,
			&p.Grams,
			&p.DateCreate,
		)

		if err != nil {
			continue
		}

		log.Println(p)

		products = append(products, p)
	}

	return products, nil
}
