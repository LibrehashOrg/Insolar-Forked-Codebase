package profiles

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "ActiveNode" can be found in github.com/insolar/insolar/network/consensus/gcpv2/api/profiles
*/
import (
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"
	insolar "github.com/insolar/insolar/insolar"
	cryptkit "github.com/insolar/insolar/network/consensus/common/cryptkit"
	member "github.com/insolar/insolar/network/consensus/gcpv2/api/member"
)

//ActiveNodeMock implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode
type ActiveNodeMock struct {
	t minimock.Tester

	CanIntroduceJoinerFunc       func() (r bool)
	CanIntroduceJoinerCounter    uint64
	CanIntroduceJoinerPreCounter uint64
	CanIntroduceJoinerMock       mActiveNodeMockCanIntroduceJoiner

	GetDeclaredPowerFunc       func() (r member.Power)
	GetDeclaredPowerCounter    uint64
	GetDeclaredPowerPreCounter uint64
	GetDeclaredPowerMock       mActiveNodeMockGetDeclaredPower

	GetIndexFunc       func() (r member.Index)
	GetIndexCounter    uint64
	GetIndexPreCounter uint64
	GetIndexMock       mActiveNodeMockGetIndex

	GetNodeIDFunc       func() (r insolar.ShortNodeID)
	GetNodeIDCounter    uint64
	GetNodeIDPreCounter uint64
	GetNodeIDMock       mActiveNodeMockGetNodeID

	GetOpModeFunc       func() (r member.OpMode)
	GetOpModeCounter    uint64
	GetOpModePreCounter uint64
	GetOpModeMock       mActiveNodeMockGetOpMode

	GetSignatureVerifierFunc       func() (r cryptkit.SignatureVerifier)
	GetSignatureVerifierCounter    uint64
	GetSignatureVerifierPreCounter uint64
	GetSignatureVerifierMock       mActiveNodeMockGetSignatureVerifier

	GetStaticFunc       func() (r StaticProfile)
	GetStaticCounter    uint64
	GetStaticPreCounter uint64
	GetStaticMock       mActiveNodeMockGetStatic

	HasFullProfileFunc       func() (r bool)
	HasFullProfileCounter    uint64
	HasFullProfilePreCounter uint64
	HasFullProfileMock       mActiveNodeMockHasFullProfile

	IsJoinerFunc       func() (r bool)
	IsJoinerCounter    uint64
	IsJoinerPreCounter uint64
	IsJoinerMock       mActiveNodeMockIsJoiner

	IsPoweredFunc       func() (r bool)
	IsPoweredCounter    uint64
	IsPoweredPreCounter uint64
	IsPoweredMock       mActiveNodeMockIsPowered

	IsStatefulFunc       func() (r bool)
	IsStatefulCounter    uint64
	IsStatefulPreCounter uint64
	IsStatefulMock       mActiveNodeMockIsStateful

	IsVoterFunc       func() (r bool)
	IsVoterCounter    uint64
	IsVoterPreCounter uint64
	IsVoterMock       mActiveNodeMockIsVoter
}

//NewActiveNodeMock returns a mock for github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode
func NewActiveNodeMock(t minimock.Tester) *ActiveNodeMock {
	m := &ActiveNodeMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CanIntroduceJoinerMock = mActiveNodeMockCanIntroduceJoiner{mock: m}
	m.GetDeclaredPowerMock = mActiveNodeMockGetDeclaredPower{mock: m}
	m.GetIndexMock = mActiveNodeMockGetIndex{mock: m}
	m.GetNodeIDMock = mActiveNodeMockGetNodeID{mock: m}
	m.GetOpModeMock = mActiveNodeMockGetOpMode{mock: m}
	m.GetSignatureVerifierMock = mActiveNodeMockGetSignatureVerifier{mock: m}
	m.GetStaticMock = mActiveNodeMockGetStatic{mock: m}
	m.HasFullProfileMock = mActiveNodeMockHasFullProfile{mock: m}
	m.IsJoinerMock = mActiveNodeMockIsJoiner{mock: m}
	m.IsPoweredMock = mActiveNodeMockIsPowered{mock: m}
	m.IsStatefulMock = mActiveNodeMockIsStateful{mock: m}
	m.IsVoterMock = mActiveNodeMockIsVoter{mock: m}

	return m
}

type mActiveNodeMockCanIntroduceJoiner struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockCanIntroduceJoinerExpectation
	expectationSeries []*ActiveNodeMockCanIntroduceJoinerExpectation
}

type ActiveNodeMockCanIntroduceJoinerExpectation struct {
	result *ActiveNodeMockCanIntroduceJoinerResult
}

type ActiveNodeMockCanIntroduceJoinerResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.CanIntroduceJoiner is expected from 1 to Infinity times
func (m *mActiveNodeMockCanIntroduceJoiner) Expect() *mActiveNodeMockCanIntroduceJoiner {
	m.mock.CanIntroduceJoinerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockCanIntroduceJoinerExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.CanIntroduceJoiner
func (m *mActiveNodeMockCanIntroduceJoiner) Return(r bool) *ActiveNodeMock {
	m.mock.CanIntroduceJoinerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockCanIntroduceJoinerExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockCanIntroduceJoinerResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.CanIntroduceJoiner is expected once
func (m *mActiveNodeMockCanIntroduceJoiner) ExpectOnce() *ActiveNodeMockCanIntroduceJoinerExpectation {
	m.mock.CanIntroduceJoinerFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockCanIntroduceJoinerExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockCanIntroduceJoinerExpectation) Return(r bool) {
	e.result = &ActiveNodeMockCanIntroduceJoinerResult{r}
}

//Set uses given function f as a mock of ActiveNode.CanIntroduceJoiner method
func (m *mActiveNodeMockCanIntroduceJoiner) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.CanIntroduceJoinerFunc = f
	return m.mock
}

//CanIntroduceJoiner implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) CanIntroduceJoiner() (r bool) {
	counter := atomic.AddUint64(&m.CanIntroduceJoinerPreCounter, 1)
	defer atomic.AddUint64(&m.CanIntroduceJoinerCounter, 1)

	if len(m.CanIntroduceJoinerMock.expectationSeries) > 0 {
		if counter > uint64(len(m.CanIntroduceJoinerMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.CanIntroduceJoiner.")
			return
		}

		result := m.CanIntroduceJoinerMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.CanIntroduceJoiner")
			return
		}

		r = result.r

		return
	}

	if m.CanIntroduceJoinerMock.mainExpectation != nil {

		result := m.CanIntroduceJoinerMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.CanIntroduceJoiner")
		}

		r = result.r

		return
	}

	if m.CanIntroduceJoinerFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.CanIntroduceJoiner.")
		return
	}

	return m.CanIntroduceJoinerFunc()
}

//CanIntroduceJoinerMinimockCounter returns a count of ActiveNodeMock.CanIntroduceJoinerFunc invocations
func (m *ActiveNodeMock) CanIntroduceJoinerMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CanIntroduceJoinerCounter)
}

//CanIntroduceJoinerMinimockPreCounter returns the value of ActiveNodeMock.CanIntroduceJoiner invocations
func (m *ActiveNodeMock) CanIntroduceJoinerMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CanIntroduceJoinerPreCounter)
}

//CanIntroduceJoinerFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) CanIntroduceJoinerFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.CanIntroduceJoinerMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.CanIntroduceJoinerCounter) == uint64(len(m.CanIntroduceJoinerMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.CanIntroduceJoinerMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.CanIntroduceJoinerCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.CanIntroduceJoinerFunc != nil {
		return atomic.LoadUint64(&m.CanIntroduceJoinerCounter) > 0
	}

	return true
}

type mActiveNodeMockGetDeclaredPower struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetDeclaredPowerExpectation
	expectationSeries []*ActiveNodeMockGetDeclaredPowerExpectation
}

type ActiveNodeMockGetDeclaredPowerExpectation struct {
	result *ActiveNodeMockGetDeclaredPowerResult
}

type ActiveNodeMockGetDeclaredPowerResult struct {
	r member.Power
}

//Expect specifies that invocation of ActiveNode.GetDeclaredPower is expected from 1 to Infinity times
func (m *mActiveNodeMockGetDeclaredPower) Expect() *mActiveNodeMockGetDeclaredPower {
	m.mock.GetDeclaredPowerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetDeclaredPowerExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetDeclaredPower
func (m *mActiveNodeMockGetDeclaredPower) Return(r member.Power) *ActiveNodeMock {
	m.mock.GetDeclaredPowerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetDeclaredPowerExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetDeclaredPowerResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetDeclaredPower is expected once
func (m *mActiveNodeMockGetDeclaredPower) ExpectOnce() *ActiveNodeMockGetDeclaredPowerExpectation {
	m.mock.GetDeclaredPowerFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetDeclaredPowerExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetDeclaredPowerExpectation) Return(r member.Power) {
	e.result = &ActiveNodeMockGetDeclaredPowerResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetDeclaredPower method
func (m *mActiveNodeMockGetDeclaredPower) Set(f func() (r member.Power)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetDeclaredPowerFunc = f
	return m.mock
}

//GetDeclaredPower implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetDeclaredPower() (r member.Power) {
	counter := atomic.AddUint64(&m.GetDeclaredPowerPreCounter, 1)
	defer atomic.AddUint64(&m.GetDeclaredPowerCounter, 1)

	if len(m.GetDeclaredPowerMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetDeclaredPowerMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetDeclaredPower.")
			return
		}

		result := m.GetDeclaredPowerMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetDeclaredPower")
			return
		}

		r = result.r

		return
	}

	if m.GetDeclaredPowerMock.mainExpectation != nil {

		result := m.GetDeclaredPowerMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetDeclaredPower")
		}

		r = result.r

		return
	}

	if m.GetDeclaredPowerFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetDeclaredPower.")
		return
	}

	return m.GetDeclaredPowerFunc()
}

//GetDeclaredPowerMinimockCounter returns a count of ActiveNodeMock.GetDeclaredPowerFunc invocations
func (m *ActiveNodeMock) GetDeclaredPowerMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetDeclaredPowerCounter)
}

//GetDeclaredPowerMinimockPreCounter returns the value of ActiveNodeMock.GetDeclaredPower invocations
func (m *ActiveNodeMock) GetDeclaredPowerMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetDeclaredPowerPreCounter)
}

//GetDeclaredPowerFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetDeclaredPowerFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetDeclaredPowerMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetDeclaredPowerCounter) == uint64(len(m.GetDeclaredPowerMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetDeclaredPowerMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetDeclaredPowerCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetDeclaredPowerFunc != nil {
		return atomic.LoadUint64(&m.GetDeclaredPowerCounter) > 0
	}

	return true
}

type mActiveNodeMockGetIndex struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetIndexExpectation
	expectationSeries []*ActiveNodeMockGetIndexExpectation
}

type ActiveNodeMockGetIndexExpectation struct {
	result *ActiveNodeMockGetIndexResult
}

type ActiveNodeMockGetIndexResult struct {
	r member.Index
}

//Expect specifies that invocation of ActiveNode.GetIndex is expected from 1 to Infinity times
func (m *mActiveNodeMockGetIndex) Expect() *mActiveNodeMockGetIndex {
	m.mock.GetIndexFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetIndexExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetIndex
func (m *mActiveNodeMockGetIndex) Return(r member.Index) *ActiveNodeMock {
	m.mock.GetIndexFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetIndexExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetIndexResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetIndex is expected once
func (m *mActiveNodeMockGetIndex) ExpectOnce() *ActiveNodeMockGetIndexExpectation {
	m.mock.GetIndexFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetIndexExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetIndexExpectation) Return(r member.Index) {
	e.result = &ActiveNodeMockGetIndexResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetIndex method
func (m *mActiveNodeMockGetIndex) Set(f func() (r member.Index)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetIndexFunc = f
	return m.mock
}

//GetIndex implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetIndex() (r member.Index) {
	counter := atomic.AddUint64(&m.GetIndexPreCounter, 1)
	defer atomic.AddUint64(&m.GetIndexCounter, 1)

	if len(m.GetIndexMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetIndexMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetIndex.")
			return
		}

		result := m.GetIndexMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetIndex")
			return
		}

		r = result.r

		return
	}

	if m.GetIndexMock.mainExpectation != nil {

		result := m.GetIndexMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetIndex")
		}

		r = result.r

		return
	}

	if m.GetIndexFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetIndex.")
		return
	}

	return m.GetIndexFunc()
}

//GetIndexMinimockCounter returns a count of ActiveNodeMock.GetIndexFunc invocations
func (m *ActiveNodeMock) GetIndexMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetIndexCounter)
}

//GetIndexMinimockPreCounter returns the value of ActiveNodeMock.GetIndex invocations
func (m *ActiveNodeMock) GetIndexMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetIndexPreCounter)
}

//GetIndexFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetIndexFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetIndexMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetIndexCounter) == uint64(len(m.GetIndexMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetIndexMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetIndexCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetIndexFunc != nil {
		return atomic.LoadUint64(&m.GetIndexCounter) > 0
	}

	return true
}

type mActiveNodeMockGetNodeID struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetNodeIDExpectation
	expectationSeries []*ActiveNodeMockGetNodeIDExpectation
}

type ActiveNodeMockGetNodeIDExpectation struct {
	result *ActiveNodeMockGetNodeIDResult
}

type ActiveNodeMockGetNodeIDResult struct {
	r insolar.ShortNodeID
}

//Expect specifies that invocation of ActiveNode.GetNodeID is expected from 1 to Infinity times
func (m *mActiveNodeMockGetNodeID) Expect() *mActiveNodeMockGetNodeID {
	m.mock.GetNodeIDFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetNodeIDExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetNodeID
func (m *mActiveNodeMockGetNodeID) Return(r insolar.ShortNodeID) *ActiveNodeMock {
	m.mock.GetNodeIDFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetNodeIDExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetNodeIDResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetNodeID is expected once
func (m *mActiveNodeMockGetNodeID) ExpectOnce() *ActiveNodeMockGetNodeIDExpectation {
	m.mock.GetNodeIDFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetNodeIDExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetNodeIDExpectation) Return(r insolar.ShortNodeID) {
	e.result = &ActiveNodeMockGetNodeIDResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetNodeID method
func (m *mActiveNodeMockGetNodeID) Set(f func() (r insolar.ShortNodeID)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetNodeIDFunc = f
	return m.mock
}

//GetNodeID implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetNodeID() (r insolar.ShortNodeID) {
	counter := atomic.AddUint64(&m.GetNodeIDPreCounter, 1)
	defer atomic.AddUint64(&m.GetNodeIDCounter, 1)

	if len(m.GetNodeIDMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetNodeIDMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetNodeID.")
			return
		}

		result := m.GetNodeIDMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetNodeID")
			return
		}

		r = result.r

		return
	}

	if m.GetNodeIDMock.mainExpectation != nil {

		result := m.GetNodeIDMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetNodeID")
		}

		r = result.r

		return
	}

	if m.GetNodeIDFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetNodeID.")
		return
	}

	return m.GetNodeIDFunc()
}

//GetNodeIDMinimockCounter returns a count of ActiveNodeMock.GetNodeIDFunc invocations
func (m *ActiveNodeMock) GetNodeIDMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetNodeIDCounter)
}

//GetNodeIDMinimockPreCounter returns the value of ActiveNodeMock.GetNodeID invocations
func (m *ActiveNodeMock) GetNodeIDMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetNodeIDPreCounter)
}

//GetNodeIDFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetNodeIDFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetNodeIDMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetNodeIDCounter) == uint64(len(m.GetNodeIDMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetNodeIDMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetNodeIDCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetNodeIDFunc != nil {
		return atomic.LoadUint64(&m.GetNodeIDCounter) > 0
	}

	return true
}

type mActiveNodeMockGetOpMode struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetOpModeExpectation
	expectationSeries []*ActiveNodeMockGetOpModeExpectation
}

type ActiveNodeMockGetOpModeExpectation struct {
	result *ActiveNodeMockGetOpModeResult
}

type ActiveNodeMockGetOpModeResult struct {
	r member.OpMode
}

//Expect specifies that invocation of ActiveNode.GetOpMode is expected from 1 to Infinity times
func (m *mActiveNodeMockGetOpMode) Expect() *mActiveNodeMockGetOpMode {
	m.mock.GetOpModeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetOpModeExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetOpMode
func (m *mActiveNodeMockGetOpMode) Return(r member.OpMode) *ActiveNodeMock {
	m.mock.GetOpModeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetOpModeExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetOpModeResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetOpMode is expected once
func (m *mActiveNodeMockGetOpMode) ExpectOnce() *ActiveNodeMockGetOpModeExpectation {
	m.mock.GetOpModeFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetOpModeExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetOpModeExpectation) Return(r member.OpMode) {
	e.result = &ActiveNodeMockGetOpModeResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetOpMode method
func (m *mActiveNodeMockGetOpMode) Set(f func() (r member.OpMode)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetOpModeFunc = f
	return m.mock
}

//GetOpMode implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetOpMode() (r member.OpMode) {
	counter := atomic.AddUint64(&m.GetOpModePreCounter, 1)
	defer atomic.AddUint64(&m.GetOpModeCounter, 1)

	if len(m.GetOpModeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetOpModeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetOpMode.")
			return
		}

		result := m.GetOpModeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetOpMode")
			return
		}

		r = result.r

		return
	}

	if m.GetOpModeMock.mainExpectation != nil {

		result := m.GetOpModeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetOpMode")
		}

		r = result.r

		return
	}

	if m.GetOpModeFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetOpMode.")
		return
	}

	return m.GetOpModeFunc()
}

//GetOpModeMinimockCounter returns a count of ActiveNodeMock.GetOpModeFunc invocations
func (m *ActiveNodeMock) GetOpModeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetOpModeCounter)
}

//GetOpModeMinimockPreCounter returns the value of ActiveNodeMock.GetOpMode invocations
func (m *ActiveNodeMock) GetOpModeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetOpModePreCounter)
}

//GetOpModeFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetOpModeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetOpModeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetOpModeCounter) == uint64(len(m.GetOpModeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetOpModeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetOpModeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetOpModeFunc != nil {
		return atomic.LoadUint64(&m.GetOpModeCounter) > 0
	}

	return true
}

type mActiveNodeMockGetSignatureVerifier struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetSignatureVerifierExpectation
	expectationSeries []*ActiveNodeMockGetSignatureVerifierExpectation
}

type ActiveNodeMockGetSignatureVerifierExpectation struct {
	result *ActiveNodeMockGetSignatureVerifierResult
}

type ActiveNodeMockGetSignatureVerifierResult struct {
	r cryptkit.SignatureVerifier
}

//Expect specifies that invocation of ActiveNode.GetSignatureVerifier is expected from 1 to Infinity times
func (m *mActiveNodeMockGetSignatureVerifier) Expect() *mActiveNodeMockGetSignatureVerifier {
	m.mock.GetSignatureVerifierFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetSignatureVerifierExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetSignatureVerifier
func (m *mActiveNodeMockGetSignatureVerifier) Return(r cryptkit.SignatureVerifier) *ActiveNodeMock {
	m.mock.GetSignatureVerifierFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetSignatureVerifierExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetSignatureVerifierResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetSignatureVerifier is expected once
func (m *mActiveNodeMockGetSignatureVerifier) ExpectOnce() *ActiveNodeMockGetSignatureVerifierExpectation {
	m.mock.GetSignatureVerifierFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetSignatureVerifierExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetSignatureVerifierExpectation) Return(r cryptkit.SignatureVerifier) {
	e.result = &ActiveNodeMockGetSignatureVerifierResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetSignatureVerifier method
func (m *mActiveNodeMockGetSignatureVerifier) Set(f func() (r cryptkit.SignatureVerifier)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetSignatureVerifierFunc = f
	return m.mock
}

//GetSignatureVerifier implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetSignatureVerifier() (r cryptkit.SignatureVerifier) {
	counter := atomic.AddUint64(&m.GetSignatureVerifierPreCounter, 1)
	defer atomic.AddUint64(&m.GetSignatureVerifierCounter, 1)

	if len(m.GetSignatureVerifierMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetSignatureVerifierMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetSignatureVerifier.")
			return
		}

		result := m.GetSignatureVerifierMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetSignatureVerifier")
			return
		}

		r = result.r

		return
	}

	if m.GetSignatureVerifierMock.mainExpectation != nil {

		result := m.GetSignatureVerifierMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetSignatureVerifier")
		}

		r = result.r

		return
	}

	if m.GetSignatureVerifierFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetSignatureVerifier.")
		return
	}

	return m.GetSignatureVerifierFunc()
}

//GetSignatureVerifierMinimockCounter returns a count of ActiveNodeMock.GetSignatureVerifierFunc invocations
func (m *ActiveNodeMock) GetSignatureVerifierMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetSignatureVerifierCounter)
}

//GetSignatureVerifierMinimockPreCounter returns the value of ActiveNodeMock.GetSignatureVerifier invocations
func (m *ActiveNodeMock) GetSignatureVerifierMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetSignatureVerifierPreCounter)
}

//GetSignatureVerifierFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetSignatureVerifierFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetSignatureVerifierMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetSignatureVerifierCounter) == uint64(len(m.GetSignatureVerifierMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetSignatureVerifierMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetSignatureVerifierCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetSignatureVerifierFunc != nil {
		return atomic.LoadUint64(&m.GetSignatureVerifierCounter) > 0
	}

	return true
}

type mActiveNodeMockGetStatic struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockGetStaticExpectation
	expectationSeries []*ActiveNodeMockGetStaticExpectation
}

type ActiveNodeMockGetStaticExpectation struct {
	result *ActiveNodeMockGetStaticResult
}

type ActiveNodeMockGetStaticResult struct {
	r StaticProfile
}

//Expect specifies that invocation of ActiveNode.GetStatic is expected from 1 to Infinity times
func (m *mActiveNodeMockGetStatic) Expect() *mActiveNodeMockGetStatic {
	m.mock.GetStaticFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetStaticExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.GetStatic
func (m *mActiveNodeMockGetStatic) Return(r StaticProfile) *ActiveNodeMock {
	m.mock.GetStaticFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockGetStaticExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockGetStaticResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.GetStatic is expected once
func (m *mActiveNodeMockGetStatic) ExpectOnce() *ActiveNodeMockGetStaticExpectation {
	m.mock.GetStaticFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockGetStaticExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockGetStaticExpectation) Return(r StaticProfile) {
	e.result = &ActiveNodeMockGetStaticResult{r}
}

//Set uses given function f as a mock of ActiveNode.GetStatic method
func (m *mActiveNodeMockGetStatic) Set(f func() (r StaticProfile)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetStaticFunc = f
	return m.mock
}

//GetStatic implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) GetStatic() (r StaticProfile) {
	counter := atomic.AddUint64(&m.GetStaticPreCounter, 1)
	defer atomic.AddUint64(&m.GetStaticCounter, 1)

	if len(m.GetStaticMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetStaticMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.GetStatic.")
			return
		}

		result := m.GetStaticMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetStatic")
			return
		}

		r = result.r

		return
	}

	if m.GetStaticMock.mainExpectation != nil {

		result := m.GetStaticMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.GetStatic")
		}

		r = result.r

		return
	}

	if m.GetStaticFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.GetStatic.")
		return
	}

	return m.GetStaticFunc()
}

//GetStaticMinimockCounter returns a count of ActiveNodeMock.GetStaticFunc invocations
func (m *ActiveNodeMock) GetStaticMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetStaticCounter)
}

//GetStaticMinimockPreCounter returns the value of ActiveNodeMock.GetStatic invocations
func (m *ActiveNodeMock) GetStaticMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetStaticPreCounter)
}

//GetStaticFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) GetStaticFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetStaticMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetStaticCounter) == uint64(len(m.GetStaticMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetStaticMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetStaticCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetStaticFunc != nil {
		return atomic.LoadUint64(&m.GetStaticCounter) > 0
	}

	return true
}

type mActiveNodeMockHasFullProfile struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockHasFullProfileExpectation
	expectationSeries []*ActiveNodeMockHasFullProfileExpectation
}

type ActiveNodeMockHasFullProfileExpectation struct {
	result *ActiveNodeMockHasFullProfileResult
}

type ActiveNodeMockHasFullProfileResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.HasFullProfile is expected from 1 to Infinity times
func (m *mActiveNodeMockHasFullProfile) Expect() *mActiveNodeMockHasFullProfile {
	m.mock.HasFullProfileFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockHasFullProfileExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.HasFullProfile
func (m *mActiveNodeMockHasFullProfile) Return(r bool) *ActiveNodeMock {
	m.mock.HasFullProfileFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockHasFullProfileExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockHasFullProfileResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.HasFullProfile is expected once
func (m *mActiveNodeMockHasFullProfile) ExpectOnce() *ActiveNodeMockHasFullProfileExpectation {
	m.mock.HasFullProfileFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockHasFullProfileExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockHasFullProfileExpectation) Return(r bool) {
	e.result = &ActiveNodeMockHasFullProfileResult{r}
}

//Set uses given function f as a mock of ActiveNode.HasFullProfile method
func (m *mActiveNodeMockHasFullProfile) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.HasFullProfileFunc = f
	return m.mock
}

//HasFullProfile implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) HasFullProfile() (r bool) {
	counter := atomic.AddUint64(&m.HasFullProfilePreCounter, 1)
	defer atomic.AddUint64(&m.HasFullProfileCounter, 1)

	if len(m.HasFullProfileMock.expectationSeries) > 0 {
		if counter > uint64(len(m.HasFullProfileMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.HasFullProfile.")
			return
		}

		result := m.HasFullProfileMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.HasFullProfile")
			return
		}

		r = result.r

		return
	}

	if m.HasFullProfileMock.mainExpectation != nil {

		result := m.HasFullProfileMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.HasFullProfile")
		}

		r = result.r

		return
	}

	if m.HasFullProfileFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.HasFullProfile.")
		return
	}

	return m.HasFullProfileFunc()
}

//HasFullProfileMinimockCounter returns a count of ActiveNodeMock.HasFullProfileFunc invocations
func (m *ActiveNodeMock) HasFullProfileMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HasFullProfileCounter)
}

//HasFullProfileMinimockPreCounter returns the value of ActiveNodeMock.HasFullProfile invocations
func (m *ActiveNodeMock) HasFullProfileMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.HasFullProfilePreCounter)
}

//HasFullProfileFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) HasFullProfileFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.HasFullProfileMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.HasFullProfileCounter) == uint64(len(m.HasFullProfileMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.HasFullProfileMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.HasFullProfileCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.HasFullProfileFunc != nil {
		return atomic.LoadUint64(&m.HasFullProfileCounter) > 0
	}

	return true
}

type mActiveNodeMockIsJoiner struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockIsJoinerExpectation
	expectationSeries []*ActiveNodeMockIsJoinerExpectation
}

type ActiveNodeMockIsJoinerExpectation struct {
	result *ActiveNodeMockIsJoinerResult
}

type ActiveNodeMockIsJoinerResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.IsJoiner is expected from 1 to Infinity times
func (m *mActiveNodeMockIsJoiner) Expect() *mActiveNodeMockIsJoiner {
	m.mock.IsJoinerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsJoinerExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.IsJoiner
func (m *mActiveNodeMockIsJoiner) Return(r bool) *ActiveNodeMock {
	m.mock.IsJoinerFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsJoinerExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockIsJoinerResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.IsJoiner is expected once
func (m *mActiveNodeMockIsJoiner) ExpectOnce() *ActiveNodeMockIsJoinerExpectation {
	m.mock.IsJoinerFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockIsJoinerExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockIsJoinerExpectation) Return(r bool) {
	e.result = &ActiveNodeMockIsJoinerResult{r}
}

//Set uses given function f as a mock of ActiveNode.IsJoiner method
func (m *mActiveNodeMockIsJoiner) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsJoinerFunc = f
	return m.mock
}

//IsJoiner implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) IsJoiner() (r bool) {
	counter := atomic.AddUint64(&m.IsJoinerPreCounter, 1)
	defer atomic.AddUint64(&m.IsJoinerCounter, 1)

	if len(m.IsJoinerMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsJoinerMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.IsJoiner.")
			return
		}

		result := m.IsJoinerMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsJoiner")
			return
		}

		r = result.r

		return
	}

	if m.IsJoinerMock.mainExpectation != nil {

		result := m.IsJoinerMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsJoiner")
		}

		r = result.r

		return
	}

	if m.IsJoinerFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.IsJoiner.")
		return
	}

	return m.IsJoinerFunc()
}

//IsJoinerMinimockCounter returns a count of ActiveNodeMock.IsJoinerFunc invocations
func (m *ActiveNodeMock) IsJoinerMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsJoinerCounter)
}

//IsJoinerMinimockPreCounter returns the value of ActiveNodeMock.IsJoiner invocations
func (m *ActiveNodeMock) IsJoinerMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsJoinerPreCounter)
}

//IsJoinerFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) IsJoinerFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsJoinerMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsJoinerCounter) == uint64(len(m.IsJoinerMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsJoinerMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsJoinerCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsJoinerFunc != nil {
		return atomic.LoadUint64(&m.IsJoinerCounter) > 0
	}

	return true
}

type mActiveNodeMockIsPowered struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockIsPoweredExpectation
	expectationSeries []*ActiveNodeMockIsPoweredExpectation
}

type ActiveNodeMockIsPoweredExpectation struct {
	result *ActiveNodeMockIsPoweredResult
}

type ActiveNodeMockIsPoweredResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.IsPowered is expected from 1 to Infinity times
func (m *mActiveNodeMockIsPowered) Expect() *mActiveNodeMockIsPowered {
	m.mock.IsPoweredFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsPoweredExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.IsPowered
func (m *mActiveNodeMockIsPowered) Return(r bool) *ActiveNodeMock {
	m.mock.IsPoweredFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsPoweredExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockIsPoweredResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.IsPowered is expected once
func (m *mActiveNodeMockIsPowered) ExpectOnce() *ActiveNodeMockIsPoweredExpectation {
	m.mock.IsPoweredFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockIsPoweredExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockIsPoweredExpectation) Return(r bool) {
	e.result = &ActiveNodeMockIsPoweredResult{r}
}

//Set uses given function f as a mock of ActiveNode.IsPowered method
func (m *mActiveNodeMockIsPowered) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsPoweredFunc = f
	return m.mock
}

//IsPowered implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) IsPowered() (r bool) {
	counter := atomic.AddUint64(&m.IsPoweredPreCounter, 1)
	defer atomic.AddUint64(&m.IsPoweredCounter, 1)

	if len(m.IsPoweredMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsPoweredMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.IsPowered.")
			return
		}

		result := m.IsPoweredMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsPowered")
			return
		}

		r = result.r

		return
	}

	if m.IsPoweredMock.mainExpectation != nil {

		result := m.IsPoweredMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsPowered")
		}

		r = result.r

		return
	}

	if m.IsPoweredFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.IsPowered.")
		return
	}

	return m.IsPoweredFunc()
}

//IsPoweredMinimockCounter returns a count of ActiveNodeMock.IsPoweredFunc invocations
func (m *ActiveNodeMock) IsPoweredMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsPoweredCounter)
}

//IsPoweredMinimockPreCounter returns the value of ActiveNodeMock.IsPowered invocations
func (m *ActiveNodeMock) IsPoweredMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsPoweredPreCounter)
}

//IsPoweredFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) IsPoweredFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsPoweredMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsPoweredCounter) == uint64(len(m.IsPoweredMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsPoweredMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsPoweredCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsPoweredFunc != nil {
		return atomic.LoadUint64(&m.IsPoweredCounter) > 0
	}

	return true
}

type mActiveNodeMockIsStateful struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockIsStatefulExpectation
	expectationSeries []*ActiveNodeMockIsStatefulExpectation
}

type ActiveNodeMockIsStatefulExpectation struct {
	result *ActiveNodeMockIsStatefulResult
}

type ActiveNodeMockIsStatefulResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.IsStateful is expected from 1 to Infinity times
func (m *mActiveNodeMockIsStateful) Expect() *mActiveNodeMockIsStateful {
	m.mock.IsStatefulFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsStatefulExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.IsStateful
func (m *mActiveNodeMockIsStateful) Return(r bool) *ActiveNodeMock {
	m.mock.IsStatefulFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsStatefulExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockIsStatefulResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.IsStateful is expected once
func (m *mActiveNodeMockIsStateful) ExpectOnce() *ActiveNodeMockIsStatefulExpectation {
	m.mock.IsStatefulFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockIsStatefulExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockIsStatefulExpectation) Return(r bool) {
	e.result = &ActiveNodeMockIsStatefulResult{r}
}

//Set uses given function f as a mock of ActiveNode.IsStateful method
func (m *mActiveNodeMockIsStateful) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsStatefulFunc = f
	return m.mock
}

//IsStateful implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) IsStateful() (r bool) {
	counter := atomic.AddUint64(&m.IsStatefulPreCounter, 1)
	defer atomic.AddUint64(&m.IsStatefulCounter, 1)

	if len(m.IsStatefulMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsStatefulMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.IsStateful.")
			return
		}

		result := m.IsStatefulMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsStateful")
			return
		}

		r = result.r

		return
	}

	if m.IsStatefulMock.mainExpectation != nil {

		result := m.IsStatefulMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsStateful")
		}

		r = result.r

		return
	}

	if m.IsStatefulFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.IsStateful.")
		return
	}

	return m.IsStatefulFunc()
}

//IsStatefulMinimockCounter returns a count of ActiveNodeMock.IsStatefulFunc invocations
func (m *ActiveNodeMock) IsStatefulMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsStatefulCounter)
}

//IsStatefulMinimockPreCounter returns the value of ActiveNodeMock.IsStateful invocations
func (m *ActiveNodeMock) IsStatefulMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsStatefulPreCounter)
}

//IsStatefulFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) IsStatefulFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsStatefulMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsStatefulCounter) == uint64(len(m.IsStatefulMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsStatefulMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsStatefulCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsStatefulFunc != nil {
		return atomic.LoadUint64(&m.IsStatefulCounter) > 0
	}

	return true
}

type mActiveNodeMockIsVoter struct {
	mock              *ActiveNodeMock
	mainExpectation   *ActiveNodeMockIsVoterExpectation
	expectationSeries []*ActiveNodeMockIsVoterExpectation
}

type ActiveNodeMockIsVoterExpectation struct {
	result *ActiveNodeMockIsVoterResult
}

type ActiveNodeMockIsVoterResult struct {
	r bool
}

//Expect specifies that invocation of ActiveNode.IsVoter is expected from 1 to Infinity times
func (m *mActiveNodeMockIsVoter) Expect() *mActiveNodeMockIsVoter {
	m.mock.IsVoterFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsVoterExpectation{}
	}

	return m
}

//Return specifies results of invocation of ActiveNode.IsVoter
func (m *mActiveNodeMockIsVoter) Return(r bool) *ActiveNodeMock {
	m.mock.IsVoterFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveNodeMockIsVoterExpectation{}
	}
	m.mainExpectation.result = &ActiveNodeMockIsVoterResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveNode.IsVoter is expected once
func (m *mActiveNodeMockIsVoter) ExpectOnce() *ActiveNodeMockIsVoterExpectation {
	m.mock.IsVoterFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveNodeMockIsVoterExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveNodeMockIsVoterExpectation) Return(r bool) {
	e.result = &ActiveNodeMockIsVoterResult{r}
}

//Set uses given function f as a mock of ActiveNode.IsVoter method
func (m *mActiveNodeMockIsVoter) Set(f func() (r bool)) *ActiveNodeMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.IsVoterFunc = f
	return m.mock
}

//IsVoter implements github.com/insolar/insolar/network/consensus/gcpv2/api/profiles.ActiveNode interface
func (m *ActiveNodeMock) IsVoter() (r bool) {
	counter := atomic.AddUint64(&m.IsVoterPreCounter, 1)
	defer atomic.AddUint64(&m.IsVoterCounter, 1)

	if len(m.IsVoterMock.expectationSeries) > 0 {
		if counter > uint64(len(m.IsVoterMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveNodeMock.IsVoter.")
			return
		}

		result := m.IsVoterMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsVoter")
			return
		}

		r = result.r

		return
	}

	if m.IsVoterMock.mainExpectation != nil {

		result := m.IsVoterMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveNodeMock.IsVoter")
		}

		r = result.r

		return
	}

	if m.IsVoterFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveNodeMock.IsVoter.")
		return
	}

	return m.IsVoterFunc()
}

//IsVoterMinimockCounter returns a count of ActiveNodeMock.IsVoterFunc invocations
func (m *ActiveNodeMock) IsVoterMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.IsVoterCounter)
}

//IsVoterMinimockPreCounter returns the value of ActiveNodeMock.IsVoter invocations
func (m *ActiveNodeMock) IsVoterMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.IsVoterPreCounter)
}

//IsVoterFinished returns true if mock invocations count is ok
func (m *ActiveNodeMock) IsVoterFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.IsVoterMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.IsVoterCounter) == uint64(len(m.IsVoterMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.IsVoterMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.IsVoterCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.IsVoterFunc != nil {
		return atomic.LoadUint64(&m.IsVoterCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ActiveNodeMock) ValidateCallCounters() {

	if !m.CanIntroduceJoinerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.CanIntroduceJoiner")
	}

	if !m.GetDeclaredPowerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetDeclaredPower")
	}

	if !m.GetIndexFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetIndex")
	}

	if !m.GetNodeIDFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetNodeID")
	}

	if !m.GetOpModeFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetOpMode")
	}

	if !m.GetSignatureVerifierFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetSignatureVerifier")
	}

	if !m.GetStaticFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetStatic")
	}

	if !m.HasFullProfileFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.HasFullProfile")
	}

	if !m.IsJoinerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsJoiner")
	}

	if !m.IsPoweredFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsPowered")
	}

	if !m.IsStatefulFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsStateful")
	}

	if !m.IsVoterFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsVoter")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ActiveNodeMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ActiveNodeMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ActiveNodeMock) MinimockFinish() {

	if !m.CanIntroduceJoinerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.CanIntroduceJoiner")
	}

	if !m.GetDeclaredPowerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetDeclaredPower")
	}

	if !m.GetIndexFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetIndex")
	}

	if !m.GetNodeIDFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetNodeID")
	}

	if !m.GetOpModeFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetOpMode")
	}

	if !m.GetSignatureVerifierFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetSignatureVerifier")
	}

	if !m.GetStaticFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.GetStatic")
	}

	if !m.HasFullProfileFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.HasFullProfile")
	}

	if !m.IsJoinerFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsJoiner")
	}

	if !m.IsPoweredFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsPowered")
	}

	if !m.IsStatefulFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsStateful")
	}

	if !m.IsVoterFinished() {
		m.t.Fatal("Expected call to ActiveNodeMock.IsVoter")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ActiveNodeMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ActiveNodeMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.CanIntroduceJoinerFinished()
		ok = ok && m.GetDeclaredPowerFinished()
		ok = ok && m.GetIndexFinished()
		ok = ok && m.GetNodeIDFinished()
		ok = ok && m.GetOpModeFinished()
		ok = ok && m.GetSignatureVerifierFinished()
		ok = ok && m.GetStaticFinished()
		ok = ok && m.HasFullProfileFinished()
		ok = ok && m.IsJoinerFinished()
		ok = ok && m.IsPoweredFinished()
		ok = ok && m.IsStatefulFinished()
		ok = ok && m.IsVoterFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.CanIntroduceJoinerFinished() {
				m.t.Error("Expected call to ActiveNodeMock.CanIntroduceJoiner")
			}

			if !m.GetDeclaredPowerFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetDeclaredPower")
			}

			if !m.GetIndexFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetIndex")
			}

			if !m.GetNodeIDFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetNodeID")
			}

			if !m.GetOpModeFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetOpMode")
			}

			if !m.GetSignatureVerifierFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetSignatureVerifier")
			}

			if !m.GetStaticFinished() {
				m.t.Error("Expected call to ActiveNodeMock.GetStatic")
			}

			if !m.HasFullProfileFinished() {
				m.t.Error("Expected call to ActiveNodeMock.HasFullProfile")
			}

			if !m.IsJoinerFinished() {
				m.t.Error("Expected call to ActiveNodeMock.IsJoiner")
			}

			if !m.IsPoweredFinished() {
				m.t.Error("Expected call to ActiveNodeMock.IsPowered")
			}

			if !m.IsStatefulFinished() {
				m.t.Error("Expected call to ActiveNodeMock.IsStateful")
			}

			if !m.IsVoterFinished() {
				m.t.Error("Expected call to ActiveNodeMock.IsVoter")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *ActiveNodeMock) AllMocksCalled() bool {

	if !m.CanIntroduceJoinerFinished() {
		return false
	}

	if !m.GetDeclaredPowerFinished() {
		return false
	}

	if !m.GetIndexFinished() {
		return false
	}

	if !m.GetNodeIDFinished() {
		return false
	}

	if !m.GetOpModeFinished() {
		return false
	}

	if !m.GetSignatureVerifierFinished() {
		return false
	}

	if !m.GetStaticFinished() {
		return false
	}

	if !m.HasFullProfileFinished() {
		return false
	}

	if !m.IsJoinerFinished() {
		return false
	}

	if !m.IsPoweredFinished() {
		return false
	}

	if !m.IsStatefulFinished() {
		return false
	}

	if !m.IsVoterFinished() {
		return false
	}

	return true
}
