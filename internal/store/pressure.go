package store

import (
	"iot/data_simulator/common"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type PressureStore struct {
	ch *clickhouse.Conn
}

func (s *PressureStore) InsertBatch(data []common.Metrics) error {
	return nil
}
