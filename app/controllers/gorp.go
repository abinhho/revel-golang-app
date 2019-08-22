package controllers

import (
	"revel-golang-app/app/models"
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/modules/db/app"
	r "github.com/revel/revel"
	"log"
	"os"
)

var (
	Dbm *gorp.DbMap
	Txn *gorp.Transaction
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	Dbm.AddTableWithName(models.Book{}, "book").SetKeys(true, "ID")
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if Txn == nil {
		return nil
	}
	if err := Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return nil
}

func (c *GorpController) Rollback1() r.Result {
	if Txn == nil {
		return nil
	}
	if err := Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return nil
}