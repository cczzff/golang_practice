package auth

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"golang_practice/core/auth"
	"errors"
	"github.com/satori/go.uuid"
)

type (
	DBModel struct {
		db *xorm.Engine
	}
)

type authToken struct {
	Id     int64
	UserId int64 `xorm:"unique"`
	Token  string
}

func NewDBModel(dbDriver, dbSource string) (m *DBModel, err error) {
	m = &DBModel{}
	m.db, err = xorm.NewEngine(dbDriver, dbSource)
	m.db.ShowSQL(true)
	return
}

func (m *DBModel) addUser(username string, password string) (account *core_auth.Account, err error) {
	account = &core_auth.Account{
		Username: username,
		Password: password,
	}
	_, err = m.db.Insert(account)

	if err != nil {
		return
	}
	has, err := m.db.Get(account)
	if err != nil {
		return
	}
	if has == false {
		err = errors.New("failed to get account")
	}
	token, err := uuid.NewV4()
	if err != nil {
		return
	}
	_, err = m.db.Insert(&authToken{
		UserId: account.Id,
		Token:  token.String(),
	})

	if err != nil {
		return
	}

	return
}
