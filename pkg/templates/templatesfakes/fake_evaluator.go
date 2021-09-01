// Code generated by counterfeiter. DO NOT EDIT.
package templatesfakes

import (
	"sync"
)

type FakeEvaluator struct {
	EvaluateJsonPathStub        func(string, interface{}) (interface{}, error)
	evaluateJsonPathMutex       sync.RWMutex
	evaluateJsonPathArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	evaluateJsonPathReturns struct {
		result1 interface{}
		result2 error
	}
	evaluateJsonPathReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEvaluator) EvaluateJsonPath(arg1 string, arg2 interface{}) (interface{}, error) {
	fake.evaluateJsonPathMutex.Lock()
	ret, specificReturn := fake.evaluateJsonPathReturnsOnCall[len(fake.evaluateJsonPathArgsForCall)]
	fake.evaluateJsonPathArgsForCall = append(fake.evaluateJsonPathArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.EvaluateJsonPathStub
	fakeReturns := fake.evaluateJsonPathReturns
	fake.recordInvocation("EvaluateJsonPath", []interface{}{arg1, arg2})
	fake.evaluateJsonPathMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEvaluator) EvaluateJsonPathCallCount() int {
	fake.evaluateJsonPathMutex.RLock()
	defer fake.evaluateJsonPathMutex.RUnlock()
	return len(fake.evaluateJsonPathArgsForCall)
}

func (fake *FakeEvaluator) EvaluateJsonPathCalls(stub func(string, interface{}) (interface{}, error)) {
	fake.evaluateJsonPathMutex.Lock()
	defer fake.evaluateJsonPathMutex.Unlock()
	fake.EvaluateJsonPathStub = stub
}

func (fake *FakeEvaluator) EvaluateJsonPathArgsForCall(i int) (string, interface{}) {
	fake.evaluateJsonPathMutex.RLock()
	defer fake.evaluateJsonPathMutex.RUnlock()
	argsForCall := fake.evaluateJsonPathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEvaluator) EvaluateJsonPathReturns(result1 interface{}, result2 error) {
	fake.evaluateJsonPathMutex.Lock()
	defer fake.evaluateJsonPathMutex.Unlock()
	fake.EvaluateJsonPathStub = nil
	fake.evaluateJsonPathReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeEvaluator) EvaluateJsonPathReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.evaluateJsonPathMutex.Lock()
	defer fake.evaluateJsonPathMutex.Unlock()
	fake.EvaluateJsonPathStub = nil
	if fake.evaluateJsonPathReturnsOnCall == nil {
		fake.evaluateJsonPathReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.evaluateJsonPathReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeEvaluator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evaluateJsonPathMutex.RLock()
	defer fake.evaluateJsonPathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEvaluator) recordInvocation(key string, args []interface{}) {
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
