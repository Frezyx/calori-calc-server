package store

//Store ...
type Store interface {
	User() UserRepository
	Product() ProductRepository
}
