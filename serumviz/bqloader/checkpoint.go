package bqloader

import (
	"context"
	"fmt"
	"time"

	"github.com/streamingfast/derr"
	pbserumhist "github.com/streamingfast/sf-solana/pb/sf/solana/serumhist/v1"
	"google.golang.org/api/iterator"
)

type ProcessFile struct {
	Table         string    `bigquery:"table"`
	Filename      string    `bigquery:"file"`
	StartBlockNum int64     `bigquery:"start_block_num"`
	StartBlockID  []byte    `bigquery:"start_block_id"`
	EndBlockNum   int64     `bigquery:"end_block_num"`
	EndBlockID    []byte    `bigquery:"end_block_id"`
	Time          time.Time `bigquery:"timestamp"`
}

func (bq *BQLoader) GetCheckpoint(ctx context.Context) (*pbserumhist.Checkpoint, error) {
	if len(bq.checkpoints) == 0 {
		err := bq.readCheckpoints(ctx)
		if err != nil {
			return nil, err
		}
	}

	var earliestCheckpoint *pbserumhist.Checkpoint
	for _, table := range []Table{tableOrders, tableFills, tableTraders} {
		tableCheckpoint, ok := bq.checkpoints[table]
		if !ok {
			continue
		}

		if tableCheckpoint == nil { // one or more checkpoints not set.  return nil.  caller will handle nil checkpoint
			return nil, nil
		}

		validCheckpoint, err := bq.validateCheckpoint(ctx, table.String(), tableCheckpoint.LastWrittenBlockNum)
		if err != nil {
			return nil, fmt.Errorf("could not validate checkpoint: %w", err)
		}

		if !validCheckpoint {
			return nil, fmt.Errorf("invalid checkpoint for table %s: data in table exists above checkpoint slot_num %d. this data needs to be erased from the table", table, tableCheckpoint.LastWrittenBlockNum)
		}

		if earliestCheckpoint == nil {
			earliestCheckpoint = tableCheckpoint
			continue
		}

		if tableCheckpoint.LastWrittenBlockNum < earliestCheckpoint.LastWrittenBlockNum {
			earliestCheckpoint = tableCheckpoint
		}
	}

	return earliestCheckpoint, nil
}

func (bq *BQLoader) readCheckpoints(ctx context.Context) error {
	for _, table := range []Table{tableOrders, tableFills, tableTraders} {
		tableCheckpoint, err := bq.readCheckpoint(ctx, table.String())
		if err != nil {
			return err
		}
		bq.checkpoints[table] = tableCheckpoint
	}
	return nil
}

func (bq *BQLoader) readCheckpoint(ctx context.Context, forTable string) (*pbserumhist.Checkpoint, error) {
	checkpointsTableName := fmt.Sprintf("%s.serum.%s", bq.dataset.ProjectID, tableProcessedFiles)

	var result *pbserumhist.Checkpoint
	queryFunc := func(ctx context.Context) error {
		queryString := fmt.Sprintf(`SELECT * FROM %s WHERE table="%s" ORDER BY timestamp DESC LIMIT 1`, checkpointsTableName, forTable)

		q := bq.client.Query(queryString)
		j, err := q.Run(ctx)
		if err != nil {
			return fmt.Errorf("could not run query `%s`: %w", queryString, err)
		}
		it, err := j.Read(ctx)
		if err != nil {
			return fmt.Errorf("could not read query results: %w", err)
		}

		for {
			var row ProcessFile
			err := it.Next(&row)
			if err == iterator.Done {
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not read row: %w", err)
			}

			result = &pbserumhist.Checkpoint{
				LastWrittenBlockNum: uint64(row.EndBlockNum),
				LastWrittenBlockId:  row.EndBlockID,
			}
		}
	}

	err := derr.RetryContext(ctx, 5, queryFunc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bq *BQLoader) validateCheckpoint(ctx context.Context, tableName string, slotNum uint64) (bool, error) {
	tableName = fmt.Sprintf("%s.serum.%s", bq.dataset.ProjectID, tableName)

	var valid bool
	queryFunc := func(ctx context.Context) error {
		queryString := fmt.Sprintf("SELECT slot_num FROM %s WHERE slot_num > %d", tableName, slotNum)
		q := bq.client.Query(queryString)
		j, err := q.Run(ctx)
		if err != nil {
			return fmt.Errorf("could not run query `%s`: %w", queryString, err)
		}
		it, err := j.Read(ctx)
		if err != nil {
			return fmt.Errorf("could not read query results: %w", err)
		}

		type Row struct {
			SlotNum int `bigquery:"slot_num"`
		}

		count := 0
		for {
			var row Row
			err := it.Next(&row)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return fmt.Errorf("could not read account trader row: %w", err)
			}
			count++
			break
		}
		valid = bool(count == 0)
		return nil
	}

	err := derr.RetryContext(ctx, 5, queryFunc)
	return valid, err
}
