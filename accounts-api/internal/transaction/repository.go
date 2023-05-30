package transaction

import (
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	queryCreate               = "INSERT INTO transactions (amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type) VALUES (?, ?, ?, ?, ?, ?, ?)"
	querySelectByID           = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE id = ?"
	querySelectAllByAccountID = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = ?"

	querySelectAllByAmount     = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND amount between '%s' and '%s'"
	querySelectAllByAmountFrom = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND amount >= '%s'"
	querySelectAllByAmountTo   = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND amount <= '%s'"

	querySelectAllByDate     = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND transaction_date between '%s' and '%s'"
	querySelectAllByDateFrom = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND transaction_date >= '%s'"
	querySelectAllByDateTo   = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND transaction_date <= '%s'"

	querySelectAllByType = "SELECT id, amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions WHERE account_id = %d AND transaction_type='%s'"
)

type Repository interface {
	GetByID(id int) (domain.Transaction, error)
	Create(t domain.Transaction) (int, error)
	GetAllByAccountID(id int) ([]domain.Transaction, error)
	FilterTransactions(accountID int, filter domain.Filter) ([]domain.Transaction, utils.ApiError)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(id int) (domain.Transaction, error) {
	row := r.db.QueryRow(querySelectByID, id)
	transaction, err := scanTransactionFromRow(row)
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) GetAllByAccountID(id int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	rows, err := r.db.Query(querySelectAllByAccountID, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transaction, _ := scanTransactionFromRow(rows)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) FilterTransactions(accountID int, filter domain.Filter) ([]domain.Transaction, utils.ApiError) {
	transactions := []domain.Transaction{}
	query, apiErr := getFilterQuery(accountID, filter)
	if apiErr != nil {
		return nil, apiErr
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, utils.NewInternalServerError(ErrInternal.Error(), err)
	}

	for rows.Next() {
		transaction, _ := scanTransactionFromRow(rows)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) Create(t domain.Transaction) (int, error) {
	stmt, err := r.db.Prepare(queryCreate)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(t.Amount, t.TransactionDate, t.Description, t.OriginCvu, t.DestinationCvu, t.AccountID, t.TransactionType)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func getFilterQuery(accountID int, filter domain.Filter) (string, utils.ApiError) {
	if strings.ToLower(filter.Type) == "ingreso" || strings.ToLower(filter.Type) == "egreso" {
		return fmt.Sprintf(querySelectAllByType, accountID, filter.Type), nil
	}

	if filter.Type == "date" {
		switch {
		case filter.From != "" && filter.To != "":
			return fmt.Sprintf(querySelectAllByDate, accountID, filter.From, filter.To), nil
		case filter.From != "" && filter.To == "":
			return fmt.Sprintf(querySelectAllByDateFrom, accountID, filter.From), nil
		case filter.From == "" && filter.To != "":
			return fmt.Sprintf(querySelectAllByDateTo, accountID, filter.To), nil
		default:
			return "", utils.NewInternalServerError("error generating query", errors.New(""))
		}
	}

	if filter.Type == "amount" {
		switch {
		case filter.From != "" && filter.To != "":
			return fmt.Sprintf(querySelectAllByAmount, accountID, filter.From, filter.To), nil
		case filter.From != "" && filter.To == "":
			return fmt.Sprintf(querySelectAllByAmountFrom, accountID, filter.From), nil
		case filter.From == "" && filter.To != "":
			return fmt.Sprintf(querySelectAllByAmountTo, accountID, filter.To), nil
		default:
			return "", utils.NewInternalServerError("error generating query", errors.New(""))
		}
	}

	return "", utils.NewBadRequestError("invalid filter", errors.New(""))
}

func scanTransactionFromRow[R interface{ Scan(...any) error }](row R) (domain.Transaction, error) {
	var transaction domain.Transaction
	err := row.Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.TransactionDate,
		&transaction.Description,
		&transaction.OriginCvu,
		&transaction.DestinationCvu,
		&transaction.AccountID,
		&transaction.TransactionType,
	)
	return transaction, err
}
