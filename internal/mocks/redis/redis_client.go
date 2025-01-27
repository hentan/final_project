// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/hentan/final_project/internal/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// RedisClient is an autogenerated mock type for the RedisClient type
type RedisClient struct {
	mock.Mock
}

type RedisClient_Expecter struct {
	mock *mock.Mock
}

func (_m *RedisClient) EXPECT() *RedisClient_Expecter {
	return &RedisClient_Expecter{mock: &_m.Mock}
}

// DeleteFromCaсhe provides a mock function with given fields: ctx, bookID
func (_m *RedisClient) DeleteFromCaсhe(ctx context.Context, bookID int) error {
	ret := _m.Called(ctx, bookID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFromCaсhe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisClient_DeleteFromCaсhe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFromCaсhe'
type RedisClient_DeleteFromCaсhe_Call struct {
	*mock.Call
}

// DeleteFromCaсhe is a helper method to define mock.On call
//   - ctx context.Context
//   - bookID int
func (_e *RedisClient_Expecter) DeleteFromCaсhe(ctx interface{}, bookID interface{}) *RedisClient_DeleteFromCaсhe_Call {
	return &RedisClient_DeleteFromCaсhe_Call{Call: _e.mock.On("DeleteFromCaсhe", ctx, bookID)}
}

func (_c *RedisClient_DeleteFromCaсhe_Call) Run(run func(ctx context.Context, bookID int)) *RedisClient_DeleteFromCaсhe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *RedisClient_DeleteFromCaсhe_Call) Return(_a0 error) *RedisClient_DeleteFromCaсhe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RedisClient_DeleteFromCaсhe_Call) RunAndReturn(run func(context.Context, int) error) *RedisClient_DeleteFromCaсhe_Call {
	_c.Call.Return(run)
	return _c
}

// GetFromCache provides a mock function with given fields: ctx, ID, result
func (_m *RedisClient) GetFromCache(ctx context.Context, ID int, result models.Deserializable) error {
	ret := _m.Called(ctx, ID, result)

	if len(ret) == 0 {
		panic("no return value specified for GetFromCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, models.Deserializable) error); ok {
		r0 = rf(ctx, ID, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisClient_GetFromCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFromCache'
type RedisClient_GetFromCache_Call struct {
	*mock.Call
}

// GetFromCache is a helper method to define mock.On call
//   - ctx context.Context
//   - ID int
//   - result models.Deserializable
func (_e *RedisClient_Expecter) GetFromCache(ctx interface{}, ID interface{}, result interface{}) *RedisClient_GetFromCache_Call {
	return &RedisClient_GetFromCache_Call{Call: _e.mock.On("GetFromCache", ctx, ID, result)}
}

func (_c *RedisClient_GetFromCache_Call) Run(run func(ctx context.Context, ID int, result models.Deserializable)) *RedisClient_GetFromCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(models.Deserializable))
	})
	return _c
}

func (_c *RedisClient_GetFromCache_Call) Return(_a0 error) *RedisClient_GetFromCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RedisClient_GetFromCache_Call) RunAndReturn(run func(context.Context, int, models.Deserializable) error) *RedisClient_GetFromCache_Call {
	_c.Call.Return(run)
	return _c
}

// SetToCache provides a mock function with given fields: ctx, ID, bookOrAuthor, ttl
func (_m *RedisClient) SetToCache(ctx context.Context, ID int, bookOrAuthor models.Serializable, ttl time.Duration) error {
	ret := _m.Called(ctx, ID, bookOrAuthor, ttl)

	if len(ret) == 0 {
		panic("no return value specified for SetToCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, models.Serializable, time.Duration) error); ok {
		r0 = rf(ctx, ID, bookOrAuthor, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisClient_SetToCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetToCache'
type RedisClient_SetToCache_Call struct {
	*mock.Call
}

// SetToCache is a helper method to define mock.On call
//   - ctx context.Context
//   - ID int
//   - bookOrAuthor models.Serializable
//   - ttl time.Duration
func (_e *RedisClient_Expecter) SetToCache(ctx interface{}, ID interface{}, bookOrAuthor interface{}, ttl interface{}) *RedisClient_SetToCache_Call {
	return &RedisClient_SetToCache_Call{Call: _e.mock.On("SetToCache", ctx, ID, bookOrAuthor, ttl)}
}

func (_c *RedisClient_SetToCache_Call) Run(run func(ctx context.Context, ID int, bookOrAuthor models.Serializable, ttl time.Duration)) *RedisClient_SetToCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(models.Serializable), args[3].(time.Duration))
	})
	return _c
}

func (_c *RedisClient_SetToCache_Call) Return(_a0 error) *RedisClient_SetToCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RedisClient_SetToCache_Call) RunAndReturn(run func(context.Context, int, models.Serializable, time.Duration) error) *RedisClient_SetToCache_Call {
	_c.Call.Return(run)
	return _c
}

// NewRedisClient creates a new instance of RedisClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRedisClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *RedisClient {
	mock := &RedisClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}