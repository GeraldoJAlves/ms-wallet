package usecase

import (
	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/geraldojalves/ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (c *CreateAccountUseCase) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := c.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	err = c.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
