/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"
)

type Git struct {
	SyncStub        func(context.Context) error
	syncMutex       sync.RWMutex
	syncArgsForCall []struct {
		arg1 context.Context
	}
	syncReturns struct {
		result1 error
	}
	syncReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Git) Sync(arg1 context.Context) error {
	fake.syncMutex.Lock()
	ret, specificReturn := fake.syncReturnsOnCall[len(fake.syncArgsForCall)]
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Sync", []interface{}{arg1})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		return fake.SyncStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.syncReturns.result1
}

func (fake *Git) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *Git) SyncArgsForCall(i int) context.Context {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return fake.syncArgsForCall[i].arg1
}

func (fake *Git) SyncReturns(result1 error) {
	fake.SyncStub = nil
	fake.syncReturns = struct {
		result1 error
	}{result1}
}

func (fake *Git) SyncReturnsOnCall(i int, result1 error) {
	fake.SyncStub = nil
	if fake.syncReturnsOnCall == nil {
		fake.syncReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.syncReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Git) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Git) recordInvocation(key string, args []interface{}) {
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
