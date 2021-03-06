// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	eth "github.com/smartcontractkit/chainlink/core/services/eth"
	mock "github.com/stretchr/testify/mock"
)

// LogBroadcaster is an autogenerated mock type for the LogBroadcaster type
type LogBroadcaster struct {
	mock.Mock
}

// AddDependents provides a mock function with given fields: n
func (_m *LogBroadcaster) AddDependents(n int) {
	_m.Called(n)
}

// AwaitDependents provides a mock function with given fields:
func (_m *LogBroadcaster) AwaitDependents() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// DependentReady provides a mock function with given fields:
func (_m *LogBroadcaster) DependentReady() {
	_m.Called()
}

// Register provides a mock function with given fields: address, listener
func (_m *LogBroadcaster) Register(address common.Address, listener eth.LogListener) bool {
	ret := _m.Called(address, listener)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address, eth.LogListener) bool); ok {
		r0 = rf(address, listener)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *LogBroadcaster) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *LogBroadcaster) Stop() {
	_m.Called()
}

// Unregister provides a mock function with given fields: address, listener
func (_m *LogBroadcaster) Unregister(address common.Address, listener eth.LogListener) {
	_m.Called(address, listener)
}
