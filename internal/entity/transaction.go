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

	if accountFrom == nil {
		return nil, errors.New("Account from cannot be empty")
	}

	if accountTo == nil {
		return nil, errors.New("Account to cannot be empty")
	}

	if amount <= 0 {
		return nil, errors.New("Amount cannot be less than or equal to 0")
	}

	return &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
	}, nil
}
