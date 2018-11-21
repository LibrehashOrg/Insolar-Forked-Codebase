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

package nodenetwork

import (
	"sync"
	"testing"
	"time"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/consensus"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newActiveNode(ref byte) (core.RecordRef, []core.NodeRole, string, string) {
	return core.RecordRef{ref}, []core.NodeRole{core.RoleUnknown}, "127.0.0.1:12345", "1.1"
}

func testNode(ref core.RecordRef) *node {
	return &node{
		NodeID:    ref,
		NodeRoles: []core.NodeRole{core.RoleUnknown},
	}
}

func testNodeWithRole(ref core.RecordRef, role core.NodeRole) *node {
	return &node{
		NodeID:    ref,
		NodeRoles: []core.NodeRole{role},
	}
}

func newNodeKeeper() network.NodeKeeper {
	id := core.RecordRef{255}
	n := testNode(id)
	keeper := NewNodeKeeper(n)
	keeper.AddActiveNodes([]core.Node{testNode(id)})
	return keeper
}

func TestNodekeeper_GetOrigin(t *testing.T) {
	id := core.RecordRef{255}
	n := testNode(id)
	keeper := NewNodeKeeper(n)
	require.Equal(t, n, keeper.GetOrigin())
}

func TestNodekeeper_AddUnsync(t *testing.T) {
	id := core.RecordRef{}
	n := testNode(id)
	keeper := NewNodeKeeper(n)

	_, err := keeper.AddUnsync(newActiveNode(0))
	require.NoError(t, err)
	success, list := keeper.SetPulse(core.PulseNumber(0))
	require.True(t, success)
	require.Equal(t, 1, len(list.GetUnsync()))
}

func TestNodekeeper_AddUnsync2(t *testing.T) {
	keeper := newNodeKeeper()
	success, list := keeper.SetPulse(core.PulseNumber(0))
	_, err := keeper.AddUnsync(newActiveNode(0))
	require.NoError(t, err)
	require.True(t, success)
	require.Equal(t, 0, len(list.GetUnsync()))
}

func TestNodekeeper_AddUnsync3(t *testing.T) {
	keeper := newNodeKeeper()
	_, err := keeper.AddUnsync(newActiveNode(0))
	success, list := keeper.SetPulse(core.PulseNumber(0))
	_, err = keeper.AddUnsync(newActiveNode(1))
	require.NoError(t, err)
	require.True(t, success)
	require.Equal(t, 1, len(list.GetUnsync()))
}

func TestNodekeeper_pipeline(t *testing.T) {
	keeper := newNodeKeeper()
	for i := 0; i < 4; i++ {
		_, err := keeper.AddUnsync(newActiveNode(byte(2 * i)))
		require.NoError(t, err)
		pulse := core.PulseNumber(i)
		success, list := keeper.SetPulse(pulse)
		require.True(t, success)
		_, err = keeper.AddUnsync(newActiveNode(byte(2*i + 1)))
		require.NoError(t, err)
		keeper.Sync(list.GetUnsync(), pulse)
	}
	// 3 nodes should not advance to join active list
	// 5 nodes should advance + 1 origin node
	require.Equal(t, 6, len(keeper.GetActiveNodes()))
	for i := 0; i < 5; i++ {
		require.NotNil(t, keeper.GetActiveNode(core.RecordRef{byte(i)}))
	}
}

func TestNodekeeper_doubleSync(t *testing.T) {
	keeper := newNodeKeeper()
	_, err := keeper.AddUnsync(newActiveNode(0))
	require.NoError(t, err)
	pulse := core.PulseNumber(0)
	success, list := keeper.SetPulse(pulse)
	require.True(t, success)
	require.Equal(t, 1, len(list.GetUnsync()))
	keeper.Sync(list.GetUnsync(), pulse)
	// second sync should be ignored because pulse has not changed
	keeper.Sync(list.GetUnsync(), pulse)
	// and added unsync node should not advance to active list (only one origin node would be in the list)
	require.Equal(t, 1, len(keeper.GetActiveNodes()))
	require.Equal(t, keeper.GetOrigin().ID(), keeper.GetActiveNodes()[0].ID())
}

func TestNodekeeper_doubleSetPulse(t *testing.T) {
	keeper := newNodeKeeper()
	_, err := keeper.AddUnsync(newActiveNode(0))
	require.NoError(t, err)
	pulse := core.PulseNumber(0)
	_, list := keeper.SetPulse(pulse)
	keeper.Sync(list.GetUnsync(), pulse)
	_, _ = keeper.SetPulse(core.PulseNumber(1))
	_, _ = keeper.SetPulse(core.PulseNumber(2))
	// node with ref 0 advanced to active list
	require.Equal(t, 2, len(keeper.GetActiveNodes()))
	require.NotNil(t, keeper.GetActiveNode(core.RecordRef{0}))
}

func TestNodekeeper_outdatedSync(t *testing.T) {
	keeper := newNodeKeeper()
	num := 4
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		time.Sleep(100 * time.Millisecond)
		go func(k network.NodeKeeper, i int) {
			_, _ = k.AddUnsync(newActiveNode(byte(2 * i)))
			_, _ = k.AddUnsync(newActiveNode(byte(2*i + 1)))
			pulse := core.PulseNumber(i)
			success, list := k.SetPulse(pulse)
			require.True(t, success)
			// imitate long consensus process
			time.Sleep(200 * time.Millisecond)
			k.Sync(list.GetUnsync(), pulse)
			wg.Done()
		}(keeper, i)
	}
	wg.Wait()
	// All Syncs calls are executed out of date
	// So, no nodes should advance to active list (we should have only 1 origin node in active)
	require.Equal(t, 1, len(keeper.GetActiveNodes()))
}

func TestNodekeeper_SetPulse(t *testing.T) {
	keeper := newNodeKeeper()
	success, _ := keeper.SetPulse(core.PulseNumber(10))
	require.True(t, success)
	// Pulses should pass in ascending order
	success, _ = keeper.SetPulse(core.PulseNumber(9))
	require.False(t, success)
}

func TestNodekeeper_GetUnsyncHolder(t *testing.T) {
	keeper := newNodeKeeper()
	pulse := core.PulseNumber(10)
	requestedPulseLess := core.PulseNumber(11)
	nextPulse := core.PulseNumber(12)
	requestedPulseGreater := core.PulseNumber(13)
	success, _ := keeper.SetPulse(pulse)
	require.True(t, success)

	wg := sync.WaitGroup{}
	waitersNext := 10
	waitersRequestedLess := 5
	waitersRequestedGreater := 5
	wg.Add(waitersNext + waitersRequestedLess + waitersRequestedGreater)

	f := func(t *testing.T, requestedPulse core.PulseNumber, nextPulse core.PulseNumber, wg *sync.WaitGroup) {
		holder, err := keeper.GetUnsyncHolder(requestedPulse, 10*time.Millisecond)
		if requestedPulse == nextPulse {
			require.NoError(t, err)
			require.NotNil(t, holder)
			require.Equal(t, nextPulse, holder.GetPulse())
		} else {
			require.Error(t, err)
			require.Nil(t, holder)
		}
		wg.Done()
	}

	for i := 0; i < waitersNext; i++ {
		go f(t, nextPulse, nextPulse, &wg)
	}

	for i := 0; i < waitersRequestedLess; i++ {
		go f(t, requestedPulseLess, nextPulse, &wg)
	}

	for i := 0; i < waitersRequestedGreater; i++ {
		go f(t, requestedPulseGreater, nextPulse, &wg)
	}

	time.Sleep(time.Millisecond)
	success, _ = keeper.SetPulse(nextPulse)
	require.True(t, success)
	wg.Wait()
}

func TestNodekeeper_GetUnsyncHolder2(t *testing.T) {
	keeper := newNodeKeeper()
	prevPulse := core.PulseNumber(9)
	pulse := core.PulseNumber(10)
	success, _ := keeper.SetPulse(pulse)
	require.True(t, success)
	holder, err := keeper.GetUnsyncHolder(pulse, 0)
	require.NoError(t, err)
	require.NotNil(t, holder)
	require.Equal(t, pulse, holder.GetPulse())

	holder, err = keeper.GetUnsyncHolder(prevPulse, 0)
	require.Error(t, err)
	require.Nil(t, holder)
}

func TestNodekeeper_GetUnsyncHolder3(t *testing.T) {
	keeper := newNodeKeeper()
	pulse := core.PulseNumber(10)
	nextPulse := core.PulseNumber(11)
	success, _ := keeper.SetPulse(pulse)
	require.True(t, success)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(keeper network.NodeKeeper, wg *sync.WaitGroup) {
		_, err := keeper.GetUnsyncHolder(nextPulse, time.Millisecond)
		require.Error(t, err)
		wg.Done()
	}(keeper, &wg)
	time.Sleep(10 * time.Millisecond)
	wg.Wait()
}

func TestNodeKeeper_notifyAddUnsync(t *testing.T) {
	keeper := newNodeKeeper()

	nodePassesConsensus := func(ref core.RecordRef) bool {
		return ref[0] >= 5
	}

	refsCount := 10
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < refsCount; i++ {
		ref := core.RecordRef{byte(i)}
		ch, err := keeper.AddUnsync(newActiveNode(byte(i)))
		require.NoError(t, err)

		go func(t *testing.T, ch chan core.Node, ref core.RecordRef, wg *sync.WaitGroup) {
			node := <-ch
			if nodePassesConsensus(ref) {
				require.NotNil(t, node)
				require.Equal(t, ref, node.ID())
			} else {
				require.Nil(t, node)
			}
			wg.Done()
		}(t, ch, ref, &wg)
	}

	success, list := keeper.SetPulse(core.PulseNumber(133))
	require.True(t, success)
	require.NotNil(t, list)
	require.Equal(t, refsCount, len(list.GetUnsync()))

	syncCandidates := make([]core.Node, 0)
	for _, node := range list.GetUnsync() {
		if nodePassesConsensus(node.ID()) {
			syncCandidates = append(syncCandidates, node)
		}
	}
	keeper.Sync(syncCandidates, core.PulseNumber(133))
	wg.Wait()
}

func TestUnsyncList_GetUnsync(t *testing.T) {
	unsyncNodes := []core.Node{}
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), unsyncNodes)
	require.Empty(t, unsyncList.GetUnsync())
	require.Equal(t, core.PulseNumber(10), unsyncList.GetPulse())
}

func TestUnsyncList_GetHash(t *testing.T) {
	unsyncNodes := []core.Node{}
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), unsyncNodes)
	hash := []byte{'a', 'b', 'c'}
	h := make([]*network.NodeUnsyncHash, 0)
	h = append(h, &network.NodeUnsyncHash{core.RecordRef{1}, hash})
	unsyncList.SetHash(h)
	h2, err := unsyncList.GetHash(0)
	require.NoError(t, err)
	require.Equal(t, hash, h2[0].Hash)
}

func TestUnsyncList_GetHash2(t *testing.T) {
	unsyncNodes := []core.Node{}
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), unsyncNodes)
	hash := []byte{'a', 'b', 'c'}
	h := make([]*network.NodeUnsyncHash, 0)
	h = append(h, &network.NodeUnsyncHash{core.RecordRef{1}, hash})

	wg := sync.WaitGroup{}
	waiters := 10
	wg.Add(waiters)

	for i := 0; i < waiters; i++ {
		go func(list consensus.UnsyncHolder) {
			h, err := list.GetHash(time.Millisecond * 10)
			require.NoError(t, err)
			require.NotNil(t, h)
			require.Equal(t, hash, h[0].Hash)
			wg.Done()
		}(unsyncList)
	}
	time.Sleep(time.Millisecond)
	unsyncList.SetHash(h)
	wg.Wait()
}

func TestUnsyncList_GetHash3(t *testing.T) {
	unsyncNodes := []core.Node{}
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), unsyncNodes)
	hash := []byte{'a', 'b', 'c'}
	h := make([]*network.NodeUnsyncHash, 0)
	h = append(h, &network.NodeUnsyncHash{core.RecordRef{1}, hash})

	wg := sync.WaitGroup{}
	waiters := 10
	wg.Add(waiters)

	for i := 0; i < waiters; i++ {
		go func(list consensus.UnsyncHolder) {
			h, err := list.GetHash(time.Millisecond * 1)
			require.Error(t, err)
			require.Nil(t, h)
			wg.Done()
		}(unsyncList)
	}
	time.Sleep(time.Millisecond * 10)
	unsyncList.SetHash(h)
	wg.Wait()
}

func TestUnsyncList_AddUnsyncList(t *testing.T) {
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), nil)
	unsyncList.AddUnsyncList(core.RecordRef{1}, []core.Node{})
	_, exists := unsyncList.GetUnsyncList(core.RecordRef{1})
	require.True(t, exists)
}

func TestUnsyncList_AddUnsyncHash(t *testing.T) {
	unsyncList := NewUnsyncHolder(core.PulseNumber(10), nil)
	unsyncList.AddUnsyncHash(core.RecordRef{1}, []*network.NodeUnsyncHash{})
	_, exists := unsyncList.GetUnsyncHash(core.RecordRef{1})
	require.True(t, exists)
}

func TestNodekeeper_GetActiveNodesByRole(t *testing.T) {
	keeper := newNodeKeeper()
	node1 := testNode(testutils.RandomRef())
	node1.NodeRoles = []core.NodeRole{core.RoleVirtual}
	node2 := testNode(testutils.RandomRef())
	node2.NodeRoles = []core.NodeRole{core.RoleLightMaterial}
	keeper.AddActiveNodes([]core.Node{node1, node2})

	require.Equal(t, node1.NodeID, keeper.GetActiveNodesByRole(core.RoleVirtualExecutor)[0])
	require.Equal(t, node1.NodeID, keeper.GetActiveNodesByRole(core.RoleVirtualValidator)[0])
	require.Equal(t, node2.NodeID, keeper.GetActiveNodesByRole(core.RoleLightValidator)[0])
	require.Equal(t, node2.NodeID, keeper.GetActiveNodesByRole(core.RoleLightExecutor)[0])
	require.Empty(t, keeper.GetActiveNodesByRole(core.RoleHeavyExecutor))
}

func TestNodekeeper_GetActiveNodeByShortID(t *testing.T) {
	keeper := newNodeKeeper()
	node1 := testNode(testutils.RandomRef())
	keeper.AddActiveNodes([]core.Node{node1})
	require.NotNil(t, keeper.GetActiveNodeByShortID(node1.ShortID()))
	require.Nil(t, keeper.GetActiveNodeByShortID(node1.ShortID()+1))
}

func TestNodekeeper_AddActiveNodes(t *testing.T) {
	keeper := newNodeKeeper()
	node := testNodeWithRole(testutils.RandomRef(), core.RoleVirtual)
	keeper.AddActiveNodes([]core.Node{node, node, node})
	list := keeper.GetActiveNodesByRole(core.RoleVirtualValidator)
	assert.Equal(t, 1, len(list))
}
