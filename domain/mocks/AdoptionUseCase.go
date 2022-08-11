// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "petadopter/domain"

	mock "github.com/stretchr/testify/mock"
)

// AdoptionUseCase is an autogenerated mock type for the AdoptionUseCase type
type AdoptionUseCase struct {
	mock.Mock
}

// AddAdoption provides a mock function with given fields: IDUser, newAdops
func (_m *AdoptionUseCase) AddAdoption(IDUser int, newAdops domain.Adoption) (domain.Adoption, error) {
	ret := _m.Called(IDUser, newAdops)

	var r0 domain.Adoption
	if rf, ok := ret.Get(0).(func(int, domain.Adoption) domain.Adoption); ok {
		r0 = rf(IDUser, newAdops)
	} else {
		r0 = ret.Get(0).(domain.Adoption)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Adoption) error); ok {
		r1 = rf(IDUser, newAdops)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelAdoption provides a mock function with given fields: IDAdoption
func (_m *AdoptionUseCase) DelAdoption(IDAdoption int) (bool, error) {
	ret := _m.Called(IDAdoption)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDAdoption)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IDAdoption)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAP provides a mock function with given fields: userID
func (_m *AdoptionUseCase) GetAllAP(userID int) ([]domain.AdoptionPet, error) {
	ret := _m.Called(userID)

	var r0 []domain.AdoptionPet
	if rf, ok := ret.Get(0).(func(int) []domain.AdoptionPet); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AdoptionPet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpecificAdoption provides a mock function with given fields: AdoptionID
func (_m *AdoptionUseCase) GetSpecificAdoption(AdoptionID int) ([]domain.AdoptionPet, error) {
	ret := _m.Called(AdoptionID)

	var r0 []domain.AdoptionPet
	if rf, ok := ret.Get(0).(func(int) []domain.AdoptionPet); ok {
		r0 = rf(AdoptionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AdoptionPet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(AdoptionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetmyAdoption provides a mock function with given fields: userID
func (_m *AdoptionUseCase) GetmyAdoption(userID int) ([]domain.AdoptionPet, error) {
	ret := _m.Called(userID)

	var r0 []domain.AdoptionPet
	if rf, ok := ret.Get(0).(func(int) []domain.AdoptionPet); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AdoptionPet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpAdoption provides a mock function with given fields: IDAdoption, updateData
func (_m *AdoptionUseCase) UpAdoption(IDAdoption int, updateData domain.Adoption) (domain.Adoption, error) {
	ret := _m.Called(IDAdoption, updateData)

	var r0 domain.Adoption
	if rf, ok := ret.Get(0).(func(int, domain.Adoption) domain.Adoption); ok {
		r0 = rf(IDAdoption, updateData)
	} else {
		r0 = ret.Get(0).(domain.Adoption)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Adoption) error); ok {
		r1 = rf(IDAdoption, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdoptionUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdoptionUseCase creates a new instance of AdoptionUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdoptionUseCase(t mockConstructorTestingTNewAdoptionUseCase) *AdoptionUseCase {
	mock := &AdoptionUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}