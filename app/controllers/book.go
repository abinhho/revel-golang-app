package controllers

import (
	"revel-golang-app/app/models"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"github.com/revel/revel"
	"fmt"
	"strings"
	"time"
)

type ApiBook struct {
	ApiController
	GorpController
}

func BindJsonParams(i io.Reader, s interface{}) error {
	bytes, err := ioutil.ReadAll(i)
	if err != nil {
		return errors.New("can't read request data.")
	}

	if len(bytes) == 0 {
		return errors.New("data is nil")
	}

	return json.Unmarshal(bytes, s)
}


func (c *ApiBook) Create(bookName string) revel.Result {
	book := &models.Book{
		BookName: bookName,
		Updated: time.Now().Unix(),
		Created: time.Now().Unix(),
	}

	v := c.ApiController.Validation
	book.Validate(v)
	var errors []string
	for _, e := range v.Errors {
		errors = append(errors, strings.TrimSuffix(e.Message, "\n"))
	}

	fmt.Println(strings.Join(errors, ","))
	if v.HasErrors() {
		// return c.Response(&ErrorResponse{ERR_VALIDATE, c.ErrorMessage(ERR_VALIDATE)})
		return c.Response(&ErrorResponse{ERR_VALIDATE, strings.Join(errors, ", ")})
	}

	err := Txn.Insert(book)
	if err != nil {
		panic(err)
	}

	return c.Response(&Response{OK, book})
}

func (c *ApiBook) Show(id int64) revel.Result {
	var book models.Book

	bookData, err := Dbm.Get(book, id)
	if err != nil {
		panic(err)
	}

	return c.Response(&Response{OK, bookData})
}

func (c *ApiBook) List() revel.Result {
	var books []models.Book

	_, err := Dbm.Select(&books, "SELECT * FROM books LIMIT 0, 50")
	if err != nil {
		panic(err)
	}

	if len(books) == 0 {
		return c.Response(&ErrorResponse{ERR_VALIDATE, c.ErrorMessage(WARN_NOT_FOUND)})
	}

	fmt.Println(books)

	return c.Response(&Response{OK, books})
}
