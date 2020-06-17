package sqlstore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
)

//ProductRepository ...
type ProductRepository struct {
	store *Store
}

//Create ...
func (r *ProductRepository) Create(p *model.Product) error {
	if err := p.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO products (name, category, calory, squi, fat, carboh) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		p.Name,
		p.Category,
		p.Calory,
		p.Squi,
		p.Fat,
		p.Carboh,
	).Scan(&p.ID)
}
