package store

import (
	"context"
	"iot/data_simulator/common"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type TemperatureStore struct {
	ch *clickhouse.Conn
}

func (s *TemperatureStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO temperature (id, device_name, device_id, baseline_temperature, spike_probability, spike_magnitude, noise_level, updated_interval, drift_rate, current_temperature, is_spiking, last_spike_time, next_read_time, trend)")
	if err != nil {
		log.Printf("Error preparing batch: %v", err)
		return err
	}
	for _, i := range data {
		updateTime := time.Now().Add(i.UpdateInterval)

		if err := batch.Append(
			i.Id,
			i.DeviceName,
			i.DeviceId,
			i.BaselineTemp,
			i.SpikeProbability,
			i.SpikeMagnitude,
			i.NoiseLevel,
			updateTime,
			i.DriftRate,
			i.CurrentTemp,
			i.IsSpiking,
			i.LastSpikeTime,
			i.NextRead,
			i.Trend,
		); err != nil {
			log.Printf("Error appending to batch: %v", err)
			return err
		}
	}
	return batch.Send()
}
