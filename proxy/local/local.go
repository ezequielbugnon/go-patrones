package local

import (
	"github.com/ezequielbugnon/go-patrones/proxy/book"
	"github.com/ezequielbugnon/go-patrones/proxy/remote"
)

type Local struct {
	Remote *remote.Data
	cache  book.Books
}

func New(r *remote.Data) Proxy {
	return &Local{
		Remote: r,
		cache:  make(book.Books, 0),
	}
}

func (l *Local) GetByID(ID uint) book.Book {
	for _, v := range l.cache {
		if v.ID == ID {
			return v
		}
	}
	b := l.Remote.ById(ID)
	l.cache = append(l.cache, b)

	return b
}

func (l *Local) GetAll() book.Books {
	return l.Remote.All()
}
