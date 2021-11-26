// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	internal "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
	mock "github.com/stretchr/testify/mock"
)

// BrokerClient is an autogenerated mock type for the BrokerClient type
type BrokerClient struct {
	mock.Mock
}

// Deprovision provides a mock function with given fields: instance
func (_m *BrokerClient) Deprovision(instance internal.Instance) (string, error) {
	ret := _m.Called(instance)

	var r0 string
	if rf, ok := ret.Get(0).(func(internal.Instance) string); ok {
		r0 = rf(instance)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Instance) error); ok {
		r1 = rf(instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
