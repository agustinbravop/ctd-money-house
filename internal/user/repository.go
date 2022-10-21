package user

import (
	"database/sql"

	"github.com/agustinbravop/ctd-money-house/internal/domain"
)

type Repository interface {
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	// Create(p domain.User) error
	// Update(id int, p domain.User) error
	// Delete(id int) error
	// Exists(id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(id int) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Dni,
		&user.Email,
		&user.Telephone,
		&user.Password,
		&user.Cvu,
		&user.Email,
	)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) GetAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := domain.User{}
		_ = rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastName,
			&user.Dni,
			&user.Email,
			&user.Telephone,
			&user.Password,
			&user.Cvu,
			&user.Email)
		users = append(users, user)
	}

	return users, nil
}
