package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/models"
	"github.com/hentan/final_project/internal/repository"
)

type Application struct {
	Domain     string
	DB         repository.DatabaseRepo
	ServerConf config.Config
}

type Handler interface {
	Start(h http.Handler) error
	Home(w http.ResponseWriter, r *http.Request)
	AllBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	InsertBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	AllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	InsertAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthorAndBook(w http.ResponseWriter, r *http.Request)
}

func (app *Application) Start(h http.Handler) error {
	err := http.ListenAndServe(app.ServerConf.AppPort, h)
	if err != nil {
		return err
	}
	log.Println("успешный старт на порту %s", app.ServerConf.AppPort)
	return nil
}

func New(db repository.DatabaseRepo, cfg config.Config) Handler {
	return &Application{
		DB:         db,
		ServerConf: cfg,
	}
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
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

func (app *Application) AllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.DB.AllBooks()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book, err := app.DB.OneBook(bookID)

	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, book)
}

func (app *Application) InsertBook(w http.ResponseWriter, r *http.Request) {
	var bookWithID models.Book

	err := app.readJSON(w, r, &bookWithID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	newID, err := app.DB.InsertBook(bookWithID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга с id %d успешно добавлена", newID),
	}

	app.writeJSON(w, http.StatusCreated, resp)

}

func (app *Application) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	var payload models.Book

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book, err := app.DB.OneBook(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	book.Title = payload.Title
	book.AuthorID = payload.AuthorID
	book.Year = payload.Year
	book.ISBN = payload.ISBN

	err = app.DB.UpdateBook(*book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга c id %d успешно обновлена", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.DB.OneBook(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	err = app.DB.DeleteBook(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга c id %d успешно удалена", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

// с этой строки и ниже действия с авторами

func (app *Application) AllAuthors(w http.ResponseWriter, r *http.Request) {
	books, err := app.DB.AllAuthors()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	author, err := app.DB.OneAuthor(authorID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, author)
}

func (app *Application) InsertAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := app.readJSON(w, r, &author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	newID, err := app.DB.InsertAuthor(author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно добавлен", newID),
	}

	app.writeJSON(w, http.StatusCreated, resp)

}

func (app *Application) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload models.Author

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	author, err := app.DB.OneAuthor(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	author.NameAuthor = payload.NameAuthor
	author.SurnameAuthor = payload.SurnameAuthor
	author.Biography = payload.Biography
	author.Birthday = payload.Birthday

	err = app.DB.UpdateAuthor(*author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно обновлен", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.DB.OneAuthor(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	err = app.DB.DeleteAuthor(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно удален", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) UpdateAuthorAndBook(w http.ResponseWriter, r *http.Request) {
	id_book := chi.URLParam(r, "id_book")
	ID_book, err := strconv.Atoi(id_book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	id_author := chi.URLParam(r, "id_author")
	ID_author, err := strconv.Atoi(id_author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload models.AuthorAndBook

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	author, err := app.DB.OneAuthor(ID_author)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	book, err := app.DB.OneBook(ID_book)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	author.NameAuthor = payload.Author.NameAuthor
	author.SurnameAuthor = payload.Author.SurnameAuthor
	author.Biography = payload.Author.Biography
	author.Birthday = payload.Author.Birthday
	book.Title = payload.Book.Title
	book.AuthorID = ID_author
	book.Year = payload.Book.Year
	book.ISBN = payload.Book.ISBN

	err = app.DB.UpdateAuthorAndBook(*author, *book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: "Автор и книга успешно обновлены",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
