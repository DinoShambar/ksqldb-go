// Code generated by mockery v2.9.4. DO NOT EDIT.

package ksqldb

import (
	mock "github.com/stretchr/testify/mock"
	ksqldb "github.com/DinoShambar/ksqldb-go"

	net "github.com/DinoShambar/ksqldb-go/net"
)

// NewClientWithOptionsFactory is an autogenerated mock type for the NewClientWithOptionsFactory type
type NewClientWithOptionsFactory struct {
	mock.Mock
}

// NewClientWithOptions provides a mock function with given fields: options
func (_m *NewClientWithOptionsFactory) NewClientWithOptions(options net.Options) (*ksqldb.Ksqldb, error) {
	ret := _m.Called(options)

	var r0 *ksqldb.Ksqldb
	if rf, ok := ret.Get(0).(func(net.Options) *ksqldb.Ksqldb); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ksqldb.Ksqldb)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(net.Options) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
