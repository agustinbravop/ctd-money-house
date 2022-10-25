package user

import (
	"ctd-money-house/internal/domain"
	"database/sql"
)

var (
	queryCreate        = "INSERT INTO users (name, last_name, dni, email, telephone, cvu, alias) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	queryDeleteById    = "DELETE FROM sellers WHERE id=?"
	queryValidateCvu   = "SELECT id FROM users WHERE cvu=?"
	queryValidateAlias = "SELECT id FROM users WHERE alias=?"
)

type Repository interface {
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	Create(user domain.User) (int, error)
	Delete(id int) error
	ValidateCvuOrAlias(fieldMap map[string]interface{}) bool
	// Update(id int, p domain.User) error
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
			&user.Cvu,
			&user.Email)
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Create(u domain.User) (int, error) {
	stmt, err := r.db.Prepare(queryCreate)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(u.Name, u.LastName, u.Dni, u.Email, u.Telephone, u.Cvu, u.Alias)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(queryDeleteById)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) ValidateCvuOrAlias(fieldMap map[string]interface{}) bool {
	var row *sql.Row
	if fieldMap["cvu"] != nil {
		row = r.db.QueryRow(queryValidateCvu, fieldMap["cvu"])
	}
	if fieldMap["alias"] != nil {
		row = r.db.QueryRow(queryValidateAlias, fieldMap["alias"])
	}
	var id int
	err := row.Scan(&id)
	return err == nil
}
