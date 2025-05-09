// Code generated by mockery v2.52.3. DO NOT EDIT.

package tmux

import (
	model "github.com/joshmedeski/sesh/v2/model"
	mock "github.com/stretchr/testify/mock"
)

// MockTmux is an autogenerated mock type for the Tmux type
type MockTmux struct {
	mock.Mock
}

type MockTmux_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTmux) EXPECT() *MockTmux_Expecter {
	return &MockTmux_Expecter{mock: &_m.Mock}
}

// AttachSession provides a mock function with given fields: targetSession
func (_m *MockTmux) AttachSession(targetSession string) (string, error) {
	ret := _m.Called(targetSession)

	if len(ret) == 0 {
		panic("no return value specified for AttachSession")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(targetSession)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(targetSession)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(targetSession)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_AttachSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AttachSession'
type MockTmux_AttachSession_Call struct {
	*mock.Call
}

// AttachSession is a helper method to define mock.On call
//   - targetSession string
func (_e *MockTmux_Expecter) AttachSession(targetSession interface{}) *MockTmux_AttachSession_Call {
	return &MockTmux_AttachSession_Call{Call: _e.mock.On("AttachSession", targetSession)}
}

func (_c *MockTmux_AttachSession_Call) Run(run func(targetSession string)) *MockTmux_AttachSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTmux_AttachSession_Call) Return(_a0 string, _a1 error) *MockTmux_AttachSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_AttachSession_Call) RunAndReturn(run func(string) (string, error)) *MockTmux_AttachSession_Call {
	_c.Call.Return(run)
	return _c
}

// CapturePane provides a mock function with given fields: targetSession
func (_m *MockTmux) CapturePane(targetSession string) (string, error) {
	ret := _m.Called(targetSession)

	if len(ret) == 0 {
		panic("no return value specified for CapturePane")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(targetSession)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(targetSession)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(targetSession)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_CapturePane_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CapturePane'
type MockTmux_CapturePane_Call struct {
	*mock.Call
}

// CapturePane is a helper method to define mock.On call
//   - targetSession string
func (_e *MockTmux_Expecter) CapturePane(targetSession interface{}) *MockTmux_CapturePane_Call {
	return &MockTmux_CapturePane_Call{Call: _e.mock.On("CapturePane", targetSession)}
}

func (_c *MockTmux_CapturePane_Call) Run(run func(targetSession string)) *MockTmux_CapturePane_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTmux_CapturePane_Call) Return(_a0 string, _a1 error) *MockTmux_CapturePane_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_CapturePane_Call) RunAndReturn(run func(string) (string, error)) *MockTmux_CapturePane_Call {
	_c.Call.Return(run)
	return _c
}

// IsAttached provides a mock function with no fields
func (_m *MockTmux) IsAttached() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAttached")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockTmux_IsAttached_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAttached'
type MockTmux_IsAttached_Call struct {
	*mock.Call
}

// IsAttached is a helper method to define mock.On call
func (_e *MockTmux_Expecter) IsAttached() *MockTmux_IsAttached_Call {
	return &MockTmux_IsAttached_Call{Call: _e.mock.On("IsAttached")}
}

func (_c *MockTmux_IsAttached_Call) Run(run func()) *MockTmux_IsAttached_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTmux_IsAttached_Call) Return(_a0 bool) *MockTmux_IsAttached_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTmux_IsAttached_Call) RunAndReturn(run func() bool) *MockTmux_IsAttached_Call {
	_c.Call.Return(run)
	return _c
}

// ListSessions provides a mock function with no fields
func (_m *MockTmux) ListSessions() ([]*model.TmuxSession, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListSessions")
	}

	var r0 []*model.TmuxSession
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.TmuxSession, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.TmuxSession); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TmuxSession)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_ListSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSessions'
type MockTmux_ListSessions_Call struct {
	*mock.Call
}

// ListSessions is a helper method to define mock.On call
func (_e *MockTmux_Expecter) ListSessions() *MockTmux_ListSessions_Call {
	return &MockTmux_ListSessions_Call{Call: _e.mock.On("ListSessions")}
}

func (_c *MockTmux_ListSessions_Call) Run(run func()) *MockTmux_ListSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTmux_ListSessions_Call) Return(_a0 []*model.TmuxSession, _a1 error) *MockTmux_ListSessions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_ListSessions_Call) RunAndReturn(run func() ([]*model.TmuxSession, error)) *MockTmux_ListSessions_Call {
	_c.Call.Return(run)
	return _c
}

// NewSession provides a mock function with given fields: sessionName, startDir
func (_m *MockTmux) NewSession(sessionName string, startDir string) (string, error) {
	ret := _m.Called(sessionName, startDir)

	if len(ret) == 0 {
		panic("no return value specified for NewSession")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(sessionName, startDir)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(sessionName, startDir)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(sessionName, startDir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_NewSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewSession'
type MockTmux_NewSession_Call struct {
	*mock.Call
}

// NewSession is a helper method to define mock.On call
//   - sessionName string
//   - startDir string
func (_e *MockTmux_Expecter) NewSession(sessionName interface{}, startDir interface{}) *MockTmux_NewSession_Call {
	return &MockTmux_NewSession_Call{Call: _e.mock.On("NewSession", sessionName, startDir)}
}

func (_c *MockTmux_NewSession_Call) Run(run func(sessionName string, startDir string)) *MockTmux_NewSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTmux_NewSession_Call) Return(_a0 string, _a1 error) *MockTmux_NewSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_NewSession_Call) RunAndReturn(run func(string, string) (string, error)) *MockTmux_NewSession_Call {
	_c.Call.Return(run)
	return _c
}

// NewWindow provides a mock function with given fields: startDir, name
func (_m *MockTmux) NewWindow(startDir string, name string) (string, error) {
	ret := _m.Called(startDir, name)

	if len(ret) == 0 {
		panic("no return value specified for NewWindow")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(startDir, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(startDir, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(startDir, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_NewWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewWindow'
type MockTmux_NewWindow_Call struct {
	*mock.Call
}

// NewWindow is a helper method to define mock.On call
//   - startDir string
//   - name string
func (_e *MockTmux_Expecter) NewWindow(startDir interface{}, name interface{}) *MockTmux_NewWindow_Call {
	return &MockTmux_NewWindow_Call{Call: _e.mock.On("NewWindow", startDir, name)}
}

func (_c *MockTmux_NewWindow_Call) Run(run func(startDir string, name string)) *MockTmux_NewWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTmux_NewWindow_Call) Return(_a0 string, _a1 error) *MockTmux_NewWindow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_NewWindow_Call) RunAndReturn(run func(string, string) (string, error)) *MockTmux_NewWindow_Call {
	_c.Call.Return(run)
	return _c
}

// NextWindow provides a mock function with no fields
func (_m *MockTmux) NextWindow() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NextWindow")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_NextWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NextWindow'
type MockTmux_NextWindow_Call struct {
	*mock.Call
}

// NextWindow is a helper method to define mock.On call
func (_e *MockTmux_Expecter) NextWindow() *MockTmux_NextWindow_Call {
	return &MockTmux_NextWindow_Call{Call: _e.mock.On("NextWindow")}
}

func (_c *MockTmux_NextWindow_Call) Run(run func()) *MockTmux_NextWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTmux_NextWindow_Call) Return(_a0 string, _a1 error) *MockTmux_NextWindow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_NextWindow_Call) RunAndReturn(run func() (string, error)) *MockTmux_NextWindow_Call {
	_c.Call.Return(run)
	return _c
}

// SendKeys provides a mock function with given fields: name, command
func (_m *MockTmux) SendKeys(name string, command string) (string, error) {
	ret := _m.Called(name, command)

	if len(ret) == 0 {
		panic("no return value specified for SendKeys")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(name, command)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(name, command)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_SendKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendKeys'
type MockTmux_SendKeys_Call struct {
	*mock.Call
}

// SendKeys is a helper method to define mock.On call
//   - name string
//   - command string
func (_e *MockTmux_Expecter) SendKeys(name interface{}, command interface{}) *MockTmux_SendKeys_Call {
	return &MockTmux_SendKeys_Call{Call: _e.mock.On("SendKeys", name, command)}
}

func (_c *MockTmux_SendKeys_Call) Run(run func(name string, command string)) *MockTmux_SendKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTmux_SendKeys_Call) Return(_a0 string, _a1 error) *MockTmux_SendKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_SendKeys_Call) RunAndReturn(run func(string, string) (string, error)) *MockTmux_SendKeys_Call {
	_c.Call.Return(run)
	return _c
}

// SwitchClient provides a mock function with given fields: targetSession
func (_m *MockTmux) SwitchClient(targetSession string) (string, error) {
	ret := _m.Called(targetSession)

	if len(ret) == 0 {
		panic("no return value specified for SwitchClient")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(targetSession)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(targetSession)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(targetSession)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_SwitchClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SwitchClient'
type MockTmux_SwitchClient_Call struct {
	*mock.Call
}

// SwitchClient is a helper method to define mock.On call
//   - targetSession string
func (_e *MockTmux_Expecter) SwitchClient(targetSession interface{}) *MockTmux_SwitchClient_Call {
	return &MockTmux_SwitchClient_Call{Call: _e.mock.On("SwitchClient", targetSession)}
}

func (_c *MockTmux_SwitchClient_Call) Run(run func(targetSession string)) *MockTmux_SwitchClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTmux_SwitchClient_Call) Return(_a0 string, _a1 error) *MockTmux_SwitchClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_SwitchClient_Call) RunAndReturn(run func(string) (string, error)) *MockTmux_SwitchClient_Call {
	_c.Call.Return(run)
	return _c
}

// SwitchOrAttach provides a mock function with given fields: name, opts
func (_m *MockTmux) SwitchOrAttach(name string, opts model.ConnectOpts) (string, error) {
	ret := _m.Called(name, opts)

	if len(ret) == 0 {
		panic("no return value specified for SwitchOrAttach")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, model.ConnectOpts) (string, error)); ok {
		return rf(name, opts)
	}
	if rf, ok := ret.Get(0).(func(string, model.ConnectOpts) string); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, model.ConnectOpts) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTmux_SwitchOrAttach_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SwitchOrAttach'
type MockTmux_SwitchOrAttach_Call struct {
	*mock.Call
}

// SwitchOrAttach is a helper method to define mock.On call
//   - name string
//   - opts model.ConnectOpts
func (_e *MockTmux_Expecter) SwitchOrAttach(name interface{}, opts interface{}) *MockTmux_SwitchOrAttach_Call {
	return &MockTmux_SwitchOrAttach_Call{Call: _e.mock.On("SwitchOrAttach", name, opts)}
}

func (_c *MockTmux_SwitchOrAttach_Call) Run(run func(name string, opts model.ConnectOpts)) *MockTmux_SwitchOrAttach_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.ConnectOpts))
	})
	return _c
}

func (_c *MockTmux_SwitchOrAttach_Call) Return(_a0 string, _a1 error) *MockTmux_SwitchOrAttach_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTmux_SwitchOrAttach_Call) RunAndReturn(run func(string, model.ConnectOpts) (string, error)) *MockTmux_SwitchOrAttach_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTmux creates a new instance of MockTmux. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTmux(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTmux {
	mock := &MockTmux{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
