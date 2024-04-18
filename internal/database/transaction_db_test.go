package database

import (
	"database/sql"
	"testing"

	"github.com/geraldojalves/ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance decimal(10, 2), created_at date)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount decimal(10, 2), created_at date)")
	s.Nil(err)

	client, err := entity.NewClient("john", "j@j.com")
	s.Nil(err)
	s.client = client
	client2, err := entity.NewClient("john", "j@j.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom, err := entity.NewAccount(client)
	s.Nil(err)
	accountFrom.Balance = 200
	s.accountFrom = accountFrom

	accountTo, err := entity.NewAccount(client2)
	s.Nil(err)
	accountTo.Balance = 200
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	_, err := s.db.Exec("DROP TABLE accounts")
	s.Nil(err)
	_, err = s.db.Exec("DROP TABLE clients")
	s.Nil(err)
	_, err = s.db.Exec("DROP TABLE transactions")
	s.Nil(err)
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 90)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
