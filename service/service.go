package service

import (
	"github.com/aholake/book-store/model"
	"github.com/aholake/book-store/repo"
)

func AddBook(book model.Book) model.Book {
	return repo.AddBook(book)
}

func FindById(id int64) *model.Book {
	return repo.FindById(id)
}

func FindAll() []model.Book {
	return repo.FindAll()
}
