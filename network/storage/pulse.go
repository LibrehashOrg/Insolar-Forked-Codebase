//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package storage

import (
	"context"

	"github.com/insolar/insolar/insolar"
)

// PulseAccessor provides methods for accessing pulses.
//go:generate minimock -i github.com/insolar/insolar/network/storage.PulseAccessor -o ../../testutils/network -s _mock.go
type PulseAccessor interface {
	ForPulseNumber(context.Context, insolar.PulseNumber) (insolar.PulseNumber, error)
	Latest(ctx context.Context) (insolar.Pulse, error)
}

// PulseAppender provides method for appending pulses to storage.
//go:generate minimock -i github.com/insolar/insolar/network/storage.PulseAppender -o ../../testutils/network -s _mock.go
type PulseAppender interface {
	Append(ctx context.Context, pulse insolar.Pulse) error
}

// PulseCalculator performs calculations for pulses.
//go:generate minimock -i github.com/insolar/insolar/network/storage.PulseCalculator -o ../../testutils/network -s _mock.go
type PulseCalculator interface {
	Forwards(ctx context.Context, pn insolar.PulseNumber, steps int) (insolar.Pulse, error)
	Backwards(ctx context.Context, pn insolar.PulseNumber, steps int) (insolar.Pulse, error)
}

// PulseRangeHasher provides methods for hashing and validate pulse chain
//go:generate minimock -i github.com/insolar/insolar/network/storage.PulseRangeHasher -o ../../testutils/network -s _mock.go
type PulseRangeHasher interface {
	GetRangeHash(insolar.PulseRange) ([]byte, error)
	ValidateRangeHash(insolar.PulseRange, []byte) (bool, error)
}

// NewPulseStorage constructor creates PulseStorage
func NewPulseStorage() *PulseStorage {
	return &PulseStorage{}
}

type PulseStorage struct {
}

func (p *PulseStorage) GetRangeHash(insolar.PulseRange) ([]byte, error) {
	panic("implement me")
}

func (p *PulseStorage) ValidateRangeHash(insolar.PulseRange, []byte) (bool, error) {
	panic("implement me")
}

func (p *PulseStorage) Forwards(ctx context.Context, pn insolar.PulseNumber, steps int) (insolar.Pulse, error) {
	panic("implement me")
}

func (p *PulseStorage) Backwards(ctx context.Context, pn insolar.PulseNumber, steps int) (insolar.Pulse, error) {
	panic("implement me")
}

func (p *PulseStorage) Append(ctx context.Context, pulse insolar.Pulse) error {
	panic("implement me")
}

func (p *PulseStorage) ForPulseNumber(context.Context, insolar.PulseNumber) (insolar.PulseNumber, error) {
	panic("implement me")
}

func (p *PulseStorage) Latest(ctx context.Context) (insolar.Pulse, error) {
	panic("implement me")
}
