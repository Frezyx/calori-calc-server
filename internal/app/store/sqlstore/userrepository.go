package sqlstore

import (
	"database/sql"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
)

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO users (email, name, surname, encrypted_password) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Email,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

//Edit ...
func (r *UserRepository) Edit(u *model.User) error {

	return r.store.db.QueryRow("UPDATE users SET email = $1, name = $2, surname = $3, weight = $4, height = $5, age = $6, workmodel = $7, workfuturemodel = $8, gender = $9 "+
		"WHERE id =$10 RETURNING id",
		u.Email,
		u.Name,
		u.Surname,
		u.Weight,
		u.Height,
		u.Age,
		u.WorkModel,
		u.WorkFutureModel,
		u.Gender,
		u.ID,
	).Scan(&u.ID)
}

//DeleteUser ...
func (r *UserRepository) DeleteUser(ID int) (bool, error) {
	res, err := r.store.db.Exec("DELETE FROM users WHERE id=$1", ID)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil && count != 1 {
		if err == sql.ErrNoRows {
			return false, store.ErrRecordNotFound
		}
	}

	return count == 1, nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// GetByID ...
func (r *UserRepository) GetByID(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, name, surname, weight, height, age, workmodel, workfuturemodel, gender FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Name,
		&u.Surname,
		&u.Weight,
		&u.Height,
		&u.Age,
		&u.WorkModel,
		&u.WorkFutureModel,
		&u.Gender,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
