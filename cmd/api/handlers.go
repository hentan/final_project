package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	warAndPiece := models.Book{
		ID:       1,
		Title:    "War and Piece",
		AuthorId: 1,
		Year:     1869,
		ISBN:     "978-5-389-06256-6",
	}

	books = append(books, warAndPiece)

	theIdiot := models.Book{
		ID:       2,
		Title:    "the Idiot",
		AuthorId: 2,
		Year:     1868,
		ISBN:     "916-5-219-03218-4",
	}

	books = append(books, theIdiot)

	out, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
