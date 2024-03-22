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
