package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hentan/final_project/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Приложение запущено",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.AllBooks()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	book, err := app.DB.OneBook(bookID)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, book)
}

func (app *application) InsertBook(w http.ResponseWriter, r *http.Request) {
	var bookWithID models.BookID

	err := app.readJSON(w, r, &bookWithID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var book models.Book
	book.ID = bookWithID.ID
	book.Title = bookWithID.Title
	book.Author = strconv.Itoa(bookWithID.AuthorID)
	book.Year = bookWithID.Year
	book.ISBN = bookWithID.ISBN

	newID, err := app.DB.InsertBook(book)
	if err != nil {
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга с id %d успешно добавлена", newID),
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}

func (app *application) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var payload models.BookID

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book, err := app.DB.OneBook(payload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book.Title = payload.Title
	book.Author = strconv.Itoa(payload.AuthorID)
	book.Year = payload.Year
	book.ISBN = payload.ISBN

	err = app.DB.UpdateBook(*book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: "Книга успешно обновлена",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}

func (app *application) AllAuthors(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.AllAuthors()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *application) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	author, err := app.DB.OneAuthor(authorID)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, author)
}
