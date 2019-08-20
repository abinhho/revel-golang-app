package models

import (
	"errors"
	"github.com/revel/revel"
)

type BookData struct {
	ID int `json: "id"`
	BookName string `json:bookname`
	BookCatId int `json:bookcatid`
	Created int64 `json:-`
	Updated int64 `json:-`
}

func (book *BookData) Validate() error {
	var v revel.Validation

	v.Check(
		book.BookName,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{255},
	)

	if v.HasErrors() {
		return errors.New("Book name is not valid")
	}

	return nil
}
