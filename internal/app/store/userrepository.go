package store

import "github.com/Frezyx/calory-calc-server/model"

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow("INSERT INTO users (email, encripted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncriptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encripted_password FROM users WHERE email = &1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncriptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
