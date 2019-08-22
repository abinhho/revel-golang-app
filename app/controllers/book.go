package controllers

import (
	"revel-golang-app/app/models"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"github.com/revel/revel"
	"fmt"
)

type ApiBook struct {
	ApiController
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


func (c *ApiBook) Create() revel.Result {
	book := &models.Book{}
	fmt.Println(c)
	// c.BindParams(book)

	var hasError = book.Validate()
	if hasError != nil {
		return c.Response(&ErrorResponse{ERR_VALIDATE, c.ErrorMessage(ERR_VALIDATE)})
	}

	err := c.Txn.Insert(book)
	if err != nil {
		panic(err)
	}

	return c.App.RenderJson(&Response{OK, book})
}
