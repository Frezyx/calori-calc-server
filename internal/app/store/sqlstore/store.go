package sqlstore

import (
	"database/sql"

	"github.com/Frezyx/calory-calc-server/internal/app/store"

	// pq ...
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                    *sql.DB
	userRepository        *UserRepository
	productRepository     *ProductRepository
	userProductRepository *UserProductRepository
	datesRepository       *DatesRepository
	dietsRepository       *DietsRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Product ...
func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}

	return s.productRepository
}

// UserProduct ...
func (s *Store) UserProduct() store.UserProductRepository {
	if s.userProductRepository != nil {
		return s.userProductRepository
	}

	s.userProductRepository = &UserProductRepository{
		store: s,
	}

	return s.userProductRepository
}

// Dates ...
func (s *Store) Dates() store.DatesRepository {
	if s.datesRepository != nil {
		return s.datesRepository
	}

	s.datesRepository = &DatesRepository{
		store: s,
	}

	return s.datesRepository
}

// Diets ...
func (s *Store) Diets() store.DietsRepository {
	if s.dietsRepository != nil {
		return s.dietsRepository
	}

	s.dietsRepository = &DietsRepository{
		store: s,
	}

	return s.dietsRepository
}
