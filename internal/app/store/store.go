package store

//Store ...
type Store interface {
	User() UserRepository
	Product() ProductRepository
	UserProduct() UserProductRepository
	Dates() DatesRepository
	Diets() DietsRepository
}
