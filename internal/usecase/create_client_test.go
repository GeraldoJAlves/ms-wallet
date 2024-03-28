package usecase

import (
	"testing"

	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)

	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUseCase(t *testing.T) {

	input := CreateClientInputDTO{
		Name:  "john",
		Email: "john@j.com",
	}

	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateClientUseCase(clientGatewayMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	assert.Equal(t, output.Email, input.Email)
	assert.Equal(t, output.Name, input.Name)
	assert.NotNil(t, output.CreatedAt)
	assert.NotNil(t, output.UpdatedAt)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
