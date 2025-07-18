// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
	"time"

	"github.com/hyperledger/fabric-x-orderer/common/types"
	"github.com/hyperledger/fabric-x-orderer/node/assembler"
)

type FakePrefetchIndexerFactory struct {
	CreateStub        func([]types.ShardID, []types.PartyID, types.Logger, time.Duration, int, int, assembler.TimerFactory, assembler.BatchCacheFactory, assembler.PartitionPrefetchIndexerFactory) assembler.PrefetchIndexer
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []types.ShardID
		arg2 []types.PartyID
		arg3 types.Logger
		arg4 time.Duration
		arg5 int
		arg6 int
		arg7 assembler.TimerFactory
		arg8 assembler.BatchCacheFactory
		arg9 assembler.PartitionPrefetchIndexerFactory
	}
	createReturns struct {
		result1 assembler.PrefetchIndexer
	}
	createReturnsOnCall map[int]struct {
		result1 assembler.PrefetchIndexer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePrefetchIndexerFactory) Create(arg1 []types.ShardID, arg2 []types.PartyID, arg3 types.Logger, arg4 time.Duration, arg5 int, arg6 int, arg7 assembler.TimerFactory, arg8 assembler.BatchCacheFactory, arg9 assembler.PartitionPrefetchIndexerFactory) assembler.PrefetchIndexer {
	var arg1Copy []types.ShardID
	if arg1 != nil {
		arg1Copy = make([]types.ShardID, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []types.PartyID
	if arg2 != nil {
		arg2Copy = make([]types.PartyID, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []types.ShardID
		arg2 []types.PartyID
		arg3 types.Logger
		arg4 time.Duration
		arg5 int
		arg6 int
		arg7 assembler.TimerFactory
		arg8 assembler.BatchCacheFactory
		arg9 assembler.PartitionPrefetchIndexerFactory
	}{arg1Copy, arg2Copy, arg3, arg4, arg5, arg6, arg7, arg8, arg9})
	fake.recordInvocation("Create", []interface{}{arg1Copy, arg2Copy, arg3, arg4, arg5, arg6, arg7, arg8, arg9})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakePrefetchIndexerFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakePrefetchIndexerFactory) CreateCalls(stub func([]types.ShardID, []types.PartyID, types.Logger, time.Duration, int, int, assembler.TimerFactory, assembler.BatchCacheFactory, assembler.PartitionPrefetchIndexerFactory) assembler.PrefetchIndexer) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakePrefetchIndexerFactory) CreateArgsForCall(i int) ([]types.ShardID, []types.PartyID, types.Logger, time.Duration, int, int, assembler.TimerFactory, assembler.BatchCacheFactory, assembler.PartitionPrefetchIndexerFactory) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9
}

func (fake *FakePrefetchIndexerFactory) CreateReturns(result1 assembler.PrefetchIndexer) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 assembler.PrefetchIndexer
	}{result1}
}

func (fake *FakePrefetchIndexerFactory) CreateReturnsOnCall(i int, result1 assembler.PrefetchIndexer) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 assembler.PrefetchIndexer
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 assembler.PrefetchIndexer
	}{result1}
}

func (fake *FakePrefetchIndexerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePrefetchIndexerFactory) recordInvocation(key string, args []interface{}) {
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

var _ assembler.PrefetchIndexerFactory = new(FakePrefetchIndexerFactory)
