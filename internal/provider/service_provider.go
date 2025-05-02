package provider

import "github.com/hentan/internal_api_books/gen/go/books"

type ServiceProvider struct {
	bookService *books.BookServiceServer
}
