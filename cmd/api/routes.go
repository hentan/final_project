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
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)
	mux.Get("/books", app.AllBooks)
	mux.Get("/books/{id}", app.GetBook)
	mux.Get("/authors", app.AllAuthors)
	mux.Get("/authors/{id}", app.GetAuthor)
	return mux
}
