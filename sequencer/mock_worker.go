// Code generated by mockery v2.39.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	types "github.com/ethereum/go-ethereum/core/types"
)

// WorkerMock is an autogenerated mock type for the workerInterface type
type WorkerMock struct {
	mock.Mock
}

// AddForcedTx provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) AddForcedTx(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// AddTxTracker provides a mock function with given fields: ctx, txTracker
func (_m *WorkerMock) AddTxTracker(ctx context.Context, txTracker *TxTracker) (*TxTracker, error) {
	ret := _m.Called(ctx, txTracker)

	if len(ret) == 0 {
		panic("no return value specified for AddTxTracker")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) (*TxTracker, error)); ok {
		return rf(ctx, txTracker)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) *TxTracker); ok {
		r0 = rf(ctx, txTracker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *TxTracker) error); ok {
		r1 = rf(ctx, txTracker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForcedTx provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) DeleteForcedTx(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// DeleteTx provides a mock function with given fields: txHash, from
func (_m *WorkerMock) DeleteTx(txHash common.Hash, from common.Address) {
	_m.Called(txHash, from)
}

// DeleteTxPendingToStore provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) DeleteTxPendingToStore(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// GetBestFittingTx provides a mock function with given fields: remainingResources, highReservedCounters
func (_m *WorkerMock) GetBestFittingTx(remainingResources state.BatchResources, highReservedCounters state.ZKCounters) (*TxTracker, error) {
	ret := _m.Called(remainingResources, highReservedCounters)

	if len(ret) == 0 {
		panic("no return value specified for GetBestFittingTx")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(state.BatchResources, state.ZKCounters) (*TxTracker, error)); ok {
		return rf(remainingResources, highReservedCounters)
	}
	if rf, ok := ret.Get(0).(func(state.BatchResources, state.ZKCounters) *TxTracker); ok {
		r0 = rf(remainingResources, highReservedCounters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(state.BatchResources, state.ZKCounters) error); ok {
		r1 = rf(remainingResources, highReservedCounters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveTxPendingToStore provides a mock function with given fields: txHash, addr
func (_m *WorkerMock) MoveTxPendingToStore(txHash common.Hash, addr common.Address) {
	_m.Called(txHash, addr)
}

// MoveTxToNotReady provides a mock function with given fields: txHash, from, actualNonce, actualBalance
func (_m *WorkerMock) MoveTxToNotReady(txHash common.Hash, from common.Address, actualNonce *uint64, actualBalance *big.Int) []*TxTracker {
	ret := _m.Called(txHash, from, actualNonce, actualBalance)

	if len(ret) == 0 {
		panic("no return value specified for MoveTxToNotReady")
	}

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Hash, common.Address, *uint64, *big.Int) []*TxTracker); ok {
		r0 = rf(txHash, from, actualNonce, actualBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// NewTxTracker provides a mock function with given fields: tx, usedZKcounters, reservedZKCouners, ip
func (_m *WorkerMock) NewTxTracker(tx types.Transaction, usedZKcounters state.ZKCounters, reservedZKCouners state.ZKCounters, ip string) (*TxTracker, error) {
	ret := _m.Called(tx, usedZKcounters, reservedZKCouners, ip)

	if len(ret) == 0 {
		panic("no return value specified for NewTxTracker")
	}

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Transaction, state.ZKCounters, state.ZKCounters, string) (*TxTracker, error)); ok {
		return rf(tx, usedZKcounters, reservedZKCouners, ip)
	}
	if rf, ok := ret.Get(0).(func(types.Transaction, state.ZKCounters, state.ZKCounters, string) *TxTracker); ok {
		r0 = rf(tx, usedZKcounters, reservedZKCouners, ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Transaction, state.ZKCounters, state.ZKCounters, string) error); ok {
		r1 = rf(tx, usedZKcounters, reservedZKCouners, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreTxsPendingToStore provides a mock function with given fields: ctx
func (_m *WorkerMock) RestoreTxsPendingToStore(ctx context.Context) ([]*TxTracker, []*TxTracker) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for RestoreTxsPendingToStore")
	}

	var r0 []*TxTracker
	var r1 []*TxTracker
	if rf, ok := ret.Get(0).(func(context.Context) ([]*TxTracker, []*TxTracker)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*TxTracker); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) []*TxTracker); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*TxTracker)
		}
	}

	return r0, r1
}

// UpdateAfterSingleSuccessfulTxExecution provides a mock function with given fields: from, touchedAddresses
func (_m *WorkerMock) UpdateAfterSingleSuccessfulTxExecution(from common.Address, touchedAddresses map[common.Address]*state.InfoReadWrite) []*TxTracker {
	ret := _m.Called(from, touchedAddresses)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAfterSingleSuccessfulTxExecution")
	}

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Address, map[common.Address]*state.InfoReadWrite) []*TxTracker); ok {
		r0 = rf(from, touchedAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// UpdateTxZKCounters provides a mock function with given fields: txHash, from, usedZKCounters, reservedZKCounters
func (_m *WorkerMock) UpdateTxZKCounters(txHash common.Hash, from common.Address, usedZKCounters state.ZKCounters, reservedZKCounters state.ZKCounters) {
	_m.Called(txHash, from, usedZKCounters, reservedZKCounters)
}

// NewWorkerMock creates a new instance of WorkerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkerMock {
	mock := &WorkerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
