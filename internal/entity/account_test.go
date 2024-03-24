package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("john", "john@email.com")

	account, _ := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithInvalidClient(t *testing.T) {
	account, err := NewAccount(nil)

	assert.NotNil(t, err)
	assert.Nil(t, account)
}
