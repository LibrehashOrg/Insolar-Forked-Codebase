//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package smachine

import (
	"sync/atomic"

	"github.com/insolar/insolar/network/consensus/common/rwlock"
)

func NewInfiniteLock(name string) SyncLink {
	return NewSyncLinkNoLock(&infiniteLock{name: name})
}

type infiniteLock struct {
	name  string
	count int32 //atomic
}

func (p *infiniteLock) CheckState() Decision {
	return NotPassed
}

func (p *infiniteLock) CheckDependency(dep SlotDependency) Decision {
	if entry, ok := dep.(*infiniteLockEntry); ok && entry.ctl == p {
		return NotPassed
	}
	return Impossible
}

func (p *infiniteLock) UseDependency(dep SlotDependency, flags SlotDependencyFlags) Decision {
	if entry, ok := dep.(*infiniteLockEntry); ok {
		switch {
		case !entry.IsCompatibleWith(flags):
			return Impossible
		case entry.ctl == p:
			return NotPassed
		}
	}
	return Impossible
}

func (p *infiniteLock) CreateDependency(slot *Slot, flags SlotDependencyFlags, syncer rwlock.RWLocker) (BoolDecision, SlotDependency) {
	atomic.AddInt32(&p.count, 1)
	return false, &infiniteLockEntry{p, flags}
}

func (p *infiniteLock) GetLimit() (limit int, isAdjustable bool) {
	return 0, false
}

func (p *infiniteLock) AdjustLimit(limit int) ([]StepLink, bool) {
	panic("illegal state")
}

func (p *infiniteLock) GetCounts() (active, inactive int) {
	return 0, int(p.count)
}

func (p *infiniteLock) GetName() string {
	return p.name
}

var _ SlotDependency = &infiniteLockEntry{}

type infiniteLockEntry struct {
	ctl       *infiniteLock
	slotFlags SlotDependencyFlags
}

func (v infiniteLockEntry) IsReleaseOnStepping() bool {
	return v.slotFlags&syncForOneStep != 0
}

func (infiniteLockEntry) IsReleaseOnWorking() bool {
	return false
}

func (v infiniteLockEntry) Release() []StepLink {
	atomic.AddInt32(&v.ctl.count, -1)
	return nil
}

func (v infiniteLockEntry) IsCompatibleWith(flags SlotDependencyFlags) bool {
	f := v.slotFlags
	return f&flags == flags
}