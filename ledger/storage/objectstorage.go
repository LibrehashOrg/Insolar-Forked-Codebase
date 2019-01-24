package storage

import (
	"context"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/ledger/storage/index"
	"github.com/insolar/insolar/ledger/storage/record"
)

// ObjectStorage returns objects and their meta
type ObjectStorage interface {
	GetBlob(ctx context.Context, jetID core.RecordID, id *core.RecordID) ([]byte, error)
	SetBlob(ctx context.Context, jetID core.RecordID, pulseNumber core.PulseNumber, blob []byte) (*core.RecordID, error)

	GetRecord(ctx context.Context, jetID core.RecordID, id *core.RecordID) (record.Record, error)
	SetRecord(ctx context.Context, jetID core.RecordID, pulseNumber core.PulseNumber, rec record.Record) (*core.RecordID, error)

	GetObjectIndex(
		ctx context.Context,
		jetID core.RecordID,
		id *core.RecordID,
		forupdate bool,
	) (*index.ObjectLifeline, error)

	SetObjectIndex(
		ctx context.Context,
		jetID core.RecordID,
		id *core.RecordID,
		idx *index.ObjectLifeline,
	) error

	RemoveObjectIndex(
		ctx context.Context,
		jetID core.RecordID,
		ref *core.RecordID,
	) error
}
