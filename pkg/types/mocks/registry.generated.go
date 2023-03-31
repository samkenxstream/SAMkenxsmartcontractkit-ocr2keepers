// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/smartcontractkit/ocr2keepers/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

// CheckUpkeep provides a mock function with given fields: _a0, _a1
func (_m *Registry) CheckUpkeep(_a0 context.Context, _a1 ...types.UpkeepKey) (types.UpkeepResults, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 types.UpkeepResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...types.UpkeepKey) (types.UpkeepResults, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...types.UpkeepKey) types.UpkeepResults); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.UpkeepResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...types.UpkeepKey) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveUpkeepIDs provides a mock function with given fields: _a0
func (_m *Registry) GetActiveUpkeepIDs(_a0 context.Context) ([]types.UpkeepIdentifier, error) {
	ret := _m.Called(_a0)

	var r0 []types.UpkeepIdentifier
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]types.UpkeepIdentifier, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []types.UpkeepIdentifier); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.UpkeepIdentifier)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegistry interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegistry creates a new instance of Registry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegistry(t mockConstructorTestingTNewRegistry) *Registry {
	mock := &Registry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
