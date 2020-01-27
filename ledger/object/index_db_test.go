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

// +build slowtest

package object

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/insolar/store"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/heavy/migration"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/tests/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var pool *pgxpool.Pool

// TestMain does the before and after setup
func TestMain(m *testing.M) {
	ctx := context.Background()
	log.Info("[TestMain] About to start PostgreSQL...")
	pgURL, stopPostgreSQL := common.StartPostgreSQL()
	log.Info("[TestMain] PostgreSQL started!")

	p, err := pgxpool.Connect(ctx, pgURL)
	if err != nil {
		stopPostgreSQL()
		log.Panicf("[TestMain] pgxpool.Connect() failed: %v", err)
	}
	pool = p

	migrationPath := "../../migration"
	cwd, err := os.Getwd()
	if err != nil {
		stopPostgreSQL()
		panic(errors.Wrap(err, "[TestMain] os.Getwd failed"))
	}
	log.Infof("[TestMain] About to run PostgreSQL migration, cwd = %s, migration migrationPath = %s", cwd, migrationPath)
	ver, err := migration.MigrateDatabase(ctx, pool, migrationPath)
	if err != nil {
		stopPostgreSQL()
		panic(errors.Wrap(err, "Unable to migrate database"))
	}
	log.Infof("[TestMain] PostgreSQL database migration done, current schema version: %d", ver)

	// Run all tests
	code := m.Run()

	log.Info("[TestMain] Cleaning up...")
	stopPostgreSQL()
	os.Exit(code)
}

func truncateIndexTables() {
	_, err := pool.Exec(context.Background(), "TRUNCATE last_known_pulse_for_indexes, indexes")
	if err != nil {
		panic(err)
	}
}

// AALEKSEEV TODO rewrite

func BadgerDefaultOptions(dir string) badger.Options {
	ops := badger.DefaultOptions(dir)
	ops.CompactL0OnClose = false
	ops.SyncWrites = false

	return ops
}

const indexCount = 5

func TestIndexDB_DontLooseIndexAfterTruncate(t *testing.T) {
	defer truncateIndexTables()

	ctx := inslogger.TestContext(t)
	tmpdir, err := ioutil.TempDir("", "bdb-test-")
	defer os.RemoveAll(tmpdir)
	assert.NoError(t, err)

	ops := BadgerDefaultOptions(tmpdir)
	dbMock, err := store.NewBadgerDB(ops)
	require.NoError(t, err)
	defer dbMock.Stop(ctx)
	require.NoError(t, err)

	indexStore := NewIndexDB(pool, nil)

	testPulse := insolar.GenesisPulse.PulseNumber
	nextPulse := testPulse + 1
	bucket := record.Index{}
	bucket.ObjID = gen.ID()

	err = indexStore.SetIndex(ctx, testPulse, bucket)
	require.NoError(t, err)
	_, err = indexStore.ForID(ctx, testPulse, bucket.ObjID)
	require.NoError(t, err)

	err = indexStore.SetIndex(ctx, nextPulse, bucket)
	require.NoError(t, err)

	_, err = indexStore.ForID(ctx, nextPulse, bucket.ObjID)
	require.NoError(t, err)

	err = indexStore.TruncateHead(ctx, nextPulse)
	require.NoError(t, err)

	_, err = indexStore.ForID(ctx, nextPulse, bucket.ObjID)
	require.EqualError(t, err, ErrIndexNotFound.Error())

	// no update such object in that pulse -> try to get last known pulse but it refers to nextPulse
	// , but we Truncate index with that pulse -> couldn't find that object
	_, err = indexStore.ForID(ctx, nextPulse+1, bucket.ObjID)
	require.EqualError(t, err, ErrIndexNotFound.Error())

	err = indexStore.UpdateLastKnownPulse(ctx, testPulse)
	require.NoError(t, err)
	_, err = indexStore.ForID(ctx, testPulse+2, bucket.ObjID)
	require.NoError(t, err)
}

func TestIndexDB_TruncateHead(t *testing.T) {
	defer truncateIndexTables()

	ctx := inslogger.TestContext(t)
	tmpdir, err := ioutil.TempDir("", "bdb-test-")
	defer os.RemoveAll(tmpdir)
	assert.NoError(t, err)

	ops := BadgerDefaultOptions(tmpdir)
	dbMock, err := store.NewBadgerDB(ops)
	require.NoError(t, err)
	defer dbMock.Stop(ctx)
	require.NoError(t, err)

	indexStore := NewIndexDB(pool, NewRecordDB(dbMock))

	numElements := 10

	// it's used for writing pulses in random order to db
	indexes := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		indexes[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })

	startPulseNumber := insolar.GenesisPulse.PulseNumber
	objects := make([]insolar.ID, numElements)
	for _, idx := range indexes {
		pulse := startPulseNumber + insolar.PulseNumber(idx)
		objects[idx] = gen.ID()

		bucket := record.Index{}

		bucket.ObjID = objects[idx]
		err := indexStore.SetIndex(ctx, pulse, bucket)
		require.NoError(t, err)

		for i := 0; i < indexCount; i++ {
			bucket := record.Index{}

			bucket.ObjID = gen.ID()
			err := indexStore.SetIndex(ctx, pulse, bucket)
			require.NoError(t, err)
		}

	}

	for i := 0; i < numElements; i++ {
		_, err := indexStore.ForID(ctx, startPulseNumber+insolar.PulseNumber(i), objects[i])
		require.NoError(t, err)
	}

	numLeftElements := numElements / 2
	err = indexStore.TruncateHead(ctx, startPulseNumber+insolar.PulseNumber(numLeftElements))
	require.NoError(t, err)

	for i := 0; i < numLeftElements; i++ {
		_, err := indexStore.ForID(ctx, startPulseNumber+insolar.PulseNumber(i), objects[i])
		require.NoError(t, err)
	}

	for i := numElements - 1; i >= numLeftElements; i-- {
		_, err := indexStore.ForID(ctx, startPulseNumber+insolar.PulseNumber(i), objects[i])
		require.EqualError(t, err, ErrIndexNotFound.Error())
	}
}

func TestDBIndexStorage_ForID(t *testing.T) {
	defer truncateIndexTables()

	ctx := inslogger.TestContext(t)

	id := gen.ID()

	t.Run("returns error when no index-value for id", func(t *testing.T) {
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		storage := NewIndexDB(pool, NewRecordDB(db))
		pn := gen.PulseNumber()

		_, err = storage.ForID(ctx, pn, id)

		assert.Equal(t, ErrIndexNotFound, err)
	})
}

func TestDBIndexStorage_ForPulse(t *testing.T) {
	defer truncateIndexTables()

	ctx := inslogger.TestContext(t)
	prevPn := gen.PulseNumber()
	pn := prevPn + 10
	nextPn := pn + 20

	// Sort indexes for proper compare
	// For now badger iterator already sorted by key but this behavior can change
	sortIndexes := func(slice []record.Index) {
		cmp := func(i, j int) bool {
			cmp := bytes.Compare(slice[i].ObjID.Bytes(), slice[j].ObjID.Bytes())
			return cmp < 0
		}
		sort.Slice(slice, cmp)
	}

	t.Run("empty index storage", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		storage := NewIndexDB(pool, nil)

		indexes, err := storage.ForPulse(ctx, pn)
		require.Error(t, err)
		require.Equal(t, err, ErrIndexNotFound)
		require.Nil(t, indexes)
	})

	t.Run("index storage with couple values", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		storage := NewIndexDB(pool, nil)

		var indexes []record.Index
		for i := 0; i < indexCount; i++ {
			indexes = append(indexes, record.Index{ObjID: gen.ID()})
			err = storage.SetIndex(ctx, pn, indexes[i])
			require.NoError(t, err)
		}

		realIndexes, err := storage.ForPulse(ctx, pn)
		require.NoError(t, err)
		require.Equal(t, len(indexes), len(realIndexes))

		sortIndexes(realIndexes)
		sortIndexes(indexes)
		for i := 0; i < indexCount; i++ {
			require.Equal(t, indexes[i], realIndexes[i])
		}
	})

	t.Run("index storage with couple values in different pulses", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		storage := NewIndexDB(pool, nil)

		var indexes []record.Index
		for i := 0; i < indexCount; i++ {
			indexes = append(indexes, record.Index{ObjID: gen.ID()})
			err = storage.SetIndex(ctx, pn, indexes[i])
			require.NoError(t, err)
		}

		// add some values in prev pulse
		for i := 0; i < indexCount; i++ {
			err = storage.SetIndex(ctx, prevPn, record.Index{ObjID: gen.ID()})
			require.NoError(t, err)
		}

		// add some values in next pulse
		for i := 0; i < indexCount; i++ {
			err = storage.SetIndex(ctx, nextPn, record.Index{ObjID: gen.ID()})
			require.NoError(t, err)
		}

		realIndexes, err := storage.ForPulse(ctx, pn)
		require.NoError(t, err)
		require.Equal(t, len(indexes), len(realIndexes))

		sortIndexes(realIndexes)
		sortIndexes(indexes)
		for i := 0; i < indexCount; i++ {
			require.Equal(t, indexes[i], realIndexes[i])
		}
	})
}

func TestDBIndex_SetBucket(t *testing.T) {
	ctx := inslogger.TestContext(t)
	objID := gen.ID()
	lflID := gen.ID()
	buck := record.Index{
		ObjID: objID,
		Lifeline: record.Lifeline{
			LatestState: &lflID,
		},
	}

	t.Run("saves correct bucket", func(t *testing.T) {
		defer truncateIndexTables()

		pn := gen.PulseNumber()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		ops := BadgerDefaultOptions(tmpdir)
		db, err := store.NewBadgerDB(ops)
		require.NoError(t, err)
		defer db.Stop(context.Background())

		index := NewIndexDB(pool, NewRecordDB(db))

		err = index.SetIndex(ctx, pn, buck)
		require.NoError(t, err)

		res, err := index.ForID(ctx, pn, objID)
		require.NoError(t, err)

		idxBuf, _ := buck.Marshal()
		resBuf, _ := res.Marshal()

		assert.Equal(t, idxBuf, resBuf)
	})

	t.Run("re-save works fine", func(t *testing.T) {
		defer truncateIndexTables()

		pn := gen.PulseNumber()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		index := NewIndexDB(pool, NewRecordDB(db))

		err = index.SetIndex(ctx, pn, buck)
		require.NoError(t, err)

		sLlflID := gen.ID()
		sBuck := record.Index{
			ObjID: objID,
			Lifeline: record.Lifeline{
				LatestState: &sLlflID,
			},
		}

		err = index.SetIndex(ctx, pn, sBuck)
		require.NoError(t, err)

		res, err := index.ForID(ctx, pn, objID)
		require.NoError(t, err)

		idxBuf, _ := sBuck.Marshal()
		resBuf, _ := res.Marshal()

		assert.Equal(t, idxBuf, resBuf)
	})
}

func TestIndexDB_FetchFilament(t *testing.T) {
	defer truncateIndexTables()

	ctx := inslogger.TestContext(t)
	tmpdir, err := ioutil.TempDir("", "bdb-test-")
	defer os.RemoveAll(tmpdir)
	require.NoError(t, err)

	ops := BadgerDefaultOptions(tmpdir)
	db, err := store.NewBadgerDB(ops)
	require.NoError(t, err)
	defer db.Stop(context.Background())
	recordStorage := NewRecordDB(db)
	index := NewIndexDB(pool, NewRecordDB(db))

	first := insolar.NewID(1, nil)
	second := insolar.NewID(2, nil)

	firstMeta := *insolar.NewID(11, nil)
	secondMeta := *insolar.NewID(22, nil)

	firstFil := record.PendingFilament{
		RecordID: *first,
	}
	firstFilV := record.Wrap(&firstFil)
	secondFil := record.PendingFilament{
		RecordID:       *second,
		PreviousRecord: first,
	}
	secondFilV := record.Wrap(&secondFil)

	_ = recordStorage.Set(ctx, record.Material{ID: *first})
	_ = recordStorage.Set(ctx, record.Material{ID: *second})
	_ = recordStorage.Set(ctx, record.Material{Virtual: firstFilV, ID: firstMeta})
	_ = recordStorage.Set(ctx, record.Material{Virtual: secondFilV, ID: secondMeta})

	fi := &record.Index{
		PendingRecords: []insolar.ID{firstMeta, secondMeta},
	}

	res, err := index.filament(fi)

	require.NoError(t, err)
	require.Equal(t, 2, len(res))

	require.Equal(t, *first, res[0].RecordID)
	require.Equal(t, firstMeta, res[0].MetaID)

	require.Equal(t, *second, res[1].RecordID)
	require.Equal(t, secondMeta, res[1].MetaID)
}

func TestIndexDB_NextFilament(t *testing.T) {
	defer truncateIndexTables()
	ctx := inslogger.TestContext(t)

	first := insolar.NewID(1, nil)
	firstMeta := *insolar.NewID(11, nil)

	t.Run("previous exists", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		recordStorage := NewRecordDB(db)
		index := NewIndexDB(pool, NewRecordDB(db))

		firstFil := record.PendingFilament{
			PreviousRecord: first,
		}
		firstFilV := record.Wrap(&firstFil)

		_ = recordStorage.Set(ctx, record.Material{Virtual: firstFilV, ID: firstMeta})

		fi := &record.Index{
			PendingRecords: []insolar.ID{firstMeta},
		}

		cc, npn, err := index.nextFilament(fi)

		require.NoError(t, err)
		require.Equal(t, true, cc)

		require.Equal(t, insolar.PulseNumber(1), npn)
	})

	t.Run("previous doesn't exist", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		recordStorage := NewRecordDB(db)
		index := NewIndexDB(pool, NewRecordDB(db))

		firstFil := record.PendingFilament{}
		firstFilV := record.Wrap(&firstFil)

		_ = recordStorage.Set(ctx, record.Material{Virtual: firstFilV, ID: firstMeta})

		fi := &record.Index{
			PendingRecords: []insolar.ID{firstMeta},
		}

		cc, _, err := index.nextFilament(fi)

		require.NoError(t, err)
		require.Equal(t, false, cc)
	})

	t.Run("doesn't exist", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		index := NewIndexDB(pool, NewRecordDB(db))

		fi := &record.Index{
			PendingRecords: []insolar.ID{firstMeta},
		}

		cc, _, err := index.nextFilament(fi)

		require.Error(t, err, store.ErrNotFound)
		require.Equal(t, false, cc)
	})
}

func TestIndexDB_Records(t *testing.T) {
	defer truncateIndexTables()
	ctx := inslogger.TestContext(t)

	t.Run("returns err, if readUntil > readFrom", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		index := NewIndexDB(pool, NewRecordDB(db))

		res, err := index.Records(ctx, 1, 10, insolar.ID{})

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("works fine", func(t *testing.T) {
		defer truncateIndexTables()
		tmpdir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(tmpdir)
		require.NoError(t, err)

		db, err := store.NewBadgerDB(BadgerDefaultOptions(tmpdir))
		require.NoError(t, err)
		defer db.Stop(context.Background())
		index := NewIndexDB(pool, NewRecordDB(db))
		rms := NewRecordDB(db)

		pn := insolar.PulseNumber(3)
		pnS := insolar.PulseNumber(2)
		pnT := insolar.PulseNumber(1)

		// Records
		idT := insolar.NewID(pnT, nil)
		rT := record.IncomingRequest{Object: insolar.NewReference(gen.ID())}
		rTV := record.Wrap(&rT)
		_ = rms.Set(ctx, record.Material{Virtual: rTV, ID: *idT})

		idS := insolar.NewID(pnS, nil)
		rS := record.IncomingRequest{Object: insolar.NewReference(gen.ID())}
		rSV := record.Wrap(&rS)
		_ = rms.Set(ctx, record.Material{Virtual: rSV, ID: *idS})

		id := insolar.NewID(pn, nil)
		r := record.IncomingRequest{Object: insolar.NewReference(gen.ID())}
		rv := record.Wrap(&r)
		_ = rms.Set(ctx, record.Material{Virtual: rv, ID: *id})

		// Pending filaments
		midT := insolar.NewID(pnT, []byte{1})
		mT := record.PendingFilament{RecordID: *idT}
		mTV := record.Wrap(&mT)
		_ = rms.Set(ctx, record.Material{Virtual: mTV, ID: *midT})

		midS := insolar.NewID(pnS, []byte{1})
		mS := record.PendingFilament{RecordID: *idS, PreviousRecord: midT}
		mSV := record.Wrap(&mS)
		_ = rms.Set(ctx, record.Material{Virtual: mSV, ID: *midS})

		mid := insolar.NewID(pn, []byte{1})
		m := record.PendingFilament{RecordID: *id, PreviousRecord: midS}
		mV := record.Wrap(&m)
		_ = rms.Set(ctx, record.Material{Virtual: mV, ID: *mid})

		objID := gen.ID()

		third := record.Index{ObjID: objID, PendingRecords: []insolar.ID{*midT}}
		second := record.Index{ObjID: objID, PendingRecords: []insolar.ID{*midS}}
		first := record.Index{ObjID: objID, PendingRecords: []insolar.ID{*mid}}

		err = index.SetIndex(ctx, pn, first)
		require.NoError(t, err)
		err = index.SetIndex(ctx, pnS, second)
		require.NoError(t, err)
		err = index.SetIndex(ctx, pnT, third)
		require.NoError(t, err)

		res, err := index.Records(ctx, insolar.PulseNumber(3), insolar.PulseNumber(2), objID)

		require.NoError(t, err)
		require.Equal(t, 2, len(res))

		require.Equal(t, *idS, res[0].RecordID)
		require.Equal(t, *id, res[1].RecordID)

		require.Equal(t, *midS, res[0].MetaID)
		require.Equal(t, *mid, res[1].MetaID)
	})

}
