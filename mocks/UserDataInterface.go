// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	users "be15/clean/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserDataInterface is an autogenerated mock type for the UserDataInterface type
type UserDataInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserDataInterface) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *UserDataInterface) Insert(input users.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UserDataInterface) Login(email string, password string) (users.Core, string, error) {
	ret := _m.Called(email, password)

	var r0 users.Core
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (users.Core, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) users.Core); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectAll provides a mock function with given fields:
func (_m *UserDataInterface) SelectAll() ([]users.Core, error) {
	ret := _m.Called()

	var r0 []users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]users.Core, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []users.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserDataInterface creates a new instance of UserDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserDataInterface(t mockConstructorTestingTNewUserDataInterface) *UserDataInterface {
	mock := &UserDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
