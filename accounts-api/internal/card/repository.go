package card

import (
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
)

var (
	queryGetAllByAccountID       = "SELECT id, card_number, expiration_date, owner, security_code, brand, account_id FROM cards WHERE account_id = ?"
	queryGetByID                 = "SELECT id, card_number, expiration_date, owner, security_code, brand, account_id FROM cards WHERE id = ?"
	queryGetByCardNumber         = "SELECT account_id FROM cards WHERE card_number = ?"
	queryGetByAccIDAndCardNumber = "SELECT id, card_number, expiration_date, owner, security_code, brand, account_id FROM cards WHERE account_id = ? and card_number = ?"
	queryCreate                  = "INSERT INTO cards (card_number, expiration_date, owner, security_code, brand, account_id) VALUES (?, ?, ?, ?, ?, ?)"
	queryDeleteById              = "DELETE FROM cards WHERE account_id=? AND id=?"
)

type Repository interface {
	GetAllByAccountID(accountID int) ([]domain.Card, utils.ApiError)
	GetByCardNumber(id string) utils.ApiError
	GetByAccIDAndCardNumber(accountID int, cardNumber string) (domain.Card, utils.ApiError)
	GetByID(id int) (domain.Card, utils.ApiError)
	Create(user domain.Card) (int, utils.ApiError)
	Delete(idAccount, idCard int) utils.ApiError
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllByAccountID(accountID int) ([]domain.Card, utils.ApiError) {
	rows, err := r.db.Query(queryGetAllByAccountID, accountID)
	if err != nil {
		return nil, utils.NewInternalServerError(ErrInternal, err)
	}

	var cards []domain.Card

	for rows.Next() {
		c := domain.Card{}
		err = rows.Scan(&c.ID, &c.CardNumber, &c.ExpirationDate, &c.Owner, &c.SecurityCode, &c.Brand, &c.AccountID)
		if err != nil {
			return nil, utils.NewInternalServerError(ErrInternal, err)
		}
		cards = append(cards, c)
	}

	if len(cards) == 0 {
		return nil, utils.NewNotFoundError("the account doesn't have any cards associated", errors.New("not found"))
	}

	return cards, nil
}

func (r *repository) GetByID(id int) (domain.Card, utils.ApiError) {
	row := r.db.QueryRow(queryGetByID, id)

	c := domain.Card{}
	err := row.Scan(&c.ID, &c.CardNumber, &c.ExpirationDate, &c.Owner, &c.SecurityCode, &c.Brand, &c.AccountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Card{}, utils.NewNotFoundError(ErrNotFound, err)
		}
		return domain.Card{}, utils.NewInternalServerError(ErrInternal, err)
	}
	return c, nil
}

func (r *repository) GetByAccIDAndCardNumber(accountID int, cardNumber string) (domain.Card, utils.ApiError) {
	row := r.db.QueryRow(queryGetByAccIDAndCardNumber, accountID, cardNumber)

	c := domain.Card{}
	err := row.Scan(&c.ID, &c.CardNumber, &c.ExpirationDate, &c.Owner, &c.SecurityCode, &c.Brand, &c.AccountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Card{}, utils.NewNotFoundError(ErrNotFound, err)
		}
		return domain.Card{}, utils.NewInternalServerError(ErrInternal, err)
	}
	return c, nil
}

func (r *repository) GetByCardNumber(id string) utils.ApiError {
	row := r.db.QueryRow(queryGetByCardNumber, id)

	var accountID int
	if err := row.Scan(&accountID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return utils.NewInternalServerError(ErrInternal, err)
	}

	return utils.NewConflictError(fmt.Sprintf(ErrConflict, accountID), errors.New("conflict"))
}

func (r *repository) Create(c domain.Card) (int, utils.ApiError) {
	stmt, err := r.db.Prepare(queryCreate)
	if err != nil {
		return 0, utils.NewInternalServerError(ErrInternal, err)
	}

	res, err := stmt.Exec(c.CardNumber, c.ExpirationDate, c.Owner, c.SecurityCode, c.Brand, c.AccountID)
	if err != nil {
		return 0, utils.NewInternalServerError(ErrInternal, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, utils.NewInternalServerError(ErrInternal, err)
	}

	return int(id), nil
}

func (r *repository) Delete(idAccount, idCard int) utils.ApiError {

	stmt, err := r.db.Prepare(queryDeleteById)
	if err != nil {
		return utils.NewInternalServerError(ErrInternal, err)
	}

	res, err := stmt.Exec(idAccount, idCard)
	if err != nil {
		return utils.NewInternalServerError(ErrInternal, err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return utils.NewInternalServerError(ErrInternal, err)
	}

	if affect < 1 {
		return utils.NewNotFoundError(fmt.Sprintf(ErrNotFound, idCard), errors.New(""))
	}

	return nil
}
