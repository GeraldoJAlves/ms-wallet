package usecase

import (
	"testing"

	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (t *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase(t *testing.T) {

	transactionGateway := &TransactionGatewayMock{}
	transactionGateway.On("Create", mock.Anything).Return(nil)

	client, _ := entity.NewClient("john", "j@j.com")
	account, _ := entity.NewAccount(client)
	account.Credit(40)

	account2, _ := entity.NewAccount(client)
	account2.Credit(20)

	accountGateway := &AccountGatewayMock{}
	accountGateway.On("FindByID", "1").Return(account, nil)
	accountGateway.On("FindByID", "2").Return(account2, nil)

	uc := &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
	}

	input := CreateTransactionInputDTO{
		AccountIDFrom: "1",
		AccountIDTo:   "2",
		Amount:        20,
	}

	output, err := uc.Execute(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, 20.0, account.Balance)
	assert.Equal(t, 40.0, account2.Balance)
}
