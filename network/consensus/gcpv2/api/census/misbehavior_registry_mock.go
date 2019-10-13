package census

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/misbehavior"
)

// MisbehaviorRegistryMock implements MisbehaviorRegistry
type MisbehaviorRegistryMock struct {
	t minimock.Tester

	funcAddReport          func(report misbehavior.Report)
	inspectFuncAddReport   func(report misbehavior.Report)
	afterAddReportCounter  uint64
	beforeAddReportCounter uint64
	AddReportMock          mMisbehaviorRegistryMockAddReport
}

// NewMisbehaviorRegistryMock returns a mock for MisbehaviorRegistry
func NewMisbehaviorRegistryMock(t minimock.Tester) *MisbehaviorRegistryMock {
	m := &MisbehaviorRegistryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddReportMock = mMisbehaviorRegistryMockAddReport{mock: m}
	m.AddReportMock.callArgs = []*MisbehaviorRegistryMockAddReportParams{}

	return m
}

type mMisbehaviorRegistryMockAddReport struct {
	mock               *MisbehaviorRegistryMock
	defaultExpectation *MisbehaviorRegistryMockAddReportExpectation
	expectations       []*MisbehaviorRegistryMockAddReportExpectation

	callArgs []*MisbehaviorRegistryMockAddReportParams
	mutex    sync.RWMutex
}

// MisbehaviorRegistryMockAddReportExpectation specifies expectation struct of the MisbehaviorRegistry.AddReport
type MisbehaviorRegistryMockAddReportExpectation struct {
	mock   *MisbehaviorRegistryMock
	params *MisbehaviorRegistryMockAddReportParams

	Counter uint64
}

// MisbehaviorRegistryMockAddReportParams contains parameters of the MisbehaviorRegistry.AddReport
type MisbehaviorRegistryMockAddReportParams struct {
	report misbehavior.Report
}

// Expect sets up expected params for MisbehaviorRegistry.AddReport
func (mmAddReport *mMisbehaviorRegistryMockAddReport) Expect(report misbehavior.Report) *mMisbehaviorRegistryMockAddReport {
	if mmAddReport.mock.funcAddReport != nil {
		mmAddReport.mock.t.Fatalf("MisbehaviorRegistryMock.AddReport mock is already set by Set")
	}

	if mmAddReport.defaultExpectation == nil {
		mmAddReport.defaultExpectation = &MisbehaviorRegistryMockAddReportExpectation{}
	}

	mmAddReport.defaultExpectation.params = &MisbehaviorRegistryMockAddReportParams{report}
	for _, e := range mmAddReport.expectations {
		if minimock.Equal(e.params, mmAddReport.defaultExpectation.params) {
			mmAddReport.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddReport.defaultExpectation.params)
		}
	}

	return mmAddReport
}

// Inspect accepts an inspector function that has same arguments as the MisbehaviorRegistry.AddReport
func (mmAddReport *mMisbehaviorRegistryMockAddReport) Inspect(f func(report misbehavior.Report)) *mMisbehaviorRegistryMockAddReport {
	if mmAddReport.mock.inspectFuncAddReport != nil {
		mmAddReport.mock.t.Fatalf("Inspect function is already set for MisbehaviorRegistryMock.AddReport")
	}

	mmAddReport.mock.inspectFuncAddReport = f

	return mmAddReport
}

// Return sets up results that will be returned by MisbehaviorRegistry.AddReport
func (mmAddReport *mMisbehaviorRegistryMockAddReport) Return() *MisbehaviorRegistryMock {
	if mmAddReport.mock.funcAddReport != nil {
		mmAddReport.mock.t.Fatalf("MisbehaviorRegistryMock.AddReport mock is already set by Set")
	}

	if mmAddReport.defaultExpectation == nil {
		mmAddReport.defaultExpectation = &MisbehaviorRegistryMockAddReportExpectation{mock: mmAddReport.mock}
	}

	return mmAddReport.mock
}

//Set uses given function f to mock the MisbehaviorRegistry.AddReport method
func (mmAddReport *mMisbehaviorRegistryMockAddReport) Set(f func(report misbehavior.Report)) *MisbehaviorRegistryMock {
	if mmAddReport.defaultExpectation != nil {
		mmAddReport.mock.t.Fatalf("Default expectation is already set for the MisbehaviorRegistry.AddReport method")
	}

	if len(mmAddReport.expectations) > 0 {
		mmAddReport.mock.t.Fatalf("Some expectations are already set for the MisbehaviorRegistry.AddReport method")
	}

	mmAddReport.mock.funcAddReport = f
	return mmAddReport.mock
}

// AddReport implements MisbehaviorRegistry
func (mmAddReport *MisbehaviorRegistryMock) AddReport(report misbehavior.Report) {
	mm_atomic.AddUint64(&mmAddReport.beforeAddReportCounter, 1)
	defer mm_atomic.AddUint64(&mmAddReport.afterAddReportCounter, 1)

	if mmAddReport.inspectFuncAddReport != nil {
		mmAddReport.inspectFuncAddReport(report)
	}

	mm_params := &MisbehaviorRegistryMockAddReportParams{report}

	// Record call args
	mmAddReport.AddReportMock.mutex.Lock()
	mmAddReport.AddReportMock.callArgs = append(mmAddReport.AddReportMock.callArgs, mm_params)
	mmAddReport.AddReportMock.mutex.Unlock()

	for _, e := range mmAddReport.AddReportMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmAddReport.AddReportMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddReport.AddReportMock.defaultExpectation.Counter, 1)
		mm_want := mmAddReport.AddReportMock.defaultExpectation.params
		mm_got := MisbehaviorRegistryMockAddReportParams{report}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddReport.t.Errorf("MisbehaviorRegistryMock.AddReport got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmAddReport.funcAddReport != nil {
		mmAddReport.funcAddReport(report)
		return
	}
	mmAddReport.t.Fatalf("Unexpected call to MisbehaviorRegistryMock.AddReport. %v", report)

}

// AddReportAfterCounter returns a count of finished MisbehaviorRegistryMock.AddReport invocations
func (mmAddReport *MisbehaviorRegistryMock) AddReportAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddReport.afterAddReportCounter)
}

// AddReportBeforeCounter returns a count of MisbehaviorRegistryMock.AddReport invocations
func (mmAddReport *MisbehaviorRegistryMock) AddReportBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddReport.beforeAddReportCounter)
}

// Calls returns a list of arguments used in each call to MisbehaviorRegistryMock.AddReport.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddReport *mMisbehaviorRegistryMockAddReport) Calls() []*MisbehaviorRegistryMockAddReportParams {
	mmAddReport.mutex.RLock()

	argCopy := make([]*MisbehaviorRegistryMockAddReportParams, len(mmAddReport.callArgs))
	copy(argCopy, mmAddReport.callArgs)

	mmAddReport.mutex.RUnlock()

	return argCopy
}

// MinimockAddReportDone returns true if the count of the AddReport invocations corresponds
// the number of defined expectations
func (m *MisbehaviorRegistryMock) MinimockAddReportDone() bool {
	for _, e := range m.AddReportMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddReportMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddReportCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddReport != nil && mm_atomic.LoadUint64(&m.afterAddReportCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddReportInspect logs each unmet expectation
func (m *MisbehaviorRegistryMock) MinimockAddReportInspect() {
	for _, e := range m.AddReportMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MisbehaviorRegistryMock.AddReport with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddReportMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddReportCounter) < 1 {
		if m.AddReportMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MisbehaviorRegistryMock.AddReport")
		} else {
			m.t.Errorf("Expected call to MisbehaviorRegistryMock.AddReport with params: %#v", *m.AddReportMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddReport != nil && mm_atomic.LoadUint64(&m.afterAddReportCounter) < 1 {
		m.t.Error("Expected call to MisbehaviorRegistryMock.AddReport")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MisbehaviorRegistryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddReportInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MisbehaviorRegistryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MisbehaviorRegistryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddReportDone()
}
