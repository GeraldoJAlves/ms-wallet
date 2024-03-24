package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("john doe", "john@example.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "john doe", client.Name)
	assert.Equal(t, "john@example.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("john doe", "john@example.com")

	err := client.Update("john silva", "silva@example.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.Equal(t, "john silva", client.Name)
	assert.Equal(t, "silva@example.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("john doe", "john@example.com")

	err := client.Update("", "")

	assert.NotNil(t, err)

	assert.Equal(t, "john doe", client.Name)
	assert.Equal(t, "john@example.com", client.Email)
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("john doe", "john@example.com")

	assert.Equal(t, 0, len(client.Accounts))

	account, _ := NewAccount(client)
	client.AddAccount(account)

	assert.Equal(t, 1, len(client.Accounts))
}
