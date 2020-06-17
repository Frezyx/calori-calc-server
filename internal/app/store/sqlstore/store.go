package sqlstore

import (
	"database/sql"

	"github.com/Frezyx/calory-calc-server/internal/app/store"

	// pq ...
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                *sql.DB
	userRepository    *UserRepository
	productRepository *ProductRepository
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
