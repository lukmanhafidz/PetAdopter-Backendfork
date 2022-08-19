// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SpeciesHandler is an autogenerated mock type for the SpeciesHandler type
type SpeciesHandler struct {
	mock.Mock
}

// AddSpecies provides a mock function with given fields:
func (_m *SpeciesHandler) AddSpecies() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// DeleteDataSpecies provides a mock function with given fields:
func (_m *SpeciesHandler) DeleteDataSpecies() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetSpecies provides a mock function with given fields:
func (_m *SpeciesHandler) GetSpecies() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// UpdateDataSpecies provides a mock function with given fields:
func (_m *SpeciesHandler) UpdateDataSpecies() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewSpeciesHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpeciesHandler creates a new instance of SpeciesHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpeciesHandler(t mockConstructorTestingTNewSpeciesHandler) *SpeciesHandler {
	mock := &SpeciesHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
