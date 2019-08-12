/*
 *    Copyright 2019 Insolar Technologies
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

package pulsar

import (
	"context"
	"testing"
	"time"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/pulsar/entropygenerator"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestPulsar_Send(t *testing.T) {
	t.Skip()
	distMock := testutils.NewPulseDistributorMock(t)
	pn := insolar.PulseNumber(123)

	distMock.DistributeMock.Set(func(ctx context.Context, p1 insolar.Pulse) {
		require.Equal(t, pn, p1.PulseNumber)
	})

	pcs := platformpolicy.NewPlatformCryptographyScheme()
	crypto := testutils.NewCryptographyServiceMock(t)
	crypto.SignMock.Return(&insolar.Signature{}, nil)

	p := NewPulsar(
		configuration.NewPulsar(),
		crypto,
		pcs,
		platformpolicy.NewKeyProcessor(),
		distMock,
		&entropygenerator.StandardEntropyGenerator{},
	)

	err := p.Send(context.TODO(), pn)

	require.NoError(t, err)
	require.Equal(t, pn, p.LastPN())

	distMock.MinimockWait(1 * time.Minute)
}
