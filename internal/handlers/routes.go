package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/hentan/final_project/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(h Handler) http.Handler {
	//mux роутер
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Get("/", h.Home)
	mux.Get("/books", h.AllBooks)
	mux.Get("/books/{id}", h.GetBook)
	mux.Put("/books/{id}", h.UpdateBook)
	mux.Post("/books", h.InsertBook)
	mux.Delete("/books/{id}", h.DeleteBook)
	mux.Get("/authors", h.AllAuthors)
	mux.Get("/authors/{id}", h.GetAuthor)
	mux.Put("/authors/{id}", h.UpdateAuthor)
	mux.Post("/authors", h.InsertAuthor)
	mux.Delete("/authors/{id}", h.DeleteAuthor)
	mux.Put("/books/{id_book}/authors/{id_author}", h.UpdateAuthorAndBook)
	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	return mux
}
