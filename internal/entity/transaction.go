package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("account from cannot be empty")
	}

	if t.AccountTo == nil {
		return errors.New("account to cannot be empty")
	}

	if t.Amount <= 0 {
		return errors.New("amount cannot be less than or equal to 0")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insuficient funds")
	}

	return nil
}
