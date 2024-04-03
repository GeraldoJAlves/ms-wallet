package usecase

import (
	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/geraldojalves/ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
}

func NewCreateTransactionUseCase(accountGateway gateway.AccountGateway, transactionGateway gateway.TransactionGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {

	accountFrom, err := c.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		TransactionID: transaction.ID,
	}, nil
}
