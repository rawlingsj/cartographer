// Code generated by counterfeiter. DO NOT EDIT.
package deliverablefakes

import (
	"context"
	"github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	"sync"

	"github.com/vmware-tanzu/cartographer/pkg/realizer/deliverable"
	"github.com/vmware-tanzu/cartographer/pkg/templates"
)

type FakeResourceRealizer struct {
	DoStub        func(context.Context, *v1alpha1.ClusterDeliveryResource, string, deliverable.Outputs) (*templates.Output, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 context.Context
		arg2 *v1alpha1.ClusterDeliveryResource
		arg3 string
		arg4 deliverable.Outputs
	}
	doReturns struct {
		result1 *templates.Output
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 *templates.Output
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceRealizer) Do(arg1 context.Context, arg2 *v1alpha1.ClusterDeliveryResource, arg3 string, arg4 deliverable.Outputs) (*templates.Output, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 context.Context
		arg2 *v1alpha1.ClusterDeliveryResource
		arg3 string
		arg4 deliverable.Outputs
	}{arg1, arg2, arg3, arg4})
	stub := fake.DoStub
	fakeReturns := fake.doReturns
	fake.recordInvocation("Do", []interface{}{arg1, arg2, arg3, arg4})
	fake.doMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceRealizer) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeResourceRealizer) DoCalls(stub func(context.Context, *v1alpha1.ClusterDeliveryResource, string, deliverable.Outputs) (*templates.Output, error)) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *FakeResourceRealizer) DoArgsForCall(i int) (context.Context, *v1alpha1.ClusterDeliveryResource, string, deliverable.Outputs) {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeResourceRealizer) DoReturns(result1 *templates.Output, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *templates.Output
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceRealizer) DoReturnsOnCall(i int, result1 *templates.Output, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 *templates.Output
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 *templates.Output
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceRealizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceRealizer) recordInvocation(key string, args []interface{}) {
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

var _ deliverable.ResourceRealizer = new(FakeResourceRealizer)
