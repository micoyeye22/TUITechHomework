// Code generated by mockery v2.14.0. DO NOT EDIT.

package client

import (
	config "musement/src/internal/infrastructure/providers/musement/config"

	mock "github.com/stretchr/testify/mock"
)

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

// MusementProviderClientConfig provides a mock function with given fields:
func (_m *MockConfig) MusementProviderClientConfig() config.MusementProviderClientConfig {
	ret := _m.Called()

	var r0 config.MusementProviderClientConfig
	if rf, ok := ret.Get(0).(func() config.MusementProviderClientConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.MusementProviderClientConfig)
	}

	return r0
}

type mockConstructorTestingTNewMockConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConfig creates a new instance of MockConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConfig(t mockConstructorTestingTNewMockConfig) *MockConfig {
	mock := &MockConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
