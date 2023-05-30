package user

import (
	"database/sql"
	"users-api/internal/domain"
)

var (
	queryUpdate        = "UPDATE users SET first_name=?, last_name=?, dni=?, email=?, telephone=? WHERE id=?"
	queryCreate        = "INSERT INTO users (id, first_name, last_name, dni, email, telephone) VALUES (?, ?, ?, ?, ?, ?)"
	queryDeleteById    = "DELETE FROM users WHERE id=?"
	queryExistsByEmail = "SELECT id FROM users WHERE email=?"
	queryGetAll        = "SELECT id, first_name, last_name, dni, email, telephone FROM users"
	queryGetByID       = "SELECT id, first_name, last_name, dni, email, telephone FROM users WHERE id = ?"
	queryGetByEmail    = "SELECT id, first_name, last_name, dni, email, telephone FROM users WHERE email = ?"
)

type Repository interface {
	GetByID(id string) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Update(p domain.User) error
	Create(user domain.User) (string, error)
	Delete(id string) error
	ExistsByEmail(email string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(id string) (domain.User, error) {
	row := r.db.QueryRow(queryGetByID, id)
	user, err := scanUserFromRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ErrNotFound
		} else {
			return domain.User{}, ErrInternal
		}
	}
	return user, nil
}

func (r *repository) GetByEmail(email string) (domain.User, error) {
	row := r.db.QueryRow(queryGetByEmail, email)
	user, err := scanUserFromRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ErrNotFound
		} else {
			return domain.User{}, ErrInternal
		}
	}
	return user, nil
}

func (r *repository) GetAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, _ := scanUserFromRow(rows)
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Create(u domain.User) (string, error) {
	stmt, err := r.db.Prepare(queryCreate)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(u.ID, u.Name, u.LastName, u.Dni, u.Email, u.Telephone)
	if err != nil {
		return "", err
	}

	return u.ID, nil
}

func (r *repository) Delete(id string) error {
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

// ExistsByEmail retorna True cuando ya existe un domain.User con ese email en la base de datos.
func (r *repository) ExistsByEmail(email string) bool {
	row := r.db.QueryRow(queryExistsByEmail, email)
	return isIdPresent(row)
}

func (r *repository) Update(u domain.User) error {
	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Name, u.LastName, u.Dni, u.Email, u.Telephone, u.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// isIdPresent retorna True si la sql.Row dada tiene solo un id.
// Retorna False cuando el row.Scan(&id) devuelve un error, porque la query no tuvo resultados.
func isIdPresent(row *sql.Row) bool {
	var id string
	err := row.Scan(&id)
	return err == nil
}

// scanUserFromRow devuelve el domain.User leido al ejecutar Scan() tanto en sql.Row como en sql.Rows.
// Nota: [R interface{ Scan(...any) error }] hace a la función genérica para cualquier tipo R que cumpla esa interfaz.
func scanUserFromRow[R interface{ Scan(...any) error }](row R) (domain.User, error) {
	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Dni,
		&user.Email,
		&user.Telephone,
	)
	return user, err
}
