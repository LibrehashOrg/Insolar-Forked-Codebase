package object

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/insolar/insolar/ledger/object.MemoryIndexModifier -o ./memory_index_modifier_mock.go

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
	record "github.com/insolar/insolar/insolar/record"
)

// MemoryIndexModifierMock implements MemoryIndexModifier
type MemoryIndexModifierMock struct {
	t minimock.Tester

	funcSet          func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	inspectFuncSet   func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mMemoryIndexModifierMockSet
}

// NewMemoryIndexModifierMock returns a mock for MemoryIndexModifier
func NewMemoryIndexModifierMock(t minimock.Tester) *MemoryIndexModifierMock {
	m := &MemoryIndexModifierMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SetMock = mMemoryIndexModifierMockSet{mock: m}
	m.SetMock.callArgs = []*MemoryIndexModifierMockSetParams{}

	return m
}

type mMemoryIndexModifierMockSet struct {
	mock               *MemoryIndexModifierMock
	defaultExpectation *MemoryIndexModifierMockSetExpectation
	expectations       []*MemoryIndexModifierMockSetExpectation

	callArgs []*MemoryIndexModifierMockSetParams
	mutex    sync.RWMutex
}

// MemoryIndexModifierMockSetExpectation specifies expectation struct of the MemoryIndexModifier.Set
type MemoryIndexModifierMockSetExpectation struct {
	mock   *MemoryIndexModifierMock
	params *MemoryIndexModifierMockSetParams

	Counter uint64
}

// MemoryIndexModifierMockSetParams contains parameters of the MemoryIndexModifier.Set
type MemoryIndexModifierMockSetParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	index record.Index
}

// Expect sets up expected params for MemoryIndexModifier.Set
func (mmSet *mMemoryIndexModifierMockSet) Expect(ctx context.Context, pn insolar.PulseNumber, index record.Index) *mMemoryIndexModifierMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("MemoryIndexModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &MemoryIndexModifierMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &MemoryIndexModifierMockSetParams{ctx, pn, index}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the MemoryIndexModifier.Set
func (mmSet *mMemoryIndexModifierMockSet) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *mMemoryIndexModifierMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for MemoryIndexModifierMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by MemoryIndexModifier.Set
func (mmSet *mMemoryIndexModifierMockSet) Return() *MemoryIndexModifierMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("MemoryIndexModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &MemoryIndexModifierMockSetExpectation{mock: mmSet.mock}
	}

	return mmSet.mock
}

//Set uses given function f to mock the MemoryIndexModifier.Set method
func (mmSet *mMemoryIndexModifierMockSet) Set(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *MemoryIndexModifierMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the MemoryIndexModifier.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the MemoryIndexModifier.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// Set implements MemoryIndexModifier
func (mmSet *MemoryIndexModifierMock) Set(ctx context.Context, pn insolar.PulseNumber, index record.Index) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, pn, index)
	}

	params := &MemoryIndexModifierMockSetParams{ctx, pn, index}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		want := mmSet.SetMock.defaultExpectation.params
		got := MemoryIndexModifierMockSetParams{ctx, pn, index}
		if want != nil && !minimock.Equal(*want, got) {
			mmSet.t.Errorf("MemoryIndexModifierMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmSet.funcSet != nil {
		mmSet.funcSet(ctx, pn, index)
		return
	}
	mmSet.t.Fatalf("Unexpected call to MemoryIndexModifierMock.Set. %v %v %v", ctx, pn, index)

}

// SetAfterCounter returns a count of finished MemoryIndexModifierMock.Set invocations
func (mmSet *MemoryIndexModifierMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of MemoryIndexModifierMock.Set invocations
func (mmSet *MemoryIndexModifierMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to MemoryIndexModifierMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mMemoryIndexModifierMockSet) Calls() []*MemoryIndexModifierMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*MemoryIndexModifierMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *MemoryIndexModifierMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *MemoryIndexModifierMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MemoryIndexModifierMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MemoryIndexModifierMock.Set")
		} else {
			m.t.Errorf("Expected call to MemoryIndexModifierMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to MemoryIndexModifierMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MemoryIndexModifierMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MemoryIndexModifierMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MemoryIndexModifierMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSetDone()
}