package models

import (
	"github.com/revel/revel"
	"fmt"
)

type Book struct {
	Id       int64                        `id`
	BookName string                       `bookname`
    Created  int64                        `created`
    Updated  int64                        `modified`
}

func (book *Book) Validate(v *revel.Validation) {
	fmt.Println(book)
	v.Required(book.BookName)
	v.MaxSize(book.BookName, 255).MessageKey("BookName min length is 255")
	v.MinSize(book.BookName, 4)

	// return v
	// v.Check(
	// 	book.BookName,
	// 	revel.Required{},
	// 	revel.MinSize{4},
	// 	revel.MaxSize{255},
	// )

	// if v.HasErrors() {
	// 	return errors.New("Book name is not valid")
	// }

	// return nil
}
