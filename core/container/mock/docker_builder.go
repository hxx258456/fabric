// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"io"
	"sync"

	"github.com/hxx258456/fabric/core/chaincode/persistence"
	"github.com/hxx258456/fabric/core/container"
)

type DockerBuilder struct {
	BuildStub        func(string, *persistence.ChaincodePackageMetadata, io.Reader) (container.Instance, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 string
		arg2 *persistence.ChaincodePackageMetadata
		arg3 io.Reader
	}
	buildReturns struct {
		result1 container.Instance
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 container.Instance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DockerBuilder) Build(arg1 string, arg2 *persistence.ChaincodePackageMetadata, arg3 io.Reader) (container.Instance, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 string
		arg2 *persistence.ChaincodePackageMetadata
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("Build", []interface{}{arg1, arg2, arg3})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DockerBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *DockerBuilder) BuildCalls(stub func(string, *persistence.ChaincodePackageMetadata, io.Reader) (container.Instance, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *DockerBuilder) BuildArgsForCall(i int) (string, *persistence.ChaincodePackageMetadata, io.Reader) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *DockerBuilder) BuildReturns(result1 container.Instance, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 container.Instance
		result2 error
	}{result1, result2}
}

func (fake *DockerBuilder) BuildReturnsOnCall(i int, result1 container.Instance, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 container.Instance
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 container.Instance
		result2 error
	}{result1, result2}
}

func (fake *DockerBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DockerBuilder) recordInvocation(key string, args []interface{}) {
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

var _ container.DockerBuilder = new(DockerBuilder)
