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

package misbehavior

import (
	"errors"
	"testing"

	"github.com/insolar/insolar/network/consensus/common/cryptkit"

	"github.com/insolar/insolar/network/consensus/common/endpoints"

	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"

	"github.com/stretchr/testify/require"
)

func TestIsFraud(t *testing.T) {
	fe := FraudError{}
	require.True(t, IsFraud(&fe))

	err := errors.New("test")
	require.False(t, IsFraud(err))
}

func TestFraudOf(t *testing.T) {
	fe := FraudError{}
	require.Equal(t, &fe, FraudOf(&fe))

	err := errors.New("test")
	require.True(t, FraudOf(err) == nil)
}

func TestFraudIsUnknown(t *testing.T) {
	fe := FraudError{}
	require.True(t, fe.IsUnknown())

	fe.fraudType = 1
	require.False(t, fe.IsUnknown())
}

func TestFraudMisbehaviorType(t *testing.T) {
	fe := FraudError{fraudType: 0}
	require.Equal(t, Type(1<<33), fe.MisbehaviorType())

	fe.fraudType = 1
	require.Equal(t, Type((1<<33)+1), fe.MisbehaviorType())
}

func TestFraudCaptureMark(t *testing.T) {
	cm := interface{}(1)
	fe := FraudError{captureMark: cm}

	require.Equal(t, cm, fe.CaptureMark())
}

func TestFraudDetails(t *testing.T) {
	dets := []interface{}{1, 2}
	fe := &FraudError{details: dets}
	require.Equal(t, dets, fe.Details())
}

func TestFraudViolatorNode(t *testing.T) {
	bn := profiles.NewBaseNodeMock(t)
	fe := &FraudError{violatorNode: bn}

	require.Equal(t, bn, fe.ViolatorNode())
}

func TestFraudViolatorHost(t *testing.T) {
	inc := endpoints.InboundConnection{}
	fe := &FraudError{violatorHost: inc}
	require.Equal(t, inc, fe.ViolatorHost())
}

func TestFraudType(t *testing.T) {
	ft := 1
	be := &FraudError{fraudType: ft}
	require.Equal(t, ft, be.FraudType())
}

func TestFraudError(t *testing.T) {
	fe := &FraudError{}
	require.True(t, fe.Error() != "")

	bn := profiles.NewBaseNodeMock(t)
	fe.violatorNode = bn
	require.True(t, fe.Error() != "")

	fe.captureMark = 1
	require.True(t, fe.Error() != "")
}

func TestNewFraudFactory(t *testing.T) {
	ff := NewFraudFactory(reportFunc)
	require.True(t, ff.capture != nil)
}

func TestNewFraud(t *testing.T) {
	bf := NewFraudFactory(reportFunc)
	fraudType := 1
	msg := "test"
	inc := endpoints.NewInboundMock(t)
	violatorHost := inc
	bn := profiles.NewBaseNodeMock(t)
	violatorNode := bn
	details := []interface{}{1, 2}
	inc.GetNameAddressMock.Set(func() endpoints.Name { return "test" })
	inc.GetTransportKeyMock.Set(func() cryptkit.SignatureKeyHolder { return nil })
	inc.GetTransportCertMock.Set(func() cryptkit.CertificateHolder { return nil })
	be := bf.NewFraud(fraudType, msg, violatorHost, violatorNode, details...)
	require.Equal(t, fraudType, be.fraudType)

	require.Equal(t, msg, be.msg)

	require.Equal(t, violatorNode, be.violatorNode)

	require.Equal(t, details[1], be.details[1])

	require.True(t, be.captureMark != nil)

	bf = NewFraudFactory(nil)
	be = bf.NewFraud(fraudType, msg, nil, violatorNode, details...)

	require.True(t, be.captureMark == nil)
}

func TestNewNodeFraud(t *testing.T) {
	bf := NewFraudFactory(reportFunc)
	fraudType := 1
	msg := "test"
	bn := profiles.NewBaseNodeMock(t)
	violatorNode := bn
	details := []interface{}{1, 2}
	be := bf.NewNodeFraud(fraudType, msg, violatorNode, details...)
	require.Equal(t, msg, be.msg)
}

func TestNewHostFraud(t *testing.T) {
	ff := NewFraudFactory(reportFunc)
	fraudType := 1
	msg := "test"
	inc := endpoints.NewInboundMock(t)
	violatorHost := inc
	details := []interface{}{1, 2}
	inc.GetNameAddressMock.Set(func() endpoints.Name { return "test" })
	inc.GetTransportKeyMock.Set(func() cryptkit.SignatureKeyHolder { return nil })
	inc.GetTransportCertMock.Set(func() cryptkit.CertificateHolder { return nil })
	fe := ff.NewHostFraud(fraudType, msg, violatorHost, details...)
	require.Equal(t, msg, fe.msg)
}

func TestNewInconsistentMembershipAnnouncement(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewInconsistentMembershipAnnouncement(profiles.NewActiveNodeMock(t),
		profiles.MembershipAnnouncement{}, profiles.MembershipAnnouncement{})
	require.Equal(t, "multiple membership profile", fe.msg)
}

func TestNewMismatchedMembershipRank(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewMismatchedMembershipRank(profiles.NewActiveNodeMock(t),
		profiles.MembershipProfile{})
	require.Equal(t, "mismatched membership profile rank", fe.msg)
}

func TestNewMismatchedMembershipRankOrNodeCount(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewMismatchedMembershipRankOrNodeCount(profiles.NewActiveNodeMock(t),
		profiles.MembershipProfile{}, 0)
	require.Equal(t, "mismatched membership profile node count", fe.msg)
}

func TestNewUnknownNeighbour(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewUnknownNeighbour(profiles.NewBaseNodeMock(t))
	require.Equal(t, "unknown neighbour", fe.(FraudError).msg)
}

func TestNewMismatchedNeighbourRank(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewMismatchedNeighbourRank(profiles.NewBaseNodeMock(t))
	require.Equal(t, "mismatched neighbour rank", fe.(FraudError).msg)
}

func TestNewNeighbourMissingTarget(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewNeighbourMissingTarget(profiles.NewBaseNodeMock(t))
	require.Equal(t, "neighbour must include target node", fe.(FraudError).msg)
}

func TestNewNeighbourContainsSource(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewNeighbourContainsSource(profiles.NewBaseNodeMock(t))
	require.Equal(t, "neighbour must NOT include source node", fe.(FraudError).msg)
}

func TestNewInconsistentNeighbourAnnouncement(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewInconsistentNeighbourAnnouncement(profiles.NewBaseNodeMock(t))
	require.Equal(t, "multiple neighbour profile", fe.msg)
}

func TestNewInvalidPowerLevel(t *testing.T) {
	fe := NewFraudFactory(reportFunc).NewInvalidPowerLevel(profiles.NewBaseNodeMock(t))
	require.Equal(t, "power level is incorrect", fe.msg)
}
