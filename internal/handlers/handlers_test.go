package handlers_test

import (
	"context"
	"encoding/json"
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

func (s *appSuite) TestGetBooks() {

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

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(appSuite))
}
