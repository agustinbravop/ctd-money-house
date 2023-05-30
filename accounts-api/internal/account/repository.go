package account

import (
	"accounts-api/internal/domain"
	"database/sql"
)

var (
	queryCreate                       = "INSERT INTO accounts (cvu, alias, amount, user_id) VALUES (?, ?, ?, ?)"
	queryExistsByID                   = "SELECT id FROM accounts WHERE id=?"
	queryExistsByCvu                  = "SELECT id FROM accounts WHERE cvu=?"
	queryExistsByAlias                = "SELECT id FROM accounts WHERE alias=?"
	queryUpdateAlias                  = "UPDATE accounts SET alias=? WHERE id=?"
	queryUpdateAmount                 = "UPDATE accounts SET amount=? WHERE id=?"
	querySelectByID                   = "SELECT id, cvu, alias, amount, user_id FROM accounts WHERE id = ?"
	querySelectByCvu                  = "SELECT id, cvu, alias, amount, user_id FROM accounts WHERE cvu = ?"
	querySelectByAliasOrCvu           = "SELECT id, cvu, alias, amount, user_id FROM accounts WHERE cvu = ? OR alias = ?"
	querySelectByUserID               = "SELECT id, cvu, alias, amount, user_id FROM accounts WHERE user_id = ?"
	queryGetLastTransactions          = "SELECT t.id, t.amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions t JOIN accounts a ON t.account_id = a.id WHERE t.account_id = ? ORDER BY transaction_date DESC"
	queryGetLastTransactionsWithLimit = "SELECT t.id, t.amount, transaction_date, description, origin_cvu, destination_cvu, account_id, transaction_type FROM transactions t JOIN accounts a ON t.account_id = a.id WHERE t.account_id = ? ORDER BY transaction_date DESC LIMIT ?"
)

type Repository interface {
	GetByID(id int) (domain.Account, error)
	GetByCvu(cvu string) (domain.Account, error)
	GetByAliasOrCvu(aliasOrCvu string) (domain.Account, error)
	GetByUserID(userID string) (domain.Account, error)
	GetAll() ([]domain.Account, error)
	GetLastTransactions(id int) ([]domain.Transaction, error)
	GetTransactionsWithLimit(id int, limit uint) ([]domain.Transaction, error)
	Create(account domain.Account) (int, error)
	UpdateAmount(account domain.Account) error
	UpdateAlias(id int, alias string) error
	ExistsByID(id int) bool
	ExistsByCvu(cvu string) bool
	ExistsByAlias(alias string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(id int) (domain.Account, error) {
	row := r.db.QueryRow(querySelectByID, id)
	account, err := scanAccountFromRow(row)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (r *repository) GetByCvu(cvu string) (domain.Account, error) {
	row := r.db.QueryRow(querySelectByCvu, cvu)
	account, err := scanAccountFromRow(row)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (r *repository) GetByAliasOrCvu(aliasOrCvu string) (domain.Account, error) {
	row := r.db.QueryRow(querySelectByAliasOrCvu, aliasOrCvu, aliasOrCvu)
	account, err := scanAccountFromRow(row)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (r *repository) GetByUserID(userID string) (domain.Account, error) {
	row := r.db.QueryRow(querySelectByUserID, userID)
	account, err := scanAccountFromRow(row)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (r *repository) GetAll() ([]domain.Account, error) {
	var accounts []domain.Account
	rows, err := r.db.Query("SELECT id, cvu, alias, amount, user_id FROM accounts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		account, _ := scanAccountFromRow(rows)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *repository) GetLastTransactions(id int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	rows, err := r.db.Query(queryGetLastTransactions, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction domain.Transaction
		_ = rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.TransactionDate,
			&transaction.Description,
			&transaction.OriginCvu,
			&transaction.DestinationCvu,
			&transaction.AccountID,
			&transaction.TransactionType,
		)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) GetTransactionsWithLimit(id int, limit uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	rows, err := r.db.Query(queryGetLastTransactionsWithLimit, id, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction domain.Transaction
		_ = rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.TransactionDate,
			&transaction.Description,
			&transaction.OriginCvu,
			&transaction.DestinationCvu,
			&transaction.AccountID,
			&transaction.TransactionType,
		)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) Create(a domain.Account) (int, error) {
	stmt, err := r.db.Prepare(queryCreate)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(a.Cvu, a.Alias, a.Amount, a.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) UpdateAlias(id int, alias string) error {
	stmt, err := r.db.Prepare(queryUpdateAlias)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(alias, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateAmount(a domain.Account) error {
	stmt, err := r.db.Prepare(queryUpdateAmount)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(a.Amount, a.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// ExistsByID retorna True cuando ya existe un domain.Account con ese ID en la base de datos.
func (r *repository) ExistsByID(id int) bool {
	row := r.db.QueryRow(queryExistsByID, id)
	return isIdPresent(row)
}

// ExistsByCvu retorna True cuando ya existe un domain.Account con ese cvu en la base de datos.
func (r *repository) ExistsByCvu(cvu string) bool {
	row := r.db.QueryRow(queryExistsByCvu, cvu)
	return isIdPresent(row)
}

// ExistsByAlias retorna True cuando ya existe un domain.Account con ese alias en la base de datos.
func (r *repository) ExistsByAlias(alias string) bool {
	row := r.db.QueryRow(queryExistsByAlias, alias)
	return isIdPresent(row)
}

// isIdPresent retorna True si la sql.Row dada tiene solo un id.
// Retorna False cuando el row.Scan(&id) devuelve un error, porque la query no tuvo resultados.
func isIdPresent(row *sql.Row) bool {
	var id int
	err := row.Scan(&id)
	return err == nil
}

func scanAccountFromRow[R interface{ Scan(...any) error }](row R) (domain.Account, error) {
	var account domain.Account
	err := row.Scan(
		&account.ID,
		&account.Cvu,
		&account.Alias,
		&account.Amount,
		&account.UserID,
	)
	return account, err
}
