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
	"fmt"
)

var (
	Dbm *gorp.DbMap
	Txn *gorp.Transaction
)

func InitDB() {
	db.Init()
	fmt.Println("Connect DB...");
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gotest")
	// if err != nil {
	// 	panic(err)
	// }

	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	// Dbm.AddTableWithName(models.Book{}, "book").SetKeys(true, "ID")
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

func (c *GorpController) Rollback() r.Result {
	if Txn == nil {
		return nil
	}
	if err := Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return nil
}

func (c *GorpController) CreateTable() r.Result {
	if Dbm == nil {
		return nil
	}
	Dbm.AddTableWithName(models.Book{}, "books").SetKeys(true, "Id")
	Dbm.CreateTables()
    fmt.Println("\n\nEnd AddTableWithName\n\n");
    return nil
}