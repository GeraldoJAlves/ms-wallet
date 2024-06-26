package database

import (
	"database/sql"

	"github.com/geraldojalves/ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	account := &entity.Account{}
	client := &entity.Client{}
	account.Client = client

	stmt, err := a.DB.Prepare("SELECT a.id, a.balance, a.created_at, a.client_id, c.name, c.email, c.created_at FROM accounts a JOIN clients c ON (a.client_id = c.id)  WHERE a.id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	if err := row.Scan(&account.ID, &account.Balance, &account.CreatedAt, &client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
