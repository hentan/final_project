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
	mocksKafka "github.com/hentan/final_project/internal/mocks/kafka"
	mocksRedis "github.com/hentan/final_project/internal/mocks/redis"
	mocks "github.com/hentan/final_project/internal/mocks/repository"
	"github.com/hentan/final_project/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type appSuite struct {
	suite.Suite
	application *handlers.Application
	repo        *mocks.DatabaseRepo
	kafka       *mocksKafka.KafkaProducer
	redis       *mocksRedis.RedisClient
}

func (s *appSuite) SetupTest() {
	// Инициализация моков и приложения перед каждым тестом
	s.repo = new(mocks.DatabaseRepo)
	s.kafka = new(mocksKafka.KafkaProducer)
	s.redis = new(mocksRedis.RedisClient)
	s.application = &handlers.Application{
		DB:          s.repo,
		KafkaClient: s.kafka,
		RedisClient: s.redis,
	}
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

	s.redis.On("GetFromCache", mock.Anything, 2).Return(nil, nil).Once()
	s.repo.On("OneBook", 2).Return(&book, nil).Once()
	s.redis.On("SetToCache", mock.Anything, 2, &book, mock.Anything).Return(nil).Once()
	fmt.Println(s.redis)
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
	s.redis.On("GetFromCache", mock.Anything, 999).Return(nil, nil).Once()
	s.repo.On("OneBook", bookID).Return(nil, sql.ErrNoRows).Once()
	req, err := http.NewRequest("GET", "/book/999", nil)
	require.NoError(s.T(), err)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()
	s.application.GetBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go GetBook error can't find book!, sql: no rows in result set"}`
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
	s.kafka.On("SendMessage", mock.Anything).Return(nil)
	s.application.InsertBook(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Contains(s.T(), resp.Message, "handlers/handlers.go InsertBook error can't parse JSON!")
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
	s.kafka.On("SendMessage", mock.Anything).Return(nil)
	rr := httptest.NewRecorder()
	s.application.UpdateBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go UpdateBook error can't find book!, sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestUpdateBook_Error_Incorrect_JSON() {
	incJSON, _ := json.Marshal('{')
	req := httptest.NewRequest(http.MethodPost, "/book/2", bytes.NewBuffer(incJSON))

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()
	s.application.UpdateBook(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Contains(s.T(), resp.Message, "handlers/handlers.go UpdateBook error can't parse JSON!")
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
	s.kafka.On("SendMessage", mock.Anything).Return(nil)
	rr := httptest.NewRecorder()
	s.application.DeleteBook(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go DeleteBook error can't find book!, sql: no rows in result set"}`
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
	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	s.application.DeleteBook(rr, req)

	s.Equal(http.StatusInternalServerError, rr.Code, "Ожидался код ответа 500 Internal Server Error")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.True(resp.Error, "Ожидалось, что ошибка будет равна true")
}

func (s *appSuite) TestAllAuthors() {

	expectedAuthors := testdata.ReadAuthor(&s.Suite, "author/read_all_author_success.json")
	s.repo.On("AllAuthors").Return(expectedAuthors, nil).Once()
	req, err := http.NewRequest("GET", "/Authors", nil)
	require.NoError(s.T(), err)

	rr := httptest.NewRecorder()
	s.application.AllAuthors(rr, req)
	assert.Equal(s.T(), http.StatusOK, rr.Code)

	var actualAuthors []*models.Author
	err = json.Unmarshal(rr.Body.Bytes(), &actualAuthors)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), expectedAuthors, actualAuthors)
}

func (s *appSuite) TestGetAuthor() {

	expectedAuthors := testdata.ReadAuthor(&s.Suite, "author/read_all_author_success.json")
	Author := *expectedAuthors[0]

	s.repo.On("OneAuthor", 2).Return(&Author, nil).Once()
	req, err := http.NewRequest("GET", "/Author/2", nil)
	require.NoError(s.T(), err)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.application.GetAuthor(rr, req)
	assert.Equal(s.T(), http.StatusOK, rr.Code)

	expected := `{"id":2,"name_author":"Fyodor","surname_author":"Dostoevsky","biography":"Fyodor Mikhailovich Dostoevsky[a] (UK: /ˌdɒstɔɪˈɛfski/,[1] US: /ˌdɒstəˈjɛfski, ˌdʌs-/;[2] Russian: Фёдор Михайлович Достоевский[b], romanized: Fyodor Mikhaylovich Dostoyevskiy, IPA: [ˈfʲɵdər mʲɪˈxajləvʲɪdʑ dəstɐˈjefskʲɪj] ⓘ; 11 November 1821 – 9 February 1881[3][c]), sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces","birthday":"1821-11-11T00:00:00Z"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestGetAuthorNotFound() {
	AuthorID := 999
	s.repo.On("OneAuthor", AuthorID).Return(nil, sql.ErrNoRows).Once()

	req, err := http.NewRequest("GET", "/Author/999", nil)
	require.NoError(s.T(), err)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()
	s.application.GetAuthor(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go GetAuthor error can't find author!, sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestInsertAuthorSuccess() {
	expectedAuthors := testdata.ReadAuthor(&s.Suite, "author/insert_one_author_success.json")
	Author := expectedAuthors[0]
	reqBody, _ := json.Marshal(Author)
	req := httptest.NewRequest(http.MethodPost, "/Authors", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	s.repo.On("InsertAuthor", *Author).Return(1, nil)
	s.application.InsertAuthor(rr, req)
	assert.Equal(s.T(), http.StatusCreated, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), false, resp.Error)
	assert.Equal(s.T(), "Автор с id 1 успешно добавлен", resp.Message)
}

func (s *appSuite) TestInsertAuthorError() {
	incJSON, _ := json.Marshal('{')
	req := httptest.NewRequest(http.MethodPost, "/Authors", bytes.NewBuffer(incJSON))

	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()
	s.application.InsertAuthor(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Contains(s.T(), resp.Message, "handlers/handlers.go InsertAuthor error can't parse JSON!")
}

func (s *appSuite) TestUpdateAuthor_Success() {
	existingAuthors := testdata.ReadAuthor(&s.Suite, "author/update_author_success.json")
	Author := existingAuthors[0]

	updatedAuthor := models.Author{
		NameAuthor:    "Антон",
		SurnameAuthor: "Чехов",
		Biography:     "Русский писатель, драматург и врач.",
		Birthday:      "1860-01-29",
	}

	reqBody, _ := json.Marshal(updatedAuthor)
	req := httptest.NewRequest(http.MethodPut, "/Author/2", bytes.NewBuffer(reqBody))

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()

	s.repo.On("OneAuthor", 2).Return(Author, nil)
	s.repo.On("UpdateAuthor", updatedAuthor).Return(nil)
	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	s.application.UpdateAuthor(rr, req)

	s.Equal(http.StatusOK, rr.Code, "Ожидался код ответа 200 OK")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.False(resp.Error, "Ожидалось, что ошибка будет равна false")
	s.Equal(fmt.Sprintf("Автор с id %d успешно обновлен", 2), resp.Message)
}

func (s *appSuite) TestUpdateAuthor_Error_Not_Found() {
	existingAuthors := testdata.ReadAuthor(&s.Suite, "author/update_author_success.json")
	Author := existingAuthors[0]

	updatedAuthor := models.Author{
		NameAuthor:    "Антон",
		SurnameAuthor: "Чехов",
		Biography:     "Русский писатель, драматург и врач.",
		Birthday:      "1860-01-29",
		ID:            Author.ID,
	}
	reqBody, _ := json.Marshal(updatedAuthor)
	AuthorID := 999

	req := httptest.NewRequest(http.MethodPut, "/Author/999", bytes.NewBuffer(reqBody))
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneAuthor", AuthorID).Return(nil, sql.ErrNoRows).Once()
	s.kafka.On("SendMessage", mock.Anything).Return(nil)
	rr := httptest.NewRecorder()
	s.application.UpdateAuthor(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go UpdateAuthor error can't find author!, sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestUpdateAuthor_Error_Incorrect_JSON() {
	incJSON, _ := json.Marshal('{')
	req := httptest.NewRequest(http.MethodPost, "/Author/2", bytes.NewBuffer(incJSON))

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()
	s.application.UpdateAuthor(rr, req)
	assert.Equal(s.T(), http.StatusBadRequest, rr.Code)

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), true, resp.Error)
	assert.Contains(s.T(), resp.Message, "handlers/handlers.go UpdateAuthor error can't parse JSON!")
}

func (s *appSuite) TestDeleteAuthor_Success() {
	expectedAuthors := testdata.ReadAuthor(&s.Suite, "author/read_all_author_success.json")
	Author := *expectedAuthors[0]

	req := httptest.NewRequest(http.MethodDelete, "/Author/2", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneAuthor", 2).Return(&Author, nil)
	s.repo.On("DeleteAuthor", 2).Return(nil)

	rr := httptest.NewRecorder()

	s.application.DeleteAuthor(rr, req)
	s.Equal(http.StatusOK, rr.Code, "Ожидался код ответа 200 OK")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.False(resp.Error, "Ожидалось, что ошибка будет равна false")
	s.Equal(fmt.Sprintf("Автор с id %d успешно удален", 2), resp.Message)
}

func (s *appSuite) TestDeleteAuthor_Error_Not_Found() {

	AuthorID := 999
	req := httptest.NewRequest(http.MethodDelete, "/Author/999", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	s.repo.On("OneAuthor", AuthorID).Return(nil, sql.ErrNoRows).Once()
	s.kafka.On("SendMessage", mock.Anything).Return(nil)
	rr := httptest.NewRecorder()
	s.application.DeleteAuthor(rr, req)
	assert.Equal(s.T(), http.StatusNotFound, rr.Code)

	expected := `{"error":true, "message":"handlers/handlers.go DeleteAuthor error can't find author!, sql: no rows in result set"}`
	s.Require().JSONEq(expected, rr.Body.String())
}

func (s *appSuite) TestDeleteAuthor_Server_Error() {
	expectedAuthors := testdata.ReadAuthor(&s.Suite, "author/read_all_author_success.json")
	Author := *expectedAuthors[0]

	req := httptest.NewRequest(http.MethodDelete, "/Author/2", nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	rr := httptest.NewRecorder()
	s.repo.On("OneAuthor", 2).Return(&Author, nil)
	s.repo.On("DeleteAuthor", 2).Return(errors.New("ошибка при удалении"))
	s.kafka.On("SendMessage", mock.Anything).Return(nil)

	s.application.DeleteAuthor(rr, req)

	s.Equal(http.StatusInternalServerError, rr.Code, "Ожидался код ответа 500 Internal Server Error")

	var resp handlers.JSONResponce
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	s.NoError(err, "Не удалось разобрать JSON ответ")
	s.True(resp.Error, "Ожидалось, что ошибка будет равна true")
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(appSuite))
}
