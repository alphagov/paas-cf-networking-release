// This file was generated by counterfeiter
package fakes

import (
	"garden-external-networker/port_allocator"
	"sync"
)

type Tracker struct {
	AcquireOneStub        func(pool *port_allocator.Pool) (int, error)
	acquireOneMutex       sync.RWMutex
	acquireOneArgsForCall []struct {
		pool *port_allocator.Pool
	}
	acquireOneReturns struct {
		result1 int
		result2 error
	}
	ReleaseManyStub        func(pool *port_allocator.Pool, toRelease []int) error
	releaseManyMutex       sync.RWMutex
	releaseManyArgsForCall []struct {
		pool      *port_allocator.Pool
		toRelease []int
	}
	releaseManyReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Tracker) AcquireOne(pool *port_allocator.Pool) (int, error) {
	fake.acquireOneMutex.Lock()
	fake.acquireOneArgsForCall = append(fake.acquireOneArgsForCall, struct {
		pool *port_allocator.Pool
	}{pool})
	fake.recordInvocation("AcquireOne", []interface{}{pool})
	fake.acquireOneMutex.Unlock()
	if fake.AcquireOneStub != nil {
		return fake.AcquireOneStub(pool)
	} else {
		return fake.acquireOneReturns.result1, fake.acquireOneReturns.result2
	}
}

func (fake *Tracker) AcquireOneCallCount() int {
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	return len(fake.acquireOneArgsForCall)
}

func (fake *Tracker) AcquireOneArgsForCall(i int) *port_allocator.Pool {
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	return fake.acquireOneArgsForCall[i].pool
}

func (fake *Tracker) AcquireOneReturns(result1 int, result2 error) {
	fake.AcquireOneStub = nil
	fake.acquireOneReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Tracker) ReleaseMany(pool *port_allocator.Pool, toRelease []int) error {
	var toReleaseCopy []int
	if toRelease != nil {
		toReleaseCopy = make([]int, len(toRelease))
		copy(toReleaseCopy, toRelease)
	}
	fake.releaseManyMutex.Lock()
	fake.releaseManyArgsForCall = append(fake.releaseManyArgsForCall, struct {
		pool      *port_allocator.Pool
		toRelease []int
	}{pool, toReleaseCopy})
	fake.recordInvocation("ReleaseMany", []interface{}{pool, toReleaseCopy})
	fake.releaseManyMutex.Unlock()
	if fake.ReleaseManyStub != nil {
		return fake.ReleaseManyStub(pool, toRelease)
	} else {
		return fake.releaseManyReturns.result1
	}
}

func (fake *Tracker) ReleaseManyCallCount() int {
	fake.releaseManyMutex.RLock()
	defer fake.releaseManyMutex.RUnlock()
	return len(fake.releaseManyArgsForCall)
}

func (fake *Tracker) ReleaseManyArgsForCall(i int) (*port_allocator.Pool, []int) {
	fake.releaseManyMutex.RLock()
	defer fake.releaseManyMutex.RUnlock()
	return fake.releaseManyArgsForCall[i].pool, fake.releaseManyArgsForCall[i].toRelease
}

func (fake *Tracker) ReleaseManyReturns(result1 error) {
	fake.ReleaseManyStub = nil
	fake.releaseManyReturns = struct {
		result1 error
	}{result1}
}

func (fake *Tracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	fake.releaseManyMutex.RLock()
	defer fake.releaseManyMutex.RUnlock()
	return fake.invocations
}

func (fake *Tracker) recordInvocation(key string, args []interface{}) {
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
