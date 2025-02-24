// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hxx258456/fabric/core/chaincode/platforms/util"
)

type Platform struct {
	DockerBuildOptionsStub        func(string) (util.DockerBuildOptions, error)
	dockerBuildOptionsMutex       sync.RWMutex
	dockerBuildOptionsArgsForCall []struct {
		arg1 string
	}
	dockerBuildOptionsReturns struct {
		result1 util.DockerBuildOptions
		result2 error
	}
	dockerBuildOptionsReturnsOnCall map[int]struct {
		result1 util.DockerBuildOptions
		result2 error
	}
	GenerateDockerfileStub        func() (string, error)
	generateDockerfileMutex       sync.RWMutex
	generateDockerfileArgsForCall []struct {
	}
	generateDockerfileReturns struct {
		result1 string
		result2 error
	}
	generateDockerfileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Platform) DockerBuildOptions(arg1 string) (util.DockerBuildOptions, error) {
	fake.dockerBuildOptionsMutex.Lock()
	ret, specificReturn := fake.dockerBuildOptionsReturnsOnCall[len(fake.dockerBuildOptionsArgsForCall)]
	fake.dockerBuildOptionsArgsForCall = append(fake.dockerBuildOptionsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DockerBuildOptions", []interface{}{arg1})
	fake.dockerBuildOptionsMutex.Unlock()
	if fake.DockerBuildOptionsStub != nil {
		return fake.DockerBuildOptionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dockerBuildOptionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Platform) DockerBuildOptionsCallCount() int {
	fake.dockerBuildOptionsMutex.RLock()
	defer fake.dockerBuildOptionsMutex.RUnlock()
	return len(fake.dockerBuildOptionsArgsForCall)
}

func (fake *Platform) DockerBuildOptionsCalls(stub func(string) (util.DockerBuildOptions, error)) {
	fake.dockerBuildOptionsMutex.Lock()
	defer fake.dockerBuildOptionsMutex.Unlock()
	fake.DockerBuildOptionsStub = stub
}

func (fake *Platform) DockerBuildOptionsArgsForCall(i int) string {
	fake.dockerBuildOptionsMutex.RLock()
	defer fake.dockerBuildOptionsMutex.RUnlock()
	argsForCall := fake.dockerBuildOptionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Platform) DockerBuildOptionsReturns(result1 util.DockerBuildOptions, result2 error) {
	fake.dockerBuildOptionsMutex.Lock()
	defer fake.dockerBuildOptionsMutex.Unlock()
	fake.DockerBuildOptionsStub = nil
	fake.dockerBuildOptionsReturns = struct {
		result1 util.DockerBuildOptions
		result2 error
	}{result1, result2}
}

func (fake *Platform) DockerBuildOptionsReturnsOnCall(i int, result1 util.DockerBuildOptions, result2 error) {
	fake.dockerBuildOptionsMutex.Lock()
	defer fake.dockerBuildOptionsMutex.Unlock()
	fake.DockerBuildOptionsStub = nil
	if fake.dockerBuildOptionsReturnsOnCall == nil {
		fake.dockerBuildOptionsReturnsOnCall = make(map[int]struct {
			result1 util.DockerBuildOptions
			result2 error
		})
	}
	fake.dockerBuildOptionsReturnsOnCall[i] = struct {
		result1 util.DockerBuildOptions
		result2 error
	}{result1, result2}
}

func (fake *Platform) GenerateDockerfile() (string, error) {
	fake.generateDockerfileMutex.Lock()
	ret, specificReturn := fake.generateDockerfileReturnsOnCall[len(fake.generateDockerfileArgsForCall)]
	fake.generateDockerfileArgsForCall = append(fake.generateDockerfileArgsForCall, struct {
	}{})
	fake.recordInvocation("GenerateDockerfile", []interface{}{})
	fake.generateDockerfileMutex.Unlock()
	if fake.GenerateDockerfileStub != nil {
		return fake.GenerateDockerfileStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateDockerfileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Platform) GenerateDockerfileCallCount() int {
	fake.generateDockerfileMutex.RLock()
	defer fake.generateDockerfileMutex.RUnlock()
	return len(fake.generateDockerfileArgsForCall)
}

func (fake *Platform) GenerateDockerfileCalls(stub func() (string, error)) {
	fake.generateDockerfileMutex.Lock()
	defer fake.generateDockerfileMutex.Unlock()
	fake.GenerateDockerfileStub = stub
}

func (fake *Platform) GenerateDockerfileReturns(result1 string, result2 error) {
	fake.generateDockerfileMutex.Lock()
	defer fake.generateDockerfileMutex.Unlock()
	fake.GenerateDockerfileStub = nil
	fake.generateDockerfileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Platform) GenerateDockerfileReturnsOnCall(i int, result1 string, result2 error) {
	fake.generateDockerfileMutex.Lock()
	defer fake.generateDockerfileMutex.Unlock()
	fake.GenerateDockerfileStub = nil
	if fake.generateDockerfileReturnsOnCall == nil {
		fake.generateDockerfileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateDockerfileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Platform) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *Platform) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *Platform) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *Platform) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *Platform) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Platform) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dockerBuildOptionsMutex.RLock()
	defer fake.dockerBuildOptionsMutex.RUnlock()
	fake.generateDockerfileMutex.RLock()
	defer fake.generateDockerfileMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Platform) recordInvocation(key string, args []interface{}) {
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
