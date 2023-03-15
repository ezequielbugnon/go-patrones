package remote

import (
	"time"

	"github.com/ezequielbugnon/go-patrones/proxy/book"
)

type Data struct {
	books    book.Books
	server   string
	port     uint16
	user     string
	password string
}

func New(server string, port uint16, user, password string) *Data {
	d := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}

	d.load()
	return d
}

func (d *Data) ById(ID uint) book.Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}

	return book.Book{}
}

func (d *Data) All() book.Books {
	time.Sleep(4 * time.Second)
	return d.books
}

func (d *Data) load() {
	d.books = make(book.Books, 0, 2)

	d.books = append(d.books,
		book.Book{
			ID:     0,
			Name:   "",
			Author: "",
		}, book.Book{
			ID:     1,
			Name:   "",
			Author: "",
		})
}
