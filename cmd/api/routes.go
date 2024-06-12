package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//mux роутер
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Get("/", app.Home)
	mux.Get("/books", app.AllBooks)
	mux.Get("/books/{id}", app.GetBook)
	mux.Put("/books/{id}", app.UpdateBook)
	mux.Post("/books", app.InsertBook)
	mux.Delete("/books/{id}", app.DeleteBook)
	mux.Get("/authors", app.AllAuthors)
	mux.Get("/authors/{id}", app.GetAuthor)
	mux.Put("/authors/{id}", app.UpdateAuthor)
	mux.Post("/authors", app.InsertAuthor)
	mux.Delete("/authors/{id}", app.DeleteAuthor)
	mux.Put("/books/{id_book}/authors/{id_author}", app.UpdateAuthorAndBook)
	return mux
}
