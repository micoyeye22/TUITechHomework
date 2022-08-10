// Code generated by mockery v2.14.0. DO NOT EDIT.

package musement

import mock "github.com/stretchr/testify/mock"

// MockMusementProviderClientConfig is an autogenerated mock type for the MusementProviderClientConfig type
type MockMusementProviderClientConfig struct {
	mock.Mock
}

// BaseURL provides a mock function with given fields:
func (_m *MockMusementProviderClientConfig) BaseURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockMusementProviderClientConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMusementProviderClientConfig creates a new instance of MockMusementProviderClientConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMusementProviderClientConfig(t mockConstructorTestingTNewMockMusementProviderClientConfig) *MockMusementProviderClientConfig {
	mock := &MockMusementProviderClientConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}