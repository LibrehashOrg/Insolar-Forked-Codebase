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

package s_sender

import (
	"context"

	"github.com/insolar/insolar/conveyor/smachine"
	"github.com/insolar/insolar/conveyor/smachine/main/example"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/network/storage"
)

// TODO[bigbes]: port it to state machine
type SenderService interface {
	bus.Sender
}

type SenderServiceAdapter struct {
	svc  SenderService
	exec smachine.ExecutionAdapter
}

func (a *SenderServiceAdapter) PrepareSync(ctx smachine.ExecutionContext, fn func(svc SenderService)) smachine.SyncCallRequester {
	return a.exec.PrepareSync(ctx, func() smachine.AsyncResultFunc {
		fn(a.svc)
		return nil
	})
}

func (a *SenderServiceAdapter) PrepareAsync(ctx smachine.ExecutionContext, fn func(svc SenderService) smachine.AsyncResultFunc) smachine.AsyncCallRequester {
	return a.exec.PrepareAsync(ctx, func() smachine.AsyncResultFunc {
		return fn(a.svc)
	})
}

func (a *SenderServiceAdapter) PrepareNotify(ctx smachine.ExecutionContext, fn func(svc SenderService)) smachine.NotifyRequester {
	return a.exec.PrepareNotify(ctx, func() { fn(a.svc) })
}

type senderService struct {
	bus.Sender
	Accessor pulse.Accessor
}

func NewSenderService(sender bus.Sender, accessor storage.PulseAccessor) *SenderServiceAdapter {
	ach := example.NewChannelAdapter(context.Background(), 0, -1)

	go func() {
		for {
			select {
			case <-ach.Context().Done():
				return
			case t := <-ach.Channel():
				t.RunAndSendResult()
			}
		}
	}()

	return &SenderServiceAdapter{
		senderService{
			Sender: sender,
		},
		smachine.NewExecutionAdapter("Sender", &ach),
	}
}