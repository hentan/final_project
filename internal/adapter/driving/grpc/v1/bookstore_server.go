package v1

import (
	"context"
	"fmt"
	"github.com/hentan/final_project/internal/logger"
	"github.com/hentan/final_project/internal/models"
	"github.com/hentan/final_project/internal/usecase"
	"github.com/hentan/internal_api_books/gen/go/books"
)

type bookstoreServer struct {
	books.BookServiceServer
	bookUseCase *usecase.UseCase
}

func NewBookstoreServer(bookUseCase *usecase.UseCase) *bookstoreServer {
	return &bookstoreServer{
		bookUseCase: bookUseCase,
	}
}

func (s *bookstoreServer) CreateBook(ctx context.Context, req *books.CreateBookRequest) (res *books.CreateBookResponse, err error) {
	book := models.Book{
		Title:    req.GetTitle(),
		AuthorID: int(req.GetAuthorId()),
		ISBN:     req.GetIsbn(),
		Year:     int(req.GetYear()),
	}
	id, err := s.bookUseCase.BookRepo().InsertBook(book)
	log := logger.GetLogger()
	if err != nil {
		wrapError := fmt.Errorf("internal/adapter/driver/grpc/v1/bookstore_server.go CreateBook error in grpc, %w", err)
		log.Error(wrapError.Error())
		return nil, err
	}

	book.ID = int(id)
	res.Book.Author = book.Author
	res.Book.Title = book.Title
	res.Book.AuthorId = int32(book.AuthorID)
	res.Book.Isbn = book.ISBN
	res.Book.Year = int32(book.Year)

	return res, nil
}
