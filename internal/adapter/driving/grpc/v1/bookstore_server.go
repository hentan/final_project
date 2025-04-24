package v1

import (
	"context"
	"fmt"
	"github.com/hentan/internal_api_books/gen/go/books"
)

type bookstoreServer struct {
	books.BookServiceServer
}

func NewBookstoreServer() books.BookServiceServer {
	return &bookstoreServer{}
}

func (s *bookstoreServer) CreateBook(ctx context.Context, req *books.CreateBookRequest) (*books.CreateBookResponse, error) {
	fmt.Println("CreateBook called")
	return nil, nil
}
