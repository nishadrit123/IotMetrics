package store

import (
	"context"
	"iot/data_simulator/common"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type GPSStore struct {
	ch *clickhouse.Conn
}

func (s *GPSStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO gps (id, device_name, device_id, updated_interval, drift_rate, latitude, longitude, altitude, speed, heading, is_moving, next_read_time)")
	if err != nil {
		log.Printf("Error preparing gps batch: %v", err)
		return err
	}
	for _, i := range data {
		updateTime := time.Now().Add(i.UpdateInterval)

		if err := batch.Append(
			i.Id,
			i.DeviceName,
			i.DeviceId,
			updateTime,
			i.DriftRate,
			i.Latitude,
			i.Longitude,
			i.Altitude,
			i.Speed,
			i.Heading,
			i.IsMoving,
			i.NextRead,
		); err != nil {
			log.Printf("Error appending to gps batch: %v", err)
			return err
		}
	}
	return batch.Send()
}
