package util

import (
	"github.com/aholake/book-store/model"
	pb "github.com/aholake/book-store/proto"
)

func Convert(bookModel model.Book) pb.Book {
	return pb.Book{
		Id:     bookModel.Id,
		Title:  bookModel.Title,
		Author: bookModel.Author,
	}
}
