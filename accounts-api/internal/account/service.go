package account

import (
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
	"errors"
)

var (
	ErrInternal       = errors.New("internal server error")
	ErrGettingAccount = errors.New("error getting account")
	ErrNotFound       = errors.New("account not found")
	ErrAliasConflict  = errors.New("requested alias already exists")
	ErrMissingFields  = errors.New("account is missing fields")
)

type Service interface {
	GetByID(id int) (domain.Account, error)
	GetByAliasOrCvu(aliasOrCvu string) (domain.Account, error)
	GetByUserID(userID string) (domain.Account, error)
	GetAll() ([]domain.Account, error)
	GetLastTransactions(id int, limit uint) ([]domain.Transaction, error)
	Create(domain.Account) (domain.Account, error)
	UpdateAlias(id int, alias string) (domain.Account, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Account, error) {
	account, err := s.r.GetByID(id)
	if err != nil {
		return domain.Account{}, ErrNotFound
	}
	return account, nil
}

func (s *service) GetByAliasOrCvu(aliasOrCvu string) (domain.Account, error) {
	account, err := s.r.GetByAliasOrCvu(aliasOrCvu)
	if err != nil {
		return domain.Account{}, ErrNotFound
	}
	return account, nil
}

func (s *service) GetByUserID(userID string) (domain.Account, error) {
	account, err := s.r.GetByUserID(userID)
	if err != nil {
		return domain.Account{}, ErrNotFound
	}
	return account, nil
}

func (s *service) GetAll() ([]domain.Account, error) {
	accounts, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *service) GetLastTransactions(id int, limit uint) ([]domain.Transaction, error) {
	if limit > 0 {
		transactions, err := s.r.GetTransactionsWithLimit(id, limit)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	transactions, err := s.r.GetLastTransactions(id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) Create(account domain.Account) (domain.Account, error) {
	if utils.IsBlank(account.UserID) {
		return domain.Account{}, ErrMissingFields
	}
	account.Amount = 0
	account.Cvu = s.generateCvu()
	account.Alias = s.generateAlias()

	accountID, err := s.r.Create(account)
	if err != nil {
		return domain.Account{}, ErrInternal
	}
	accountCreated, err := s.r.GetByID(accountID)
	if err != nil {
		return domain.Account{}, ErrGettingAccount
	}

	return accountCreated, nil
}

func (s *service) UpdateAlias(id int, alias string) (domain.Account, error) {

	if !s.r.ExistsByID(id) {
		return domain.Account{}, ErrNotFound
	}
	if s.r.ExistsByAlias(alias) {
		return domain.Account{}, ErrAliasConflict
	}

	err := s.r.UpdateAlias(id, alias)
	if err != nil {
		return domain.Account{}, ErrInternal
	}

	account, err := s.GetByID(id)
	if err != nil {
		return domain.Account{}, ErrGettingAccount
	}
	return account, nil
}

func (s *service) generateCvu() string {
	var cvu string
	for {
		cvu = utils.GenerateCvu()
		if !s.r.ExistsByCvu(cvu) {
			break
		}
	}
	return cvu
}

func (s *service) generateAlias() string {
	var alias string
	for {
		alias = utils.GenerateAlias()
		if !s.r.ExistsByAlias(alias) {
			break
		}
	}
	return alias
}
