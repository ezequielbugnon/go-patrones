package local

import "github.com/ezequielbugnon/go-patrones/proxy/book"

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
