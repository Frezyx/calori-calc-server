package sqlstore

import (
	"errors"
	"log"

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

//Search ...
func (r *ProductRepository) Search(textRequest string) ([]model.Product, error) {
	if textRequest == "" {
		return nil, errors.New("empty request text")
	}

	products := []model.Product{}
	// textRequest
	rows, err := r.store.db.Query("SELECT * FROM products WHERE name ILIKE '%' || $1 || '%' LIMIT 10", textRequest)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := model.Product{}

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Category,
			&p.Calory,
			&p.Squi,
			&p.Fat,
			&p.Carboh,
		)

		if err != nil {
			continue
		}

		log.Println(p)

		products = append(products, p)
	}

	return products, nil
}
