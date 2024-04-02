package usecase

import (
	"testing"

	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (a *AccountGatewayMock) Save(account *entity.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUsecase(t *testing.T) {

	client, _ := entity.NewClient("john", "j@j.com")
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Get", mock.Anything).Return(client, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	usecase := &CreateAccountUseCase{
		AccountGateway: accountGatewayMock,
		ClientGateway:  clientGatewayMock,
	}

	input := &CreateAccountInputDTO{
		ClientID: "1",
	}

	output, error := usecase.Execute(input)

	assert.Nil(t, error)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)

}
