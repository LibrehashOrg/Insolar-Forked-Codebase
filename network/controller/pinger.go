/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package controller

import (
	"time"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/network/hostnetwork"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/insolar/insolar/network/transport/packet/types"
	"github.com/pkg/errors"
)

// Pinger is a light and stateless component that can ping remote host to receive its NodeID
type Pinger struct {
	transport hostnetwork.InternalTransport
}

// PingWithTimeout ping remote host with timeout
func (p *Pinger) Ping(address string, timeout time.Duration) (core.RecordRef, error) {
	request := p.transport.NewRequestBuilder().Type(types.Ping).Build()
	addr, err := host.NewAddress(address)
	if err != nil {
		return core.RecordRef{}, errors.Wrapf(err, "failed to resolve address %s", address)
	}
	h := host.NewHost(addr)
	future, err := p.transport.SendRequestPacket(request, h)
	if err != nil {
		return core.RecordRef{}, errors.Wrapf(err, "failed to ping address %s", address)
	}
	result, err := future.GetResponse(timeout)
	if err != nil {
		return core.RecordRef{}, errors.Wrapf(err, "failed to receive ping response from address %s", address)
	}
	return result.GetSender(), nil
}

func NewPinger(transport hostnetwork.InternalTransport) *Pinger {
	return &Pinger{transport: transport}
}
