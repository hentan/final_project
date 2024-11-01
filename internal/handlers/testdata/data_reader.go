package testdata

import (
	"bytes"
	"embed"
	"encoding/json"

	"github.com/hentan/final_project/internal/models"
	"github.com/stretchr/testify/require"
)

//go:embed books/*.json author/*.json
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

func ReadAuthor(s suite, path string) []*models.Author {
	data := ReadFile(s, path)
	var authors []models.Author

	err := json.Unmarshal(data, &authors)
	if err != nil {
		var singleAuthor models.Author
		err = json.Unmarshal(data, &singleAuthor)
		s.Require().NoError(err)

		authors = append(authors, singleAuthor)
	} else {
		s.Require().NoError(err)
	}

	authorPtrs := make([]*models.Author, len(authors))
	for i := range authors {
		authorPtrs[i] = &authors[i]
	}
	return authorPtrs
}

func ReadFile(s suite, path string) []byte {
	content, err := fs.ReadFile(path)
	s.Require().NoError(err)
	return bytes.TrimSpace(content)
}
