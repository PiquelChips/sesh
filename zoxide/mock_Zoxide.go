// Code generated by mockery v2.50.0. DO NOT EDIT.

package zoxide

import (
	model "github.com/joshmedeski/sesh/model"
	mock "github.com/stretchr/testify/mock"
)

// MockZoxide is an autogenerated mock type for the Zoxide type
type MockZoxide struct {
	mock.Mock
}

type MockZoxide_Expecter struct {
	mock *mock.Mock
}

func (_m *MockZoxide) EXPECT() *MockZoxide_Expecter {
	return &MockZoxide_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: path
func (_m *MockZoxide) Add(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockZoxide_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockZoxide_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - path string
func (_e *MockZoxide_Expecter) Add(path interface{}) *MockZoxide_Add_Call {
	return &MockZoxide_Add_Call{Call: _e.mock.On("Add", path)}
}

func (_c *MockZoxide_Add_Call) Run(run func(path string)) *MockZoxide_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockZoxide_Add_Call) Return(_a0 error) *MockZoxide_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockZoxide_Add_Call) RunAndReturn(run func(string) error) *MockZoxide_Add_Call {
	_c.Call.Return(run)
	return _c
}

// ListResults provides a mock function with no fields
func (_m *MockZoxide) ListResults() ([]*model.ZoxideResult, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListResults")
	}

	var r0 []*model.ZoxideResult
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.ZoxideResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.ZoxideResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ZoxideResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockZoxide_ListResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListResults'
type MockZoxide_ListResults_Call struct {
	*mock.Call
}

// ListResults is a helper method to define mock.On call
func (_e *MockZoxide_Expecter) ListResults() *MockZoxide_ListResults_Call {
	return &MockZoxide_ListResults_Call{Call: _e.mock.On("ListResults")}
}

func (_c *MockZoxide_ListResults_Call) Run(run func()) *MockZoxide_ListResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockZoxide_ListResults_Call) Return(_a0 []*model.ZoxideResult, _a1 error) *MockZoxide_ListResults_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockZoxide_ListResults_Call) RunAndReturn(run func() ([]*model.ZoxideResult, error)) *MockZoxide_ListResults_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: path
func (_m *MockZoxide) Query(path string) (*model.ZoxideResult, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *model.ZoxideResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.ZoxideResult, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *model.ZoxideResult); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ZoxideResult)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockZoxide_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockZoxide_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - path string
func (_e *MockZoxide_Expecter) Query(path interface{}) *MockZoxide_Query_Call {
	return &MockZoxide_Query_Call{Call: _e.mock.On("Query", path)}
}

func (_c *MockZoxide_Query_Call) Run(run func(path string)) *MockZoxide_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockZoxide_Query_Call) Return(_a0 *model.ZoxideResult, _a1 error) *MockZoxide_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockZoxide_Query_Call) RunAndReturn(run func(string) (*model.ZoxideResult, error)) *MockZoxide_Query_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockZoxide creates a new instance of MockZoxide. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockZoxide(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockZoxide {
	mock := &MockZoxide{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
