// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	require "github.com/stretchr/testify/require"
)

// suite is an autogenerated mock type for the suite type
type suite struct {
	mock.Mock
}

type suite_Expecter struct {
	mock *mock.Mock
}

func (_m *suite) EXPECT() *suite_Expecter {
	return &suite_Expecter{mock: &_m.Mock}
}

// Require provides a mock function with given fields:
func (_m *suite) Require() *require.Assertions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Require")
	}

	var r0 *require.Assertions
	if rf, ok := ret.Get(0).(func() *require.Assertions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*require.Assertions)
		}
	}

	return r0
}

// suite_Require_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Require'
type suite_Require_Call struct {
	*mock.Call
}

// Require is a helper method to define mock.On call
func (_e *suite_Expecter) Require() *suite_Require_Call {
	return &suite_Require_Call{Call: _e.mock.On("Require")}
}

func (_c *suite_Require_Call) Run(run func()) *suite_Require_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *suite_Require_Call) Return(_a0 *require.Assertions) *suite_Require_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *suite_Require_Call) RunAndReturn(run func() *require.Assertions) *suite_Require_Call {
	_c.Call.Return(run)
	return _c
}

// newSuite creates a new instance of suite. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newSuite(t interface {
	mock.TestingT
	Cleanup(func())
}) *suite {
	mock := &suite{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}