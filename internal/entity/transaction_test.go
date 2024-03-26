package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("john", "j@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("mary", "m@m.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.NotNil(t, transaction)
	assert.Nil(t, err)

	assert.Equal(t, 900.00, account1.Balance)
	assert.Equal(t, 1100.00, account2.Balance)
}
