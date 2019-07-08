//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package artifacts

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	wmmsg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/insolar/payload"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/internal/ledger/store"

	"github.com/insolar/insolar/component"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/delegationtoken"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/node"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/drop"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
)

type amSuite struct {
	suite.Suite

	cm  *component.Manager
	ctx context.Context

	scheme insolar.PlatformCryptographyScheme

	nodeStorage  node.Accessor
	jetStorage   jet.Storage
	dropModifier drop.Modifier
	dropAccessor drop.Accessor
}

func NewAmSuite() *amSuite {
	return &amSuite{
		Suite: suite.Suite{},
	}
}

// Init and run suite
func TestArtifactManager(t *testing.T) {
	suite.Run(t, NewAmSuite())
}

func (s *amSuite) BeforeTest(suiteName, testName string) {
	s.cm = &component.Manager{}
	s.ctx = inslogger.TestContext(s.T())

	s.scheme = platformpolicy.NewPlatformCryptographyScheme()
	s.jetStorage = jet.NewStore()
	s.nodeStorage = node.NewStorage()

	dbStore := store.NewMemoryMockDB()
	dropStorage := drop.NewDB(dbStore)
	s.dropAccessor = dropStorage
	s.dropModifier = dropStorage

	s.cm.Inject(
		s.scheme,
		store.NewMemoryMockDB(),
		s.jetStorage,
		s.nodeStorage,
		pulse.NewStorageMem(),
		s.dropAccessor,
		s.dropModifier,
	)

	err := s.cm.Init(s.ctx)
	if err != nil {
		s.T().Error("ComponentManager init failed", err)
	}
	err = s.cm.Start(s.ctx)
	if err != nil {
		s.T().Error("ComponentManager start failed", err)
	}
}

func (s *amSuite) AfterTest(suiteName, testName string) {
	err := s.cm.Stop(s.ctx)
	if err != nil {
		s.T().Error("ComponentManager stop failed", err)
	}
}

func genRandomID(pulse insolar.PulseNumber) *insolar.ID {
	buff := [insolar.RecordIDSize - insolar.PulseNumberSize]byte{}
	_, err := rand.Read(buff[:])
	if err != nil {
		panic(err)
	}
	return insolar.NewID(pulse, buff[:])
}

func genRefWithID(id *insolar.ID) *insolar.Reference {
	return insolar.NewReference(*id)
}

func genRandomRef(pulse insolar.PulseNumber) *insolar.Reference {
	return genRefWithID(genRandomID(pulse))
}

func (s *amSuite) TestLedgerArtifactManager_GetChildren_FollowsRedirect() {
	mc := minimock.NewController(s.T())
	am := NewClient(nil)
	mb := testutils.NewMessageBusMock(mc)

	objRef := genRandomRef(0)
	nodeRef := genRandomRef(0)
	mb.SendFunc = func(c context.Context, m insolar.Message, o *insolar.MessageSendOptions) (r insolar.Reply, r1 error) {
		o = o.Safe()
		if o.Receiver == nil {
			return &reply.GetChildrenRedirectReply{
				Receiver: nodeRef,
				Token:    &delegationtoken.GetChildrenRedirectToken{Signature: []byte{1, 2, 3}},
			}, nil
		}

		token, ok := o.Token.(*delegationtoken.GetChildrenRedirectToken)
		assert.True(s.T(), ok)
		assert.Equal(s.T(), []byte{1, 2, 3}, token.Signature)
		assert.Equal(s.T(), nodeRef, o.Receiver)
		return &reply.Children{}, nil
	}
	am.DefaultBus = mb

	pa := pulse.NewAccessorMock(s.T())
	pa.LatestMock.Return(*insolar.GenesisPulse, nil)
	am.PulseAccessor = pa

	_, err := am.GetChildren(s.ctx, *objRef, nil)
	require.NoError(s.T(), err)
}

func (s *amSuite) TestLedgerArtifactManager_GetRequest_Success() {
	// Arrange
	mc := minimock.NewController(s.T())
	defer mc.Finish()
	objectID := testutils.RandomID()
	requestID := testutils.RandomID()

	node := testutils.RandomRef()

	jc := jet.NewCoordinatorMock(mc)
	jc.NodeForObjectMock.Return(&node, nil)

	pulseAccessor := pulse.NewAccessorMock(s.T())
	pulseAccessor.LatestMock.Return(*insolar.GenesisPulse, nil)

	req := record.IncomingRequest{
		Method: "test",
	}
	virtRec := record.Wrap(req)
	data, err := virtRec.Marshal()
	require.NoError(s.T(), err)
	finalResponse := &reply.Request{Record: data}

	mb := testutils.NewMessageBusMock(s.T())
	mb.SendFunc = func(p context.Context, p1 insolar.Message, p2 *insolar.MessageSendOptions) (r insolar.Reply, r1 error) {
		switch mb.SendCounter {
		case 0:
			casted, ok := p1.(*message.GetPendingRequestID)
			require.Equal(s.T(), true, ok)
			require.Equal(s.T(), objectID, casted.ObjectID)
			return &reply.ID{ID: requestID}, nil
		case 1:
			casted, ok := p1.(*message.GetRequest)
			require.Equal(s.T(), true, ok)
			require.Equal(s.T(), requestID, casted.Request)
			require.Equal(s.T(), node, *p2.Receiver)
			return finalResponse, nil
		default:
			panic("test is totally broken")
		}
	}

	am := NewClient(nil)
	am.JetCoordinator = jc
	am.DefaultBus = mb
	am.PulseAccessor = pulseAccessor

	// Act
	_, res, err := am.GetPendingRequest(inslogger.TestContext(s.T()), objectID)

	// Assert
	require.NoError(s.T(), err)
	require.Equal(s.T(), "test", res.Message().(*message.CallMethod).Method)
}

// Send msg, bus.Sender gets error and closes resp chan
func TestRetryerSend_SendErrored(t *testing.T) {
	sender := &bus.SenderMock{}
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		res := make(chan *wmmsg.Message)
		close(res)
		return res, func() {}
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		return insolar.Pulse{PulseNumber: 10}, nil
	}
	c.PulseAccessor = p

	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	defer done()
	for range reps {
		require.Fail(t, "we are not expect any replays")
	}
}

// Send msg, close reply channel by timeout
func TestRetryerSend_Send_Timeout(t *testing.T) {
	once := sync.Once{}
	sender := &bus.SenderMock{}
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		done := func() {
			once.Do(func() { close(innerReps) })
		}
		go func() {
			time.Sleep(time.Second * 2)
			done()
		}()
		return innerReps, done
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		return insolar.Pulse{PulseNumber: 10}, nil
	}
	c.PulseAccessor = p

	reps, _ := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	select {
	case _, ok := <-reps:
		require.False(t, ok, "channel with replies must be closed, without any messages received")
	}
}

// Send msg, client stops waiting for response before request was actually done
func TestRetryerSend_Send_ClientDone(t *testing.T) {
	sender := &bus.SenderMock{}
	c := NewClient(sender)

	// code from r.send() func
	r := retryer{
		ppl:           &payload.State{},
		role:          insolar.DynamicRoleLightExecutor,
		ref:           testutils.RandomRef(),
		tries:         3,
		sender:        c.sender,
		channelClosed: make(chan interface{}),
		replyChan:     make(chan *wmmsg.Message),
		isDone:        false,
	}
	r.clientDone = func() {
		r.once.Do(func() {
			close(r.channelClosed)
			r.processingStarted.Lock()
			close(r.replyChan)
			r.isDone = true
			r.processingStarted.Unlock()
		})
	}
	r.clientDone()
	r.send(context.Background())

	for range r.replyChan {
		require.Fail(t, "we are not expect any replays")
	}
}

func sendTestReply(pl payload.Payload, ch chan<- *wmmsg.Message, isDone chan<- interface{}) {
	msg, _ := payload.NewMessage(pl)
	meta := payload.Meta{
		Payload: msg.Payload,
	}
	buf, _ := meta.Marshal()
	msg.Payload = buf
	ch <- msg
	close(isDone)
}

// Send msg, get one response
func TestRetryerSend(t *testing.T) {
	sender := &bus.SenderMock{}
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		return innerReps, func() { close(innerReps) }
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		return insolar.Pulse{PulseNumber: 10}, nil
	}
	c.PulseAccessor = p

	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)

	isDone := make(chan<- interface{})
	go sendTestReply(&payload.Error{Text: "object is deactivated", Code: payload.CodeUnknown}, innerReps, isDone)

	var success bool
	for rep := range reps {
		replyPayload, err := payload.UnmarshalFromMeta(rep.Payload)
		require.Nil(t, err)

		switch p := replyPayload.(type) {
		case *payload.Error:
			switch p.Code {
			case payload.CodeUnknown:
				success = true
			}
		}

		if success {
			break
		}
	}
	done()

	ok := true
	select {
	case _, ok = <-innerReps:
	default:
	}
	require.False(t, ok)
}

// Send msg, get "flow cancelled" error, than get one response
func TestRetryerSend_FlowCancelled_Once(t *testing.T) {
	sender := bus.NewSenderMock(t)
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		innerReps = make(chan *wmmsg.Message)
		if sender.SendRoleCounter == 0 {
			go sendTestReply(&payload.Error{Text: "test error", Code: payload.CodeFlowCanceled}, innerReps, make(chan<- interface{}))
		} else {
			go sendTestReply(&payload.State{}, innerReps, make(chan<- interface{}))
		}
		return innerReps, func() { close(innerReps) }
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	pulseNumber := 10
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		pulseNumber = pulseNumber + 10
		return insolar.Pulse{PulseNumber: insolar.PulseNumber(pulseNumber)}, nil
	}
	c.PulseAccessor = p

	var success bool
	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	defer done()
	for rep := range reps {
		replyPayload, _ := payload.UnmarshalFromMeta(rep.Payload)

		switch replyPayload.(type) {
		case *payload.State:
			success = true
		}

		if success {
			break
		}
	}
	done()

	ok := true
	select {
	case _, ok = <-innerReps:
	default:
	}
	require.False(t, ok)
}

// Send msg, get "flow cancelled" error, than get two responses
func TestRetryerSend_FlowCancelled_Once_SeveralReply(t *testing.T) {
	sender := bus.NewSenderMock(t)
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		innerReps = make(chan *wmmsg.Message)
		if sender.SendRoleCounter == 0 {
			go sendTestReply(&payload.Error{Text: "test error", Code: payload.CodeFlowCanceled}, innerReps, make(chan<- interface{}))
		} else {
			go sendTestReply(&payload.State{}, innerReps, make(chan<- interface{}))
			go sendTestReply(&payload.State{}, innerReps, make(chan<- interface{}))
		}
		return innerReps, func() { close(innerReps) }
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	pulseNumber := 10
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		pulseNumber = pulseNumber + 10
		return insolar.Pulse{PulseNumber: insolar.PulseNumber(pulseNumber)}, nil
	}
	c.PulseAccessor = p

	var success int
	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	for rep := range reps {
		replyPayload, _ := payload.UnmarshalFromMeta(rep.Payload)

		switch replyPayload.(type) {
		case *payload.State:
			success = success + 1
		}

		if success == 2 {
			break
		}
	}
	done()

	ok := true
	select {
	case _, ok = <-innerReps:
	default:
	}
	require.False(t, ok)
}

// Send msg, get "flow cancelled" error on every tries
func TestRetryerSend_FlowCancelled_RetryExceeded(t *testing.T) {
	sender := bus.NewSenderMock(t)
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		innerReps = make(chan *wmmsg.Message)
		go sendTestReply(&payload.Error{Text: "test error", Code: payload.CodeFlowCanceled}, innerReps, make(chan<- interface{}))
		return innerReps, func() { close(innerReps) }
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	pulseNumber := 10
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		pulseNumber = pulseNumber + 10
		return insolar.Pulse{PulseNumber: insolar.PulseNumber(pulseNumber)}, nil
	}
	c.PulseAccessor = p

	var success bool
	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	for range reps {
		success = true
		break
	}
	require.False(t, success)

	done()

	ok := true
	select {
	case _, ok = <-innerReps:
	default:
	}
	require.False(t, ok)
}

// Send msg, get response, than get "flow cancelled" error, than get two responses
func TestRetryerSend_FlowCancelled_Between(t *testing.T) {
	sender := bus.NewSenderMock(t)
	innerReps := make(chan *wmmsg.Message)
	sender.SendRoleFunc = func(p context.Context, p1 *wmmsg.Message, p2 insolar.DynamicRole, p3 insolar.Reference) (r <-chan *wmmsg.Message, r1 func()) {
		innerReps = make(chan *wmmsg.Message)
		if sender.SendRoleCounter == 0 {
			go func() {
				isDone := make(chan interface{})
				go sendTestReply(&payload.State{}, innerReps, isDone)
				<-isDone
				go sendTestReply(&payload.Error{Text: "test error", Code: payload.CodeFlowCanceled}, innerReps, make(chan<- interface{}))
			}()
		} else {
			go func() {
				isDone := make(chan interface{})
				go sendTestReply(&payload.State{}, innerReps, isDone)
				<-isDone
				go sendTestReply(&payload.State{}, innerReps, make(chan<- interface{}))
			}()
		}
		return innerReps, func() { close(innerReps) }
	}
	c := NewClient(sender)
	p := pulse.NewAccessorMock(t)
	pulseNumber := 10
	p.LatestFunc = func(p context.Context) (r insolar.Pulse, r1 error) {
		pulseNumber = pulseNumber + 10
		return insolar.Pulse{PulseNumber: insolar.PulseNumber(pulseNumber)}, nil
	}
	c.PulseAccessor = p

	var success int
	reps, done := c.retryerSend(context.Background(), &payload.State{}, insolar.DynamicRoleLightExecutor, testutils.RandomRef(), 3)
	for rep := range reps {
		replyPayload, _ := payload.UnmarshalFromMeta(rep.Payload)

		switch replyPayload.(type) {
		case *payload.State:
			success = success + 1
		default:
		}

		if success == 3 {
			break
		}
	}

	done()

	ok := true
	select {
	case _, ok = <-innerReps:
	default:
	}
	require.False(t, ok)
}
