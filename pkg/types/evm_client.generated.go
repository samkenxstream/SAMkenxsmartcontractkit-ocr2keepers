// Code generated by mockery v2.12.1. DO NOT EDIT.

package types

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ethereum/go-ethereum/rpc"

	testing "testing"
)

// MockEVMClient is an autogenerated mock type for the EVMClient type
type MockEVMClient struct {
	mock.Mock
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *MockEVMClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallContract provides a mock function with given fields: ctx, call, blockNumber
func (_m *MockEVMClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, call, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, call, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, call, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CodeAt provides a mock function with given fields: ctx, contract, blockNumber
func (_m *MockEVMClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, contract, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, contract, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, contract, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadTicker provides a mock function with given fields:
func (_m *MockEVMClient) HeadTicker() chan BlockKey {
	ret := _m.Called()

	var r0 chan BlockKey
	if rf, ok := ret.Get(0).(func() chan BlockKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan BlockKey)
		}
	}

	return r0
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *MockEVMClient) HeaderByNumber(ctx context.Context, number *big.Int) (*coretypes.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *coretypes.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *coretypes.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEVMClient creates a new instance of MockEVMClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEVMClient(t testing.TB) *MockEVMClient {
	mock := &MockEVMClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
