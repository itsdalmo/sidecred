// Code generated by counterfeiter. DO NOT EDIT.
package githubfakes

import (
	"context"
	"sync"

	"github.com/telia-oss/githubapp"
	"github.com/telia-oss/sidecred/store/github"
)

type FakeApp struct {
	CreateInstallationTokenStub        func(context.Context, string, []string, *githubapp.Permissions) (*githubapp.Token, error)
	createInstallationTokenMutex       sync.RWMutex
	createInstallationTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []string
		arg4 *githubapp.Permissions
	}
	createInstallationTokenReturns struct {
		result1 *githubapp.Token
		result2 error
	}
	createInstallationTokenReturnsOnCall map[int]struct {
		result1 *githubapp.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApp) CreateInstallationToken(arg1 context.Context, arg2 string, arg3 []string, arg4 *githubapp.Permissions) (*githubapp.Token, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.createInstallationTokenMutex.Lock()
	ret, specificReturn := fake.createInstallationTokenReturnsOnCall[len(fake.createInstallationTokenArgsForCall)]
	fake.createInstallationTokenArgsForCall = append(fake.createInstallationTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []string
		arg4 *githubapp.Permissions
	}{arg1, arg2, arg3Copy, arg4})
	stub := fake.CreateInstallationTokenStub
	fakeReturns := fake.createInstallationTokenReturns
	fake.recordInvocation("CreateInstallationToken", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.createInstallationTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApp) CreateInstallationTokenCallCount() int {
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	return len(fake.createInstallationTokenArgsForCall)
}

func (fake *FakeApp) CreateInstallationTokenCalls(stub func(context.Context, string, []string, *githubapp.Permissions) (*githubapp.Token, error)) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = stub
}

func (fake *FakeApp) CreateInstallationTokenArgsForCall(i int) (context.Context, string, []string, *githubapp.Permissions) {
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	argsForCall := fake.createInstallationTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeApp) CreateInstallationTokenReturns(result1 *githubapp.Token, result2 error) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = nil
	fake.createInstallationTokenReturns = struct {
		result1 *githubapp.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeApp) CreateInstallationTokenReturnsOnCall(i int, result1 *githubapp.Token, result2 error) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = nil
	if fake.createInstallationTokenReturnsOnCall == nil {
		fake.createInstallationTokenReturnsOnCall = make(map[int]struct {
			result1 *githubapp.Token
			result2 error
		})
	}
	fake.createInstallationTokenReturnsOnCall[i] = struct {
		result1 *githubapp.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeApp) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApp) recordInvocation(key string, args []interface{}) {
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

var _ github.App = new(FakeApp)
