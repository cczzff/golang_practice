package auth

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"golang_practice/core/auth"
	"errors"
)

type (
	DBModel struct {
		db *xorm.Engine
	}
)

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
	return
}
