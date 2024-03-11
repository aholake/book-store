package server

import (
	"context"

	"github.com/aholake/book-store/model"
	pb "github.com/aholake/book-store/proto"
	"github.com/aholake/book-store/repo"
	"github.com/aholake/book-store/service"
	"github.com/aholake/book-store/util"
)

type Server struct {
	pb.UnimplementedBookServiceServer
}

func (s Server) AddBook(ctx context.Context, b *pb.Book) (*pb.Book, error) {
	bookModel := model.Book{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
	}

	updatedBook := service.AddBook(bookModel)
	return &pb.Book{
		Id:     updatedBook.Id,
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
	}, nil
}

func (s Server) Get(ctx context.Context, id *pb.BookIdRequest) (*pb.Book, error) {
	book := service.FindById(id.Id)
	if book == nil {
		return nil, nil
	}
	return &pb.Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
	}, nil
}

func (s Server) GetAll(context.Context, *pb.EmptyMessage) (*pb.GetAllResponse, error) {
	books := repo.FindAll()
	response := []*pb.Book{}
	for _, b := range books {
		res := util.Convert(b)
		response = append(response, &res)
	}
	return &pb.GetAllResponse{
		Books: response,
	}, nil
}
