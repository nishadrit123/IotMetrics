package store

import (
	"iot/data_simulator/common"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type HumidityStore struct {
	ch *clickhouse.Conn
}

func (s *HumidityStore) InsertBatch(data []common.Metrics) error {
	return nil
}
