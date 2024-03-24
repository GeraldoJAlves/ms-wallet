package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	if client == nil {
		return nil, errors.New("invalid client")
	}

	return &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (a *Account) Credit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount to credit")
	}

	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Debit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount to credit")
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return nil
}
