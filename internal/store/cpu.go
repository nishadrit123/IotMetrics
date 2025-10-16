package store

import (
	"context"
	"iot/data_simulator/common"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type CPUStore struct {
	ch *clickhouse.Conn
}

func (s *CPUStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO cpu (id, device_name, device_id, baseline_usage, spike_probability, spike_magnitude, noise_level, updated_interval, current_usage, cpu_temperature, is_spiking, last_spike_time, next_read_time)")
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
			i.BaselineUsage,
			i.SpikeProbability,
			i.SpikeMagnitude,
			i.NoiseLevel,
			updateTime,
			i.CurrentUsage,
			i.Temperature,
			i.IsSpiking,
			i.LastSpikeTime,
			i.NextRead,
		); err != nil {
			log.Printf("Error appending to batch: %v", err)
			return err
		}
	}
	return batch.Send()
}
