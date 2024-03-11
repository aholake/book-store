package repo

import (
	"log"

	"github.com/aholake/book-store/database"
	"github.com/aholake/book-store/model"
)

func AddBook(book model.Book) model.Book {
	result, err := database.DB.Exec("INSERT INTO book (title, author) VALUES (?, ?)",
		book.Title, book.Author,
	)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	returnedBook := model.Book{}
	err = database.DB.QueryRow("SELECT * FROM book WHERE id=?", lastInsertID).Scan(&returnedBook.Id, &returnedBook.Title, &returnedBook.Author)
	if err != nil {
		log.Fatal(err)
	}

	return returnedBook
}

func FindById(id int64) *model.Book {
	returnedBook := model.Book{}
	err := database.DB.QueryRow("SELECT * FROM book WHERE id=?", id).Scan(&returnedBook.Id, &returnedBook.Title, &returnedBook.Author)
	if err != nil {
		log.Print("No records found")
		return nil
	}
	return &returnedBook
}

func FindAll() []model.Book {
	books := []model.Book{}
	row, err := database.DB.Query("SELECT * FROM book")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		book := model.Book{}
		err = row.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books

}
