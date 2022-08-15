// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MeetingHandler is an autogenerated mock type for the MeetingHandler type
type MeetingHandler struct {
	mock.Mock
}

// DeleteDataMeeting provides a mock function with given fields:
func (_m *MeetingHandler) DeleteDataMeeting() echo.HandlerFunc {
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

// GetAdopt provides a mock function with given fields:
func (_m *MeetingHandler) GetAdopt() echo.HandlerFunc {
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

// InsertMeeting provides a mock function with given fields:
func (_m *MeetingHandler) InsertMeeting() echo.HandlerFunc {
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

// UpdateDataMeeting provides a mock function with given fields:
func (_m *MeetingHandler) UpdateDataMeeting() echo.HandlerFunc {
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

type mockConstructorTestingTNewMeetingHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMeetingHandler creates a new instance of MeetingHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMeetingHandler(t mockConstructorTestingTNewMeetingHandler) *MeetingHandler {
	mock := &MeetingHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
