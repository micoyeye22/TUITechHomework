// Code generated by mockery v2.14.0. DO NOT EDIT.

package usecase

import (
	contracts "musement/src/internal/core/contracts"

	mock "github.com/stretchr/testify/mock"
)

// MockMusementProvider is an autogenerated mock type for the MusementProvider type
type MockMusementProvider struct {
	mock.Mock
}

// GetCities provides a mock function with given fields:
func (_m *MockMusementProvider) GetCities() ([]contracts.City, error) {
	ret := _m.Called()

	var r0 []contracts.City
	if rf, ok := ret.Get(0).(func() []contracts.City); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contracts.City)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockMusementProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMusementProvider creates a new instance of MockMusementProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMusementProvider(t mockConstructorTestingTNewMockMusementProvider) *MockMusementProvider {
	mock := &MockMusementProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
