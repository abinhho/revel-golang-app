package models

import (
	"errors"
	"github.com/revel/revel"
)

type Book struct {
	Id       int64                        `id`
	BookName string                       `bookname`
    Created  int64                        `created`
    Updated  int64                        `modified`
}

func (book *Book) Validate() error {
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
