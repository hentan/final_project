package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/kafka"
	"github.com/hentan/final_project/internal/logger"
	"github.com/hentan/final_project/internal/models"
	"github.com/hentan/final_project/internal/repository"
)

type Application struct {
	Domain      string
	DB          repository.DatabaseRepo
	ServerConf  config.Config
	KafkaClient kafka.KafkaProducer
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
	newLogger := logger.GetLogger()
	err := http.ListenAndServe(app.ServerConf.AppPort, h)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("успешный старт на порту %s", app.ServerConf.AppPort)
	newLogger.Info(msg)
	return nil
}

func New(db repository.DatabaseRepo, cfg config.Config, kafkaProducer kafka.KafkaProducer) Handler {
	return &Application{
		DB:          db,
		ServerConf:  cfg,
		KafkaClient: kafkaProducer,
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
		wrapError := fmt.Errorf("handlers/handlers.go AllBooks error can't read from DB, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go GetBook error ID must be INT, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	book, err := app.DB.OneBook(bookID)

	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go GetBook error can't find book!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, book)
}

func (app *Application) InsertBook(w http.ResponseWriter, r *http.Request) {
	var bookWithID models.Book

	err := app.readJSON(w, r, &bookWithID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go InsertBook error can't parse JSON!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	newID, err := app.DB.InsertBook(bookWithID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go InsertBook error can't insert book in database!, %w", err)
		app.errorJSON(w, wrapError)
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
		wrapError := fmt.Errorf("handlers/handlers.go UpdateBook error ID must be INT!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	var payload models.Book

	err = app.readJSON(w, r, &payload)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateBook error can't parse JSON!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	book, err := app.DB.OneBook(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateBook error can't find book!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	book.Title = payload.Title
	book.AuthorID = payload.AuthorID
	book.Year = payload.Year
	book.ISBN = payload.ISBN

	err = app.DB.UpdateBook(*book)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateBook error problems with update book!, %w", err)
		app.errorJSON(w, wrapError, 500)
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
		wrapError := fmt.Errorf("handlers/handlers.go DeleteBook ID must be INT!, %w", err)
		app.errorJSON(w, wrapError, 400)
		return
	}

	_, err = app.DB.OneBook(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go DeleteBook error can't find book!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	err = app.DB.DeleteBook(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go DeleteBook error problems with delete from database!, %w", err)
		app.errorJSON(w, wrapError, 500)
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
		wrapError := fmt.Errorf("handlers/handlers.go AllAuthors error reading authors from DB, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go GetAuthor error ID must be INT!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	author, err := app.DB.OneAuthor(authorID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go GetAuthor error can't find author!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, author)
}

func (app *Application) InsertAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := app.readJSON(w, r, &author)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go InsertAuthor error can't parse JSON!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	newID, err := app.DB.InsertAuthor(author)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go InsertAuthor error can't insert author in database!, %w", err)
		app.errorJSON(w, wrapError)
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
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthor error ID must be INT!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	var payload models.Author

	err = app.readJSON(w, r, &payload)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthor error can't parse JSON!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	author, err := app.DB.OneAuthor(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthor error can't find author!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	author.NameAuthor = payload.NameAuthor
	author.SurnameAuthor = payload.SurnameAuthor
	author.Biography = payload.Biography
	author.Birthday = payload.Birthday

	err = app.DB.UpdateAuthor(*author)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthor error problems with update author!, %w", err)
		app.errorJSON(w, wrapError, 500)
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
		wrapError := fmt.Errorf("handlers/handlers.go DeleteAuthor ID must be INT!, %w", err)
		app.errorJSON(w, wrapError, 400)
		return
	}

	_, err = app.DB.OneAuthor(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go DeleteAuthor error can't find author!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	err = app.DB.DeleteAuthor(ID)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go DeleteAuthor error problems with delete from database!, %w", err)
		app.errorJSON(w, wrapError, 500)
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
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error ID must be INT!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	id_author := chi.URLParam(r, "id_author")
	ID_author, err := strconv.Atoi(id_author)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error ID must be INT!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	var payload models.AuthorAndBook

	err = app.readJSON(w, r, &payload)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error can't parse JSON!, %w", err)
		app.errorJSON(w, wrapError)
		return
	}

	author, err := app.DB.OneAuthor(ID_author)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error can't find author!, %w", err)
		app.errorJSON(w, wrapError, 404)
		return
	}

	book, err := app.DB.OneBook(ID_book)
	if err != nil {
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error can't find book!, %w", err)
		app.errorJSON(w, wrapError, 404)
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
		wrapError := fmt.Errorf("handlers/handlers.go UpdateAuthorAndBook error problems with update author or book!, %w", err)
		app.errorJSON(w, wrapError, 500)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: "Автор и книга успешно обновлены",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
