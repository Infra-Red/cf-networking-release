// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/policy-server/store"
)

type EgressDestinationRepo struct {
	AllStub        func(db.Transaction) ([]store.EgressDestination, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
		arg1 db.Transaction
	}
	allReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	CreateIPRangeStub        func(db.Transaction, string, string, string, string, string, int64, int64, int64, int64) error
	createIPRangeMutex       sync.RWMutex
	createIPRangeArgsForCall []struct {
		arg1  db.Transaction
		arg2  string
		arg3  string
		arg4  string
		arg5  string
		arg6  string
		arg7  int64
		arg8  int64
		arg9  int64
		arg10 int64
	}
	createIPRangeReturns struct {
		result1 error
	}
	createIPRangeReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(db.Transaction, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetByGUIDStub        func(db.Transaction, ...string) ([]store.EgressDestination, error)
	getByGUIDMutex       sync.RWMutex
	getByGUIDArgsForCall []struct {
		arg1 db.Transaction
		arg2 []string
	}
	getByGUIDReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	getByGUIDReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	GetByNameStub        func(db.Transaction, ...string) ([]store.EgressDestination, error)
	getByNameMutex       sync.RWMutex
	getByNameArgsForCall []struct {
		arg1 db.Transaction
		arg2 []string
	}
	getByNameReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	getByNameReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	UpdateIPRangeStub        func(db.Transaction, string, string, string, string, string, int64, int64, int64, int64) error
	updateIPRangeMutex       sync.RWMutex
	updateIPRangeArgsForCall []struct {
		arg1  db.Transaction
		arg2  string
		arg3  string
		arg4  string
		arg5  string
		arg6  string
		arg7  int64
		arg8  int64
		arg9  int64
		arg10 int64
	}
	updateIPRangeReturns struct {
		result1 error
	}
	updateIPRangeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressDestinationRepo) All(arg1 db.Transaction) ([]store.EgressDestination, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
		arg1 db.Transaction
	}{arg1})
	stub := fake.AllStub
	fakeReturns := fake.allReturns
	fake.recordInvocation("All", []interface{}{arg1})
	fake.allMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationRepo) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressDestinationRepo) AllCalls(stub func(db.Transaction) ([]store.EgressDestination, error)) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = stub
}

func (fake *EgressDestinationRepo) AllArgsForCall(i int) db.Transaction {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	argsForCall := fake.allArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressDestinationRepo) AllReturns(result1 []store.EgressDestination, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) AllReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) CreateIPRange(arg1 db.Transaction, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 int64, arg8 int64, arg9 int64, arg10 int64) error {
	fake.createIPRangeMutex.Lock()
	ret, specificReturn := fake.createIPRangeReturnsOnCall[len(fake.createIPRangeArgsForCall)]
	fake.createIPRangeArgsForCall = append(fake.createIPRangeArgsForCall, struct {
		arg1  db.Transaction
		arg2  string
		arg3  string
		arg4  string
		arg5  string
		arg6  string
		arg7  int64
		arg8  int64
		arg9  int64
		arg10 int64
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	stub := fake.CreateIPRangeStub
	fakeReturns := fake.createIPRangeReturns
	fake.recordInvocation("CreateIPRange", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.createIPRangeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressDestinationRepo) CreateIPRangeCallCount() int {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return len(fake.createIPRangeArgsForCall)
}

func (fake *EgressDestinationRepo) CreateIPRangeCalls(stub func(db.Transaction, string, string, string, string, string, int64, int64, int64, int64) error) {
	fake.createIPRangeMutex.Lock()
	defer fake.createIPRangeMutex.Unlock()
	fake.CreateIPRangeStub = stub
}

func (fake *EgressDestinationRepo) CreateIPRangeArgsForCall(i int) (db.Transaction, string, string, string, string, string, int64, int64, int64, int64) {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	argsForCall := fake.createIPRangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10
}

func (fake *EgressDestinationRepo) CreateIPRangeReturns(result1 error) {
	fake.createIPRangeMutex.Lock()
	defer fake.createIPRangeMutex.Unlock()
	fake.CreateIPRangeStub = nil
	fake.createIPRangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) CreateIPRangeReturnsOnCall(i int, result1 error) {
	fake.createIPRangeMutex.Lock()
	defer fake.createIPRangeMutex.Unlock()
	fake.CreateIPRangeStub = nil
	if fake.createIPRangeReturnsOnCall == nil {
		fake.createIPRangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createIPRangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) Delete(arg1 db.Transaction, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressDestinationRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *EgressDestinationRepo) DeleteCalls(stub func(db.Transaction, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *EgressDestinationRepo) DeleteArgsForCall(i int) (db.Transaction, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressDestinationRepo) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) GetByGUID(arg1 db.Transaction, arg2 ...string) ([]store.EgressDestination, error) {
	fake.getByGUIDMutex.Lock()
	ret, specificReturn := fake.getByGUIDReturnsOnCall[len(fake.getByGUIDArgsForCall)]
	fake.getByGUIDArgsForCall = append(fake.getByGUIDArgsForCall, struct {
		arg1 db.Transaction
		arg2 []string
	}{arg1, arg2})
	stub := fake.GetByGUIDStub
	fakeReturns := fake.getByGUIDReturns
	fake.recordInvocation("GetByGUID", []interface{}{arg1, arg2})
	fake.getByGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationRepo) GetByGUIDCallCount() int {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	return len(fake.getByGUIDArgsForCall)
}

func (fake *EgressDestinationRepo) GetByGUIDCalls(stub func(db.Transaction, ...string) ([]store.EgressDestination, error)) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = stub
}

func (fake *EgressDestinationRepo) GetByGUIDArgsForCall(i int) (db.Transaction, []string) {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	argsForCall := fake.getByGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressDestinationRepo) GetByGUIDReturns(result1 []store.EgressDestination, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	fake.getByGUIDReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) GetByGUIDReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	if fake.getByGUIDReturnsOnCall == nil {
		fake.getByGUIDReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.getByGUIDReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) GetByName(arg1 db.Transaction, arg2 ...string) ([]store.EgressDestination, error) {
	fake.getByNameMutex.Lock()
	ret, specificReturn := fake.getByNameReturnsOnCall[len(fake.getByNameArgsForCall)]
	fake.getByNameArgsForCall = append(fake.getByNameArgsForCall, struct {
		arg1 db.Transaction
		arg2 []string
	}{arg1, arg2})
	stub := fake.GetByNameStub
	fakeReturns := fake.getByNameReturns
	fake.recordInvocation("GetByName", []interface{}{arg1, arg2})
	fake.getByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationRepo) GetByNameCallCount() int {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	return len(fake.getByNameArgsForCall)
}

func (fake *EgressDestinationRepo) GetByNameCalls(stub func(db.Transaction, ...string) ([]store.EgressDestination, error)) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = stub
}

func (fake *EgressDestinationRepo) GetByNameArgsForCall(i int) (db.Transaction, []string) {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	argsForCall := fake.getByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressDestinationRepo) GetByNameReturns(result1 []store.EgressDestination, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	fake.getByNameReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) GetByNameReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	if fake.getByNameReturnsOnCall == nil {
		fake.getByNameReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.getByNameReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationRepo) UpdateIPRange(arg1 db.Transaction, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 int64, arg8 int64, arg9 int64, arg10 int64) error {
	fake.updateIPRangeMutex.Lock()
	ret, specificReturn := fake.updateIPRangeReturnsOnCall[len(fake.updateIPRangeArgsForCall)]
	fake.updateIPRangeArgsForCall = append(fake.updateIPRangeArgsForCall, struct {
		arg1  db.Transaction
		arg2  string
		arg3  string
		arg4  string
		arg5  string
		arg6  string
		arg7  int64
		arg8  int64
		arg9  int64
		arg10 int64
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	stub := fake.UpdateIPRangeStub
	fakeReturns := fake.updateIPRangeReturns
	fake.recordInvocation("UpdateIPRange", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.updateIPRangeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressDestinationRepo) UpdateIPRangeCallCount() int {
	fake.updateIPRangeMutex.RLock()
	defer fake.updateIPRangeMutex.RUnlock()
	return len(fake.updateIPRangeArgsForCall)
}

func (fake *EgressDestinationRepo) UpdateIPRangeCalls(stub func(db.Transaction, string, string, string, string, string, int64, int64, int64, int64) error) {
	fake.updateIPRangeMutex.Lock()
	defer fake.updateIPRangeMutex.Unlock()
	fake.UpdateIPRangeStub = stub
}

func (fake *EgressDestinationRepo) UpdateIPRangeArgsForCall(i int) (db.Transaction, string, string, string, string, string, int64, int64, int64, int64) {
	fake.updateIPRangeMutex.RLock()
	defer fake.updateIPRangeMutex.RUnlock()
	argsForCall := fake.updateIPRangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10
}

func (fake *EgressDestinationRepo) UpdateIPRangeReturns(result1 error) {
	fake.updateIPRangeMutex.Lock()
	defer fake.updateIPRangeMutex.Unlock()
	fake.UpdateIPRangeStub = nil
	fake.updateIPRangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) UpdateIPRangeReturnsOnCall(i int, result1 error) {
	fake.updateIPRangeMutex.Lock()
	defer fake.updateIPRangeMutex.Unlock()
	fake.UpdateIPRangeStub = nil
	if fake.updateIPRangeReturnsOnCall == nil {
		fake.updateIPRangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateIPRangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressDestinationRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	fake.updateIPRangeMutex.RLock()
	defer fake.updateIPRangeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressDestinationRepo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
