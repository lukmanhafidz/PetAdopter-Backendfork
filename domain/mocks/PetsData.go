// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "petadopter/domain"

	mock "github.com/stretchr/testify/mock"
)

// PetsData is an autogenerated mock type for the PetsData type
type PetsData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: IDPets
func (_m *PetsData) Delete(IDPets int) bool {
	ret := _m.Called(IDPets)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDPets)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *PetsData) GetAll() []domain.Pets {
	ret := _m.Called()

	var r0 []domain.Pets
	if rf, ok := ret.Get(0).(func() []domain.Pets); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pets)
		}
	}

	return r0
}

// GetPetUser provides a mock function with given fields:
func (_m *PetsData) GetPetUser() domain.PetUser {
	ret := _m.Called()

	var r0 domain.PetUser
	if rf, ok := ret.Get(0).(func() domain.PetUser); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.PetUser)
	}

	return r0
}

// GetPetsID provides a mock function with given fields: PetsID
func (_m *PetsData) GetPetsID(PetsID int) []domain.Pets {
	ret := _m.Called(PetsID)

	var r0 []domain.Pets
	if rf, ok := ret.Get(0).(func(int) []domain.Pets); ok {
		r0 = rf(PetsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pets)
		}
	}

	return r0
}

// GetPetsbyuser provides a mock function with given fields: userID
func (_m *PetsData) GetPetsbyuser(userID int) []domain.Pets {
	ret := _m.Called(userID)

	var r0 []domain.Pets
	if rf, ok := ret.Get(0).(func(int) []domain.Pets); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pets)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: insertPets
func (_m *PetsData) Insert(insertPets domain.Pets) domain.Pets {
	ret := _m.Called(insertPets)

	var r0 domain.Pets
	if rf, ok := ret.Get(0).(func(domain.Pets) domain.Pets); ok {
		r0 = rf(insertPets)
	} else {
		r0 = ret.Get(0).(domain.Pets)
	}

	return r0
}

// Update provides a mock function with given fields: IDPets, updatedPets
func (_m *PetsData) Update(IDPets int, updatedPets domain.Pets) domain.Pets {
	ret := _m.Called(IDPets, updatedPets)

	var r0 domain.Pets
	if rf, ok := ret.Get(0).(func(int, domain.Pets) domain.Pets); ok {
		r0 = rf(IDPets, updatedPets)
	} else {
		r0 = ret.Get(0).(domain.Pets)
	}

	return r0
}

type mockConstructorTestingTNewPetsData interface {
	mock.TestingT
	Cleanup(func())
}

// NewPetsData creates a new instance of PetsData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPetsData(t mockConstructorTestingTNewPetsData) *PetsData {
	mock := &PetsData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}