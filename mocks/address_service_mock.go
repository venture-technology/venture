// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AddressService is an autogenerated mock type for the AddressService type
type AddressService struct {
	mock.Mock
}

// GetDistance provides a mock function with given fields: origin, destination
func (_m *AddressService) GetDistance(origin string, destination string) (*float64, error) {
	ret := _m.Called(origin, destination)

	if len(ret) == 0 {
		panic("no return value specified for GetDistance")
	}

	var r0 *float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*float64, error)); ok {
		return rf(origin, destination)
	}
	if rf, ok := ret.Get(0).(func(string, string) *float64); ok {
		r0 = rf(origin, destination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*float64)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(origin, destination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAddressService creates a new instance of AddressService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAddressService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AddressService {
	mock := &AddressService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
