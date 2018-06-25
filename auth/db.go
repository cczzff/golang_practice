package auth

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"

	"golang_practice/core/auth"
)

type (
	DBModel struct {
		db *xorm.Engine
	}
)

func NewDBModel(dbDriver, dbSource string) (m *DBModel, err error) {
	m = &DBModel{}
	m.db, err = xorm.NewEngine(dbDriver, dbSource)
	m.db.Sync2(new(core_auth.Account))
	return
}
