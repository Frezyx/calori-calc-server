package sqlstore

import (
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
