// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// AllAuthors provides a mock function with given fields: w, r
func (_m *Handler) AllAuthors(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_AllAuthors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllAuthors'
type Handler_AllAuthors_Call struct {
	*mock.Call
}

// AllAuthors is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) AllAuthors(w interface{}, r interface{}) *Handler_AllAuthors_Call {
	return &Handler_AllAuthors_Call{Call: _e.mock.On("AllAuthors", w, r)}
}

func (_c *Handler_AllAuthors_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_AllAuthors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_AllAuthors_Call) Return() *Handler_AllAuthors_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_AllAuthors_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_AllAuthors_Call {
	_c.Call.Return(run)
	return _c
}

// AllBooks provides a mock function with given fields: w, r
func (_m *Handler) AllBooks(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_AllBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllBooks'
type Handler_AllBooks_Call struct {
	*mock.Call
}

// AllBooks is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) AllBooks(w interface{}, r interface{}) *Handler_AllBooks_Call {
	return &Handler_AllBooks_Call{Call: _e.mock.On("AllBooks", w, r)}
}

func (_c *Handler_AllBooks_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_AllBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_AllBooks_Call) Return() *Handler_AllBooks_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_AllBooks_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_AllBooks_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAuthor provides a mock function with given fields: w, r
func (_m *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_DeleteAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAuthor'
type Handler_DeleteAuthor_Call struct {
	*mock.Call
}

// DeleteAuthor is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) DeleteAuthor(w interface{}, r interface{}) *Handler_DeleteAuthor_Call {
	return &Handler_DeleteAuthor_Call{Call: _e.mock.On("DeleteAuthor", w, r)}
}

func (_c *Handler_DeleteAuthor_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_DeleteAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_DeleteAuthor_Call) Return() *Handler_DeleteAuthor_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_DeleteAuthor_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_DeleteAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBook provides a mock function with given fields: w, r
func (_m *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_DeleteBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBook'
type Handler_DeleteBook_Call struct {
	*mock.Call
}

// DeleteBook is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) DeleteBook(w interface{}, r interface{}) *Handler_DeleteBook_Call {
	return &Handler_DeleteBook_Call{Call: _e.mock.On("DeleteBook", w, r)}
}

func (_c *Handler_DeleteBook_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_DeleteBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_DeleteBook_Call) Return() *Handler_DeleteBook_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_DeleteBook_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_DeleteBook_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthor provides a mock function with given fields: w, r
func (_m *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_GetAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthor'
type Handler_GetAuthor_Call struct {
	*mock.Call
}

// GetAuthor is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) GetAuthor(w interface{}, r interface{}) *Handler_GetAuthor_Call {
	return &Handler_GetAuthor_Call{Call: _e.mock.On("GetAuthor", w, r)}
}

func (_c *Handler_GetAuthor_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_GetAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_GetAuthor_Call) Return() *Handler_GetAuthor_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_GetAuthor_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_GetAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// GetBook provides a mock function with given fields: w, r
func (_m *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_GetBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBook'
type Handler_GetBook_Call struct {
	*mock.Call
}

// GetBook is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) GetBook(w interface{}, r interface{}) *Handler_GetBook_Call {
	return &Handler_GetBook_Call{Call: _e.mock.On("GetBook", w, r)}
}

func (_c *Handler_GetBook_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_GetBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_GetBook_Call) Return() *Handler_GetBook_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_GetBook_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_GetBook_Call {
	_c.Call.Return(run)
	return _c
}

// Home provides a mock function with given fields: w, r
func (_m *Handler) Home(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_Home_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Home'
type Handler_Home_Call struct {
	*mock.Call
}

// Home is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) Home(w interface{}, r interface{}) *Handler_Home_Call {
	return &Handler_Home_Call{Call: _e.mock.On("Home", w, r)}
}

func (_c *Handler_Home_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_Home_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_Home_Call) Return() *Handler_Home_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_Home_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_Home_Call {
	_c.Call.Return(run)
	return _c
}

// InsertAuthor provides a mock function with given fields: w, r
func (_m *Handler) InsertAuthor(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_InsertAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertAuthor'
type Handler_InsertAuthor_Call struct {
	*mock.Call
}

// InsertAuthor is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) InsertAuthor(w interface{}, r interface{}) *Handler_InsertAuthor_Call {
	return &Handler_InsertAuthor_Call{Call: _e.mock.On("InsertAuthor", w, r)}
}

func (_c *Handler_InsertAuthor_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_InsertAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_InsertAuthor_Call) Return() *Handler_InsertAuthor_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_InsertAuthor_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_InsertAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// InsertBook provides a mock function with given fields: w, r
func (_m *Handler) InsertBook(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_InsertBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertBook'
type Handler_InsertBook_Call struct {
	*mock.Call
}

// InsertBook is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) InsertBook(w interface{}, r interface{}) *Handler_InsertBook_Call {
	return &Handler_InsertBook_Call{Call: _e.mock.On("InsertBook", w, r)}
}

func (_c *Handler_InsertBook_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_InsertBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_InsertBook_Call) Return() *Handler_InsertBook_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_InsertBook_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_InsertBook_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: h
func (_m *Handler) Start(h http.Handler) error {
	ret := _m.Called(h)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.Handler) error); ok {
		r0 = rf(h)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Handler_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - h http.Handler
func (_e *Handler_Expecter) Start(h interface{}) *Handler_Start_Call {
	return &Handler_Start_Call{Call: _e.mock.On("Start", h)}
}

func (_c *Handler_Start_Call) Run(run func(h http.Handler)) *Handler_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Handler))
	})
	return _c
}

func (_c *Handler_Start_Call) Return(_a0 error) *Handler_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Start_Call) RunAndReturn(run func(http.Handler) error) *Handler_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAuthor provides a mock function with given fields: w, r
func (_m *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_UpdateAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuthor'
type Handler_UpdateAuthor_Call struct {
	*mock.Call
}

// UpdateAuthor is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) UpdateAuthor(w interface{}, r interface{}) *Handler_UpdateAuthor_Call {
	return &Handler_UpdateAuthor_Call{Call: _e.mock.On("UpdateAuthor", w, r)}
}

func (_c *Handler_UpdateAuthor_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_UpdateAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_UpdateAuthor_Call) Return() *Handler_UpdateAuthor_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_UpdateAuthor_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_UpdateAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAuthorAndBook provides a mock function with given fields: w, r
func (_m *Handler) UpdateAuthorAndBook(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_UpdateAuthorAndBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuthorAndBook'
type Handler_UpdateAuthorAndBook_Call struct {
	*mock.Call
}

// UpdateAuthorAndBook is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) UpdateAuthorAndBook(w interface{}, r interface{}) *Handler_UpdateAuthorAndBook_Call {
	return &Handler_UpdateAuthorAndBook_Call{Call: _e.mock.On("UpdateAuthorAndBook", w, r)}
}

func (_c *Handler_UpdateAuthorAndBook_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_UpdateAuthorAndBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_UpdateAuthorAndBook_Call) Return() *Handler_UpdateAuthorAndBook_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_UpdateAuthorAndBook_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_UpdateAuthorAndBook_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBook provides a mock function with given fields: w, r
func (_m *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Handler_UpdateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBook'
type Handler_UpdateBook_Call struct {
	*mock.Call
}

// UpdateBook is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) UpdateBook(w interface{}, r interface{}) *Handler_UpdateBook_Call {
	return &Handler_UpdateBook_Call{Call: _e.mock.On("UpdateBook", w, r)}
}

func (_c *Handler_UpdateBook_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_UpdateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_UpdateBook_Call) Return() *Handler_UpdateBook_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_UpdateBook_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *Handler_UpdateBook_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}