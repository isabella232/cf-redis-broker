// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/cf-redis-broker/redis"
)

type FakeLocalInstanceRepository struct {
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindByIDStub        func(string) (*redis.Instance, error)
	findByIDMutex       sync.RWMutex
	findByIDArgsForCall []struct {
		arg1 string
	}
	findByIDReturns struct {
		result1 *redis.Instance
		result2 error
	}
	findByIDReturnsOnCall map[int]struct {
		result1 *redis.Instance
		result2 error
	}
	InstanceConfigPathStub        func(string) string
	instanceConfigPathMutex       sync.RWMutex
	instanceConfigPathArgsForCall []struct {
		arg1 string
	}
	instanceConfigPathReturns struct {
		result1 string
	}
	instanceConfigPathReturnsOnCall map[int]struct {
		result1 string
	}
	InstanceCountStub        func() (int, []error)
	instanceCountMutex       sync.RWMutex
	instanceCountArgsForCall []struct {
	}
	instanceCountReturns struct {
		result1 int
		result2 []error
	}
	instanceCountReturnsOnCall map[int]struct {
		result1 int
		result2 []error
	}
	InstanceDataDirStub        func(string) string
	instanceDataDirMutex       sync.RWMutex
	instanceDataDirArgsForCall []struct {
		arg1 string
	}
	instanceDataDirReturns struct {
		result1 string
	}
	instanceDataDirReturnsOnCall map[int]struct {
		result1 string
	}
	InstanceExistsStub        func(string) (bool, error)
	instanceExistsMutex       sync.RWMutex
	instanceExistsArgsForCall []struct {
		arg1 string
	}
	instanceExistsReturns struct {
		result1 bool
		result2 error
	}
	instanceExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	InstanceLogFilePathStub        func(string) string
	instanceLogFilePathMutex       sync.RWMutex
	instanceLogFilePathArgsForCall []struct {
		arg1 string
	}
	instanceLogFilePathReturns struct {
		result1 string
	}
	instanceLogFilePathReturnsOnCall map[int]struct {
		result1 string
	}
	InstancePidFilePathStub        func(string) string
	instancePidFilePathMutex       sync.RWMutex
	instancePidFilePathArgsForCall []struct {
		arg1 string
	}
	instancePidFilePathReturns struct {
		result1 string
	}
	instancePidFilePathReturnsOnCall map[int]struct {
		result1 string
	}
	LockStub        func(*redis.Instance) error
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		arg1 *redis.Instance
	}
	lockReturns struct {
		result1 error
	}
	lockReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func(*redis.Instance) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 *redis.Instance
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	UnlockStub        func(*redis.Instance) error
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct {
		arg1 *redis.Instance
	}
	unlockReturns struct {
		result1 error
	}
	unlockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLocalInstanceRepository) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeLocalInstanceRepository) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeLocalInstanceRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeLocalInstanceRepository) FindByID(arg1 string) (*redis.Instance, error) {
	fake.findByIDMutex.Lock()
	ret, specificReturn := fake.findByIDReturnsOnCall[len(fake.findByIDArgsForCall)]
	fake.findByIDArgsForCall = append(fake.findByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindByID", []interface{}{arg1})
	fake.findByIDMutex.Unlock()
	if fake.FindByIDStub != nil {
		return fake.FindByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findByIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLocalInstanceRepository) FindByIDCallCount() int {
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	return len(fake.findByIDArgsForCall)
}

func (fake *FakeLocalInstanceRepository) FindByIDCalls(stub func(string) (*redis.Instance, error)) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = stub
}

func (fake *FakeLocalInstanceRepository) FindByIDArgsForCall(i int) string {
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	argsForCall := fake.findByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) FindByIDReturns(result1 *redis.Instance, result2 error) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = nil
	fake.findByIDReturns = struct {
		result1 *redis.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) FindByIDReturnsOnCall(i int, result1 *redis.Instance, result2 error) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = nil
	if fake.findByIDReturnsOnCall == nil {
		fake.findByIDReturnsOnCall = make(map[int]struct {
			result1 *redis.Instance
			result2 error
		})
	}
	fake.findByIDReturnsOnCall[i] = struct {
		result1 *redis.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPath(arg1 string) string {
	fake.instanceConfigPathMutex.Lock()
	ret, specificReturn := fake.instanceConfigPathReturnsOnCall[len(fake.instanceConfigPathArgsForCall)]
	fake.instanceConfigPathArgsForCall = append(fake.instanceConfigPathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstanceConfigPath", []interface{}{arg1})
	fake.instanceConfigPathMutex.Unlock()
	if fake.InstanceConfigPathStub != nil {
		return fake.InstanceConfigPathStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.instanceConfigPathReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPathCallCount() int {
	fake.instanceConfigPathMutex.RLock()
	defer fake.instanceConfigPathMutex.RUnlock()
	return len(fake.instanceConfigPathArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPathCalls(stub func(string) string) {
	fake.instanceConfigPathMutex.Lock()
	defer fake.instanceConfigPathMutex.Unlock()
	fake.InstanceConfigPathStub = stub
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPathArgsForCall(i int) string {
	fake.instanceConfigPathMutex.RLock()
	defer fake.instanceConfigPathMutex.RUnlock()
	argsForCall := fake.instanceConfigPathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPathReturns(result1 string) {
	fake.instanceConfigPathMutex.Lock()
	defer fake.instanceConfigPathMutex.Unlock()
	fake.InstanceConfigPathStub = nil
	fake.instanceConfigPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstanceConfigPathReturnsOnCall(i int, result1 string) {
	fake.instanceConfigPathMutex.Lock()
	defer fake.instanceConfigPathMutex.Unlock()
	fake.InstanceConfigPathStub = nil
	if fake.instanceConfigPathReturnsOnCall == nil {
		fake.instanceConfigPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.instanceConfigPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstanceCount() (int, []error) {
	fake.instanceCountMutex.Lock()
	ret, specificReturn := fake.instanceCountReturnsOnCall[len(fake.instanceCountArgsForCall)]
	fake.instanceCountArgsForCall = append(fake.instanceCountArgsForCall, struct {
	}{})
	fake.recordInvocation("InstanceCount", []interface{}{})
	fake.instanceCountMutex.Unlock()
	if fake.InstanceCountStub != nil {
		return fake.InstanceCountStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instanceCountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLocalInstanceRepository) InstanceCountCallCount() int {
	fake.instanceCountMutex.RLock()
	defer fake.instanceCountMutex.RUnlock()
	return len(fake.instanceCountArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstanceCountCalls(stub func() (int, []error)) {
	fake.instanceCountMutex.Lock()
	defer fake.instanceCountMutex.Unlock()
	fake.InstanceCountStub = stub
}

func (fake *FakeLocalInstanceRepository) InstanceCountReturns(result1 int, result2 []error) {
	fake.instanceCountMutex.Lock()
	defer fake.instanceCountMutex.Unlock()
	fake.InstanceCountStub = nil
	fake.instanceCountReturns = struct {
		result1 int
		result2 []error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) InstanceCountReturnsOnCall(i int, result1 int, result2 []error) {
	fake.instanceCountMutex.Lock()
	defer fake.instanceCountMutex.Unlock()
	fake.InstanceCountStub = nil
	if fake.instanceCountReturnsOnCall == nil {
		fake.instanceCountReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []error
		})
	}
	fake.instanceCountReturnsOnCall[i] = struct {
		result1 int
		result2 []error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) InstanceDataDir(arg1 string) string {
	fake.instanceDataDirMutex.Lock()
	ret, specificReturn := fake.instanceDataDirReturnsOnCall[len(fake.instanceDataDirArgsForCall)]
	fake.instanceDataDirArgsForCall = append(fake.instanceDataDirArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstanceDataDir", []interface{}{arg1})
	fake.instanceDataDirMutex.Unlock()
	if fake.InstanceDataDirStub != nil {
		return fake.InstanceDataDirStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.instanceDataDirReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) InstanceDataDirCallCount() int {
	fake.instanceDataDirMutex.RLock()
	defer fake.instanceDataDirMutex.RUnlock()
	return len(fake.instanceDataDirArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstanceDataDirCalls(stub func(string) string) {
	fake.instanceDataDirMutex.Lock()
	defer fake.instanceDataDirMutex.Unlock()
	fake.InstanceDataDirStub = stub
}

func (fake *FakeLocalInstanceRepository) InstanceDataDirArgsForCall(i int) string {
	fake.instanceDataDirMutex.RLock()
	defer fake.instanceDataDirMutex.RUnlock()
	argsForCall := fake.instanceDataDirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) InstanceDataDirReturns(result1 string) {
	fake.instanceDataDirMutex.Lock()
	defer fake.instanceDataDirMutex.Unlock()
	fake.InstanceDataDirStub = nil
	fake.instanceDataDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstanceDataDirReturnsOnCall(i int, result1 string) {
	fake.instanceDataDirMutex.Lock()
	defer fake.instanceDataDirMutex.Unlock()
	fake.InstanceDataDirStub = nil
	if fake.instanceDataDirReturnsOnCall == nil {
		fake.instanceDataDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.instanceDataDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstanceExists(arg1 string) (bool, error) {
	fake.instanceExistsMutex.Lock()
	ret, specificReturn := fake.instanceExistsReturnsOnCall[len(fake.instanceExistsArgsForCall)]
	fake.instanceExistsArgsForCall = append(fake.instanceExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstanceExists", []interface{}{arg1})
	fake.instanceExistsMutex.Unlock()
	if fake.InstanceExistsStub != nil {
		return fake.InstanceExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instanceExistsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLocalInstanceRepository) InstanceExistsCallCount() int {
	fake.instanceExistsMutex.RLock()
	defer fake.instanceExistsMutex.RUnlock()
	return len(fake.instanceExistsArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstanceExistsCalls(stub func(string) (bool, error)) {
	fake.instanceExistsMutex.Lock()
	defer fake.instanceExistsMutex.Unlock()
	fake.InstanceExistsStub = stub
}

func (fake *FakeLocalInstanceRepository) InstanceExistsArgsForCall(i int) string {
	fake.instanceExistsMutex.RLock()
	defer fake.instanceExistsMutex.RUnlock()
	argsForCall := fake.instanceExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) InstanceExistsReturns(result1 bool, result2 error) {
	fake.instanceExistsMutex.Lock()
	defer fake.instanceExistsMutex.Unlock()
	fake.InstanceExistsStub = nil
	fake.instanceExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) InstanceExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.instanceExistsMutex.Lock()
	defer fake.instanceExistsMutex.Unlock()
	fake.InstanceExistsStub = nil
	if fake.instanceExistsReturnsOnCall == nil {
		fake.instanceExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.instanceExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePath(arg1 string) string {
	fake.instanceLogFilePathMutex.Lock()
	ret, specificReturn := fake.instanceLogFilePathReturnsOnCall[len(fake.instanceLogFilePathArgsForCall)]
	fake.instanceLogFilePathArgsForCall = append(fake.instanceLogFilePathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstanceLogFilePath", []interface{}{arg1})
	fake.instanceLogFilePathMutex.Unlock()
	if fake.InstanceLogFilePathStub != nil {
		return fake.InstanceLogFilePathStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.instanceLogFilePathReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePathCallCount() int {
	fake.instanceLogFilePathMutex.RLock()
	defer fake.instanceLogFilePathMutex.RUnlock()
	return len(fake.instanceLogFilePathArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePathCalls(stub func(string) string) {
	fake.instanceLogFilePathMutex.Lock()
	defer fake.instanceLogFilePathMutex.Unlock()
	fake.InstanceLogFilePathStub = stub
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePathArgsForCall(i int) string {
	fake.instanceLogFilePathMutex.RLock()
	defer fake.instanceLogFilePathMutex.RUnlock()
	argsForCall := fake.instanceLogFilePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePathReturns(result1 string) {
	fake.instanceLogFilePathMutex.Lock()
	defer fake.instanceLogFilePathMutex.Unlock()
	fake.InstanceLogFilePathStub = nil
	fake.instanceLogFilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstanceLogFilePathReturnsOnCall(i int, result1 string) {
	fake.instanceLogFilePathMutex.Lock()
	defer fake.instanceLogFilePathMutex.Unlock()
	fake.InstanceLogFilePathStub = nil
	if fake.instanceLogFilePathReturnsOnCall == nil {
		fake.instanceLogFilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.instanceLogFilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePath(arg1 string) string {
	fake.instancePidFilePathMutex.Lock()
	ret, specificReturn := fake.instancePidFilePathReturnsOnCall[len(fake.instancePidFilePathArgsForCall)]
	fake.instancePidFilePathArgsForCall = append(fake.instancePidFilePathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstancePidFilePath", []interface{}{arg1})
	fake.instancePidFilePathMutex.Unlock()
	if fake.InstancePidFilePathStub != nil {
		return fake.InstancePidFilePathStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.instancePidFilePathReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePathCallCount() int {
	fake.instancePidFilePathMutex.RLock()
	defer fake.instancePidFilePathMutex.RUnlock()
	return len(fake.instancePidFilePathArgsForCall)
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePathCalls(stub func(string) string) {
	fake.instancePidFilePathMutex.Lock()
	defer fake.instancePidFilePathMutex.Unlock()
	fake.InstancePidFilePathStub = stub
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePathArgsForCall(i int) string {
	fake.instancePidFilePathMutex.RLock()
	defer fake.instancePidFilePathMutex.RUnlock()
	argsForCall := fake.instancePidFilePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePathReturns(result1 string) {
	fake.instancePidFilePathMutex.Lock()
	defer fake.instancePidFilePathMutex.Unlock()
	fake.InstancePidFilePathStub = nil
	fake.instancePidFilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) InstancePidFilePathReturnsOnCall(i int, result1 string) {
	fake.instancePidFilePathMutex.Lock()
	defer fake.instancePidFilePathMutex.Unlock()
	fake.InstancePidFilePathStub = nil
	if fake.instancePidFilePathReturnsOnCall == nil {
		fake.instancePidFilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.instancePidFilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalInstanceRepository) Lock(arg1 *redis.Instance) error {
	fake.lockMutex.Lock()
	ret, specificReturn := fake.lockReturnsOnCall[len(fake.lockArgsForCall)]
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		arg1 *redis.Instance
	}{arg1})
	fake.recordInvocation("Lock", []interface{}{arg1})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lockReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLocalInstanceRepository) LockCalls(stub func(*redis.Instance) error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = stub
}

func (fake *FakeLocalInstanceRepository) LockArgsForCall(i int) *redis.Instance {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	argsForCall := fake.lockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) LockReturns(result1 error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) LockReturnsOnCall(i int, result1 error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = nil
	if fake.lockReturnsOnCall == nil {
		fake.lockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.lockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) Setup(arg1 *redis.Instance) error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 *redis.Instance
	}{arg1})
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeLocalInstanceRepository) SetupCalls(stub func(*redis.Instance) error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *FakeLocalInstanceRepository) SetupArgsForCall(i int) *redis.Instance {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) SetupReturns(result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) SetupReturnsOnCall(i int, result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) Unlock(arg1 *redis.Instance) error {
	fake.unlockMutex.Lock()
	ret, specificReturn := fake.unlockReturnsOnCall[len(fake.unlockArgsForCall)]
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct {
		arg1 *redis.Instance
	}{arg1})
	fake.recordInvocation("Unlock", []interface{}{arg1})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		return fake.UnlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unlockReturns
	return fakeReturns.result1
}

func (fake *FakeLocalInstanceRepository) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *FakeLocalInstanceRepository) UnlockCalls(stub func(*redis.Instance) error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = stub
}

func (fake *FakeLocalInstanceRepository) UnlockArgsForCall(i int) *redis.Instance {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	argsForCall := fake.unlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLocalInstanceRepository) UnlockReturns(result1 error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = nil
	fake.unlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) UnlockReturnsOnCall(i int, result1 error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = nil
	if fake.unlockReturnsOnCall == nil {
		fake.unlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalInstanceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	fake.instanceConfigPathMutex.RLock()
	defer fake.instanceConfigPathMutex.RUnlock()
	fake.instanceCountMutex.RLock()
	defer fake.instanceCountMutex.RUnlock()
	fake.instanceDataDirMutex.RLock()
	defer fake.instanceDataDirMutex.RUnlock()
	fake.instanceExistsMutex.RLock()
	defer fake.instanceExistsMutex.RUnlock()
	fake.instanceLogFilePathMutex.RLock()
	defer fake.instanceLogFilePathMutex.RUnlock()
	fake.instancePidFilePathMutex.RLock()
	defer fake.instancePidFilePathMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLocalInstanceRepository) recordInvocation(key string, args []interface{}) {
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

var _ redis.LocalInstanceRepository = new(FakeLocalInstanceRepository)