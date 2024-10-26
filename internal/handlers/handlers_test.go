package handlers_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/handlers/testdata"
	mocks "github.com/hentan/final_project/internal/mocks/repository"
	"github.com/hentan/final_project/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type appSuite struct {
	suite.Suite
	application *handlers.Application
	repo        *mocks.DatabaseRepo
}

func (s *appSuite) SetupTest() {
	// Инициализация моков и приложения перед каждым тестом
	s.repo = new(mocks.DatabaseRepo)
	s.application = &handlers.Application{DB: s.repo}
}

func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	app := &appSuite{}
	app.application.Home(rec, req)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	byte_mes := rec.Body.Bytes()
	map_for_json := make(map[string]string)

	err := json.Unmarshal(byte_mes, &map_for_json)
	require.NoError(t, err)
	expected_status := "active"
	expected_message := "Приложение запущено"

	assert.Equal(t, expected_status, map_for_json["status"])
	assert.Equal(t, expected_message, map_for_json["message"])
}

func (s *appSuite) TestAllBooks() {

	expectedBooks := testdata.ReadBook(&s.Suite, "books/read_all_books_success.json")
	s.repo.On("AllBooks").Return(expectedBooks, nil).Once()
	req, err := http.NewRequest("GET", "/books", nil)
	require.NoError(s.T(), err)

	rr := httptest.NewRecorder()
	s.application.AllBooks(rr, req)
	assert.Equal(s.T(), http.StatusOK, rr.Code)

	var actualBooks []*models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &actualBooks)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), expectedBooks, actualBooks)
}

func (s *appSuite) TestGetBook() {

	expectedBooks := testdata.ReadBook(&s.Suite, "books/read_all_books_success.json")
	book := *expectedBooks[0]

	s.repo.On("OneBook", 2).Return(&book, nil).Once()
	req, err := http.NewRequest("GET", "/book/2", nil)
	require.NoError(s.T(), err)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.application.GetBook(rr, req)
	assert.Equal(s.T(), http.StatusOK, rr.Code)

	expected := `{"id":2,"title":"the Idiot","author":"Fyodor Dostoevsky","year":1868,"isbn":"978-1533695840"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestGetBookNotFound() {
	bookID := 999
	s.repo.On("OneBook", bookID).Return(nil, sql.ErrNoRows).Once()

	req, err := http.NewRequest("GET", "/book/999", nil)
	require.NoError(s.T(), err)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.application.GetBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestInsertBookSuccess() {
	expectedBooks := testdata.ReadBook(&s.Suite, "books/insert_one_book_success.json")
	book := expectedBooks[0]
	reqBody, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	s.repo.On("InsertBook", *book).Return(1, nil)
	s.application.InsertBook(rr, req)
	assert.Equal(s.T(), http.StatusCreated, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), false, resp.Error)
	assert.Equal(s.T(), "Книга с id 1 успешно добавлена", resp.Message)
}

func (s *appSuite) TestInsertBookError() {
	incJSON, _ := json.Marshal('{')
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(incJSON))

	rr := httptest.NewRecorder()
	s.application.InsertBook(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Equal(s.T(), "json: cannot unmarshal number into Go value of type models.Book", resp.Message)
}

func (s *appSuite) TestUpdateBook_Success() {
	existingBooks := testdata.ReadBook(&s.Suite, "books/update_book_success.json")
	book := existingBooks[0]

	updatedBook := models.Book{
		Title:    "Updated Book",
		AuthorID: book.AuthorID,
		Year:     book.Year,
		ISBN:     "9876543210",
	}

	reqBody, _ := json.Marshal(updatedBook)
	req := httptest.NewRequest(http.MethodPut, "/book/2", bytes.NewBuffer(reqBody))

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()

	s.repo.On("OneBook", 2).Return(book, nil)
	s.repo.On("UpdateBook", updatedBook).Return(nil)

	s.application.UpdateBook(rr, req)

	s.Equal(http.StatusOK, rr.Code, "Ожидался код ответа 200 OK")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.False(resp.Error, "Ожидалось, что ошибка будет равна false")
	s.Equal(fmt.Sprintf("Книга c id %d успешно обновлена", 2), resp.Message)
}

func (s *appSuite) TestUpdateBook_Error_Not_Found() {
	existingBooks := testdata.ReadBook(&s.Suite, "books/update_book_success.json")
	book := existingBooks[0]

	updatedBook := models.Book{
		Title:    "Updated Book",
		AuthorID: book.AuthorID,
		Year:     book.Year,
		ISBN:     "9876543210",
	}
	reqBody, _ := json.Marshal(updatedBook)
	bookID := 999

	req := httptest.NewRequest(http.MethodPut, "/book/999", bytes.NewBuffer(reqBody))
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneBook", bookID).Return(nil, sql.ErrNoRows).Once()
	rr := httptest.NewRecorder()
	s.application.UpdateBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestUpdateBook_Error_Incorrect_JSON() {
	incJSON, _ := json.Marshal('{')
	req := httptest.NewRequest(http.MethodPost, "/book/2", bytes.NewBuffer(incJSON))

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.application.UpdateBook(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Equal(s.T(), "json: cannot unmarshal number into Go value of type models.Book", resp.Message)
}

func (s *appSuite) TestDeleteBook_Success() {
	expectedBooks := testdata.ReadBook(&s.Suite, "books/read_all_books_success.json")
	book := *expectedBooks[0]

	req := httptest.NewRequest(http.MethodDelete, "/book/2", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneBook", 2).Return(&book, nil)
	s.repo.On("DeleteBook", 2).Return(nil)

	rr := httptest.NewRecorder()

	s.application.DeleteBook(rr, req)
	s.Equal(http.StatusOK, rr.Code, "Ожидался код ответа 200 OK")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.False(resp.Error, "Ожидалось, что ошибка будет равна false")
	s.Equal(fmt.Sprintf("Книга c id %d успешно удалена", 2), resp.Message)
}

func (s *appSuite) TestDeleteBook_Error_Not_Found() {

	bookID := 999
	req := httptest.NewRequest(http.MethodDelete, "/book/999", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneBook", bookID).Return(nil, sql.ErrNoRows).Once()
	rr := httptest.NewRecorder()
	s.application.DeleteBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestDeleteBook_Server_Error() {
	expectedBooks := testdata.ReadBook(&s.Suite, "books/read_all_books_success.json")
	book := *expectedBooks[0]

	req := httptest.NewRequest(http.MethodDelete, "/book/2", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.repo.On("OneBook", 2).Return(&book, nil)
	s.repo.On("DeleteBook", 2).Return(errors.New("ошибка при удалении"))

	s.application.DeleteBook(rr, req)

	s.Equal(http.StatusInternalServerError, rr.Code, "Ожидался код ответа 500 Internal Server Error")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.True(resp.Error, "Ожидалось, что ошибка будет равна true")
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(appSuite))
}
