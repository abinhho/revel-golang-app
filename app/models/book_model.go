package models

import (
	"error"
	"regexp"
	"gihub.com/revel/revel"
)

type BookModel struct {
	ID int `json: "id"`
	BookName string `json:bookname`
	BookCatId int `json:bookcatid`
	Created int64 `json:-`
	Updated int64 `json:-`
}

func (book *BookModel) Validate() error {
	var v revel.Validation

	v.check(
		book.BookName,
		revel.required{},
		revel.MinSize{4},
		revel.MaxSize{255},
	)

	if v.HasError() {
		return errors.New("Book name is not valid")
	}

	return nil
}
