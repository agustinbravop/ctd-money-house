package transaction

import (
	"accounts-api/internal/account"
	"accounts-api/internal/card"
	"accounts-api/internal/clients"
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
	"errors"
	"time"
)

var (
	ErrInternal                   = errors.New("internal server error")
	ErrGettingTransaction         = errors.New("error getting created transaction")
	ErrNotFound                   = errors.New("transaction not found")
	ErrAccountNotFound            = errors.New("account not found")
	ErrCardNotFound               = errors.New("card not found")
	ErrExtractionFromCard         = errors.New("bank interface failed to extract from card")
	ErrDestinationAccountNotFound = errors.New("destination account not found")
	ErrMismatchedAccountID        = errors.New("accountID does not match with associated card's accountID")
	ErrInsufficientFunds          = errors.New("insufficient funds")
	ErrAmountCannotBeNegative     = errors.New("amount cannot be 0 or negative")
	ErrSameAccount                = errors.New("origin account and destination account must be different")
	ErrMisingFields               = errors.New("transaction is missing fields")
)

const (
	TypeIngress = "ingreso"
	TypeEgress  = "egreso"
)

type Service interface {
	GetByID(id int) (domain.Transaction, error)
	GetAllByAccountID(id int) ([]domain.Transaction, error)
	FilterTransactions(accountID int, filter domain.Filter) ([]domain.Transaction, utils.ApiError)
	Create(domain.Transaction) (domain.Transaction, error)
	DepositFromCard(accountID int, amount float64, cardID int) (domain.Transaction, error)
}

type service struct {
	r        Repository
	accRepo  account.Repository
	cardRepo card.Repository
}

func NewService(r Repository, ar account.Repository, cr card.Repository) Service {
	return &service{
		r:        r,
		accRepo:  ar,
		cardRepo: cr,
	}
}

func (s *service) GetByID(id int) (domain.Transaction, error) {
	transaction, err := s.r.GetByID(id)
	if err != nil {
		return domain.Transaction{}, ErrNotFound
	}
	return transaction, nil
}

func (s *service) GetAllByAccountID(id int) ([]domain.Transaction, error) {
	transactions, err := s.r.GetAllByAccountID(id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) FilterTransactions(accountID int, filter domain.Filter) ([]domain.Transaction, utils.ApiError) {
	transactions, err := s.r.FilterTransactions(accountID, filter)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) Create(transaction domain.Transaction) (domain.Transaction, error) {
	transaction.TransactionType = TypeEgress
	if transaction.Amount <= 0 {
		return domain.Transaction{}, ErrAmountCannotBeNegative
	}
	if utils.AnyBlank(transaction.DestinationCvu, transaction.OriginCvu) {
		return domain.Transaction{}, ErrMisingFields
	}
	originAccount, err := s.accRepo.GetByID(transaction.AccountID)
	if err != nil || originAccount.Cvu != transaction.OriginCvu {
		return domain.Transaction{}, ErrAccountNotFound
	}
	if originAccount.Amount < transaction.Amount {
		return domain.Transaction{}, ErrInsufficientFunds
	}

	destinationAccount, err := s.accRepo.GetByCvu(transaction.DestinationCvu)
	if err != nil {
		return domain.Transaction{}, ErrDestinationAccountNotFound
	}
	if originAccount.ID == destinationAccount.ID {
		return domain.Transaction{}, ErrSameAccount
	}

	originAccount.Amount -= transaction.Amount
	destinationAccount.Amount += transaction.Amount

	if err := s.accRepo.UpdateAmount(originAccount); err != nil {
		return domain.Transaction{}, ErrInternal
	}
	if err := s.accRepo.UpdateAmount(destinationAccount); err != nil {
		return domain.Transaction{}, ErrInternal
	}

	destinationTransaction := transaction
	destinationTransaction.TransactionType = TypeIngress
	destinationTransaction.AccountID = destinationAccount.ID

	_, err = s.r.Create(destinationTransaction)
	if err != nil {
		return domain.Transaction{}, ErrInternal
	}
	transactionID, err := s.r.Create(transaction)
	if err != nil {
		return domain.Transaction{}, ErrInternal
	}
	transactionCreated, err := s.r.GetByID(transactionID)
	if err != nil {
		return domain.Transaction{}, ErrGettingTransaction
	}
	return transactionCreated, nil
}

func (s *service) DepositFromCard(accountID int, amount float64, cardID int) (domain.Transaction, error) {
	if amount <= 0 {
		return domain.Transaction{}, ErrAmountCannotBeNegative
	}

	acc, err := s.accRepo.GetByID(accountID)
	if err != nil {
		return domain.Transaction{}, ErrAccountNotFound
	}

	card, apiErr := s.cardRepo.GetByID(cardID)
	if apiErr != nil {
		return domain.Transaction{}, ErrCardNotFound
	}
	if card.AccountID != accountID {
		return domain.Transaction{}, ErrMismatchedAccountID
	}

	extractionReq := clients.ExtractionReq{
		CardNumber:     card.CardNumber,
		ExpiryDate:     card.ExpirationDate,
		Owner:          card.Owner,
		SecurityCode:   card.SecurityCode,
		Amount:         amount,
		DestinationCvu: acc.Cvu,
	}

	extractionRes, err := clients.ExtractionFromCard(extractionReq)
	if err != nil {
		return domain.Transaction{}, ErrExtractionFromCard
	}

	transactionReq := domain.Transaction{
		Amount:          extractionRes.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		Description:     "deposito desde tarjeta",
		OriginCvu:       extractionRes.OriginCvu,
		DestinationCvu:  acc.Cvu,
		AccountID:       acc.ID,
		TransactionType: TypeIngress,
	}

	acc.Amount += extractionRes.Amount
	if err := s.accRepo.UpdateAmount(acc); err != nil {
		return domain.Transaction{}, ErrInternal
	}
	transactionID, err := s.r.Create(transactionReq)
	if err != nil {
		return domain.Transaction{}, ErrInternal
	}
	transaction, err := s.r.GetByID(transactionID)
	if err != nil {
		return domain.Transaction{}, ErrGettingTransaction
	}

	return transaction, nil
}
