package store

import (
	"iot/data_simulator/common"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type TemperatureStore struct {
	ch *clickhouse.Conn
}

func (s *TemperatureStore) InsertBatch(data []common.Metrics) error {
	return nil
}
