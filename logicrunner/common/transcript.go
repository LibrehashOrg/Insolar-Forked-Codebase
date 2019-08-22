package common

import (
	"context"
	"fmt"
	"reflect"

	"go.opencensus.io/trace"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
	"github.com/insolar/insolar/logicrunner/artifacts"
)

type Transcript struct {
	ObjectDescriptor artifacts.ObjectDescriptor
	Context          context.Context
	LogicContext     *insolar.LogicCallContext
	Request          *record.IncomingRequest
	RequestRef       insolar.Reference
	Nonce            uint64
	Deactivate       bool
	OutgoingRequests []OutgoingRequest
	FromLedger       bool
}

func NewTranscript(
	ctx context.Context,
	requestRef insolar.Reference,
	request record.IncomingRequest,
) *Transcript {

	return &Transcript{
		Context:    ctx,
		Request:    &request,
		RequestRef: requestRef,
		Nonce:      0,
		Deactivate: false,

		FromLedger: false,
	}
}

// NewTranscriptCloneContext creates a transcript with fresh context created from
// contextSource which can be either other Context or ServiceData. In general
// transcript shouldn't be created with context as execution can take minutes.
func NewTranscriptCloneContext(
	ctxSource interface{},
	requestRef insolar.Reference,
	request record.IncomingRequest,
) *Transcript {
	var ctx context.Context

	switch sourceTyped := ctxSource.(type) {
	case context.Context:
		ctx = freshContextFromContext(sourceTyped, request.APIRequestID)
	case payload.ServiceData:
		ctx = contextFromServiceData(sourceTyped)
	default:
		panic(fmt.Errorf("unexpected type of context source: %T", ctxSource))
	}

	objRef := request.Object
	if objRef == nil {
		objRef = &requestRef
	}
	ctx, _ = inslogger.WithFields(
		ctx,
		map[string]interface{}{
			"request": requestRef.String(),
			"object":  objRef.String(),
			"method":  request.Method,
		},
	)

	return NewTranscript(ctx, requestRef, request)
}

func (t *Transcript) AddOutgoingRequest(
	ctx context.Context, request record.IncomingRequest, result []byte, newObject *insolar.Reference, err error,
) {
	rec := OutgoingRequest{
		Request:   request,
		Response:  result,
		NewObject: newObject,
		Error:     err,
	}
	t.OutgoingRequests = append(t.OutgoingRequests, rec)
}

func (t *Transcript) HasOutgoingRequest(
	ctx context.Context, request record.IncomingRequest,
) *OutgoingRequest {
	for i := range t.OutgoingRequests {
		if reflect.DeepEqual(t.OutgoingRequests[i].Request, request) {
			return &t.OutgoingRequests[i]
		}
	}
	return nil
}

func contextFromServiceData(data payload.ServiceData) context.Context {
	ctx := inslogger.ContextWithTrace(context.Background(), data.LogTraceID)
	ctx = inslogger.WithLoggerLevel(ctx, data.LogLevel)
	if data.TraceSpanData != nil {
		parentSpan := instracer.MustDeserialize(data.TraceSpanData)
		return instracer.WithParentSpan(ctx, parentSpan)
	}
	return ctx
}

func freshContextFromContext(ctx context.Context, reqID string) context.Context {
	res := context.Background()

	res = inslogger.SetLogger(res, inslogger.FromContext(ctx))

	logLevel := inslogger.GetLoggerLevel(ctx)
	if logLevel != insolar.NoLevel {
		res = inslogger.WithLoggerLevel(res, logLevel)
	}

	// we know that trace id is equal to APIRequestID, in a few cases we
	// call this function and don't have correct context and trace id around
	// except in request record as APIRequestID field
	res = inslogger.ContextWithTrace(res, reqID)

	parentSpan, ok := instracer.ParentSpan(ctx)
	if ok {
		res = instracer.WithParentSpan(res, parentSpan)
	}

	if pctx := trace.FromContext(ctx); pctx != nil {
		res = trace.NewContext(res, pctx)
	}

	return res
}
