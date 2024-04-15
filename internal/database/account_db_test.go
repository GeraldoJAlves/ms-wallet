package database

import (
	"database/sql"
	"testing"

	"github.com/geraldojalves/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	clientDB  *ClientDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance decimal(10, 2), created_at date)")
	s.Nil(err)

	s.accountDB = NewAccountDB(db)
	s.clientDB = NewClientDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	_, err := s.db.Exec("DROP TABLE accounts")
	s.Nil(err)
	_, err = s.db.Exec("DROP TABLE clients")
	s.Nil(err)
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {

	client, err := entity.NewClient("john", "j@j.com")
	s.Nil(err)
	account, err := entity.NewAccount(client)
	s.Nil(err)

	err = s.clientDB.Save(client)
	s.Nil(err)

	err = s.accountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)

	s.Equal(account.ID, accountDB.ID)
}
