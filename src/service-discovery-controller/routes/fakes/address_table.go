// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"service-discovery-controller/routes"
	"sync"
)

type AddressTable struct {
	LookupStub        func(hostname string) []string
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		hostname string
	}
	lookupReturns struct {
		result1 []string
	}
	lookupReturnsOnCall map[int]struct {
		result1 []string
	}
	GetAllAddressesStub        func() map[string][]string
	getAllAddressesMutex       sync.RWMutex
	getAllAddressesArgsForCall []struct{}
	getAllAddressesReturns     struct {
		result1 map[string][]string
	}
	getAllAddressesReturnsOnCall map[int]struct {
		result1 map[string][]string
	}
	IsWarmStub        func() bool
	isWarmMutex       sync.RWMutex
	isWarmArgsForCall []struct{}
	isWarmReturns     struct {
		result1 bool
	}
	isWarmReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AddressTable) Lookup(hostname string) []string {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		hostname string
	}{hostname})
	fake.recordInvocation("Lookup", []interface{}{hostname})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(hostname)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lookupReturns.result1
}

func (fake *AddressTable) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *AddressTable) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].hostname
}

func (fake *AddressTable) LookupReturns(result1 []string) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 []string
	}{result1}
}

func (fake *AddressTable) LookupReturnsOnCall(i int, result1 []string) {
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *AddressTable) GetAllAddresses() map[string][]string {
	fake.getAllAddressesMutex.Lock()
	ret, specificReturn := fake.getAllAddressesReturnsOnCall[len(fake.getAllAddressesArgsForCall)]
	fake.getAllAddressesArgsForCall = append(fake.getAllAddressesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllAddresses", []interface{}{})
	fake.getAllAddressesMutex.Unlock()
	if fake.GetAllAddressesStub != nil {
		return fake.GetAllAddressesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAllAddressesReturns.result1
}

func (fake *AddressTable) GetAllAddressesCallCount() int {
	fake.getAllAddressesMutex.RLock()
	defer fake.getAllAddressesMutex.RUnlock()
	return len(fake.getAllAddressesArgsForCall)
}

func (fake *AddressTable) GetAllAddressesReturns(result1 map[string][]string) {
	fake.GetAllAddressesStub = nil
	fake.getAllAddressesReturns = struct {
		result1 map[string][]string
	}{result1}
}

func (fake *AddressTable) GetAllAddressesReturnsOnCall(i int, result1 map[string][]string) {
	fake.GetAllAddressesStub = nil
	if fake.getAllAddressesReturnsOnCall == nil {
		fake.getAllAddressesReturnsOnCall = make(map[int]struct {
			result1 map[string][]string
		})
	}
	fake.getAllAddressesReturnsOnCall[i] = struct {
		result1 map[string][]string
	}{result1}
}

func (fake *AddressTable) IsWarm() bool {
	fake.isWarmMutex.Lock()
	ret, specificReturn := fake.isWarmReturnsOnCall[len(fake.isWarmArgsForCall)]
	fake.isWarmArgsForCall = append(fake.isWarmArgsForCall, struct{}{})
	fake.recordInvocation("IsWarm", []interface{}{})
	fake.isWarmMutex.Unlock()
	if fake.IsWarmStub != nil {
		return fake.IsWarmStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isWarmReturns.result1
}

func (fake *AddressTable) IsWarmCallCount() int {
	fake.isWarmMutex.RLock()
	defer fake.isWarmMutex.RUnlock()
	return len(fake.isWarmArgsForCall)
}

func (fake *AddressTable) IsWarmReturns(result1 bool) {
	fake.IsWarmStub = nil
	fake.isWarmReturns = struct {
		result1 bool
	}{result1}
}

func (fake *AddressTable) IsWarmReturnsOnCall(i int, result1 bool) {
	fake.IsWarmStub = nil
	if fake.isWarmReturnsOnCall == nil {
		fake.isWarmReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isWarmReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *AddressTable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.getAllAddressesMutex.RLock()
	defer fake.getAllAddressesMutex.RUnlock()
	fake.isWarmMutex.RLock()
	defer fake.isWarmMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AddressTable) recordInvocation(key string, args []interface{}) {
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

var _ routes.AddressTable = new(AddressTable)
