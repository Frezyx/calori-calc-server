package teststore

import (
	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
)

// Store ...
type Store struct {
	userRepository        *UserRepository
	productRepository     *ProductRepository
	userProductRepository *UserProductRepository
	datesRepository       *DatesRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
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
