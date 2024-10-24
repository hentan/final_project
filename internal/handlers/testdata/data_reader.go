package testdata

import (
	"bytes"
	"embed"
	"encoding/json"

	"github.com/hentan/final_project/internal/models"
	"github.com/stretchr/testify/require"
)

//go:embed books/*.json
var fs embed.FS

type suite interface {
	Require() *require.Assertions
}

func ReadBook(s suite, path string) []*models.Book {
	data := ReadFile(s, path)
	var books []models.Book

	err := json.Unmarshal(data, &books)
	if err != nil {
		var singleBook models.Book
		err = json.Unmarshal(data, &singleBook)
		s.Require().NoError(err)

		books = append(books, singleBook)
	} else {
		s.Require().NoError(err)
	}

	bookPtrs := make([]*models.Book, len(books))
	for i := range books {
		bookPtrs[i] = &books[i]
	}
	return bookPtrs
}

func ReadFile(s suite, path string) []byte {
	content, err := fs.ReadFile(path)
	s.Require().NoError(err)
	return bytes.TrimSpace(content)
}
