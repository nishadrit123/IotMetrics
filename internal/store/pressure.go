package store

import (
	"context"
	"fmt"
	"iot/data_simulator/common"
	"log"
	"net/http"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type PressureStore struct {
	ch *clickhouse.Conn
}

func (s *PressureStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO pressure (id, device_name, device_id, baseline_pressure, spike_probability, spike_magnitude, noise_level, updated_interval, drift_rate, current_pressure, is_spiking, last_spike_time, next_read_time, trend)")
	if err != nil {
		log.Printf("Error preparing Pressure batch: %v", err)
		return err
	}
	for _, i := range data {
		updateTime := time.Now().Add(i.UpdateInterval)

		if err := batch.Append(
			i.Id,
			i.DeviceName,
			i.DeviceId,
			i.BaselinePressure,
			i.SpikeProbability,
			i.SpikeMagnitude,
			i.NoiseLevel,
			updateTime,
			i.DriftRate,
			i.CurrentPressure,
			i.IsSpiking,
			i.LastSpikeTime,
			i.NextRead,
			i.Trend,
		); err != nil {
			log.Printf("Error appending to Pressure batch: %v", err)
			return err
		}
	}
	return batch.Send()
}

func (s *PressureStore) GetStatistics(r *http.Request) (any, error) {
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, _, err := Paginate(r, *s.ch, "pressure")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT id,
		dictGetString(pressure_metadatadict, 'loc', device_id) as loc, 
		dictGetString(pressure_metadatadict, 'model', device_id) as model,
		dictGetString(pressure_metadatadict, 'manufacturer', device_id) as manufacturer,
		dictGetDate(pressure_metadatadict, 'install_date', device_id) as install_date,
		device_id, baseline_pressure, spike_probability, spike_magnitude, 
	    noise_level, drift_rate, current_pressure, is_spiking, 
		last_spike_time, trend, next_read_time, updated_at 
	FROM pressure 
	ORDER BY %s %s 
	LIMIT ? OFFSET ?`, order, sort_way)

	rows, err := (*s.ch).Query(context.Background(), query, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}

	stats := []common.Metrics{}
	for rows.Next() {
		var s common.Metrics
		err := rows.Scan(
			&s.Id,
			&s.Loc,
			&s.Model,
			&s.Manufacturer,
			&s.InstallDate,
			&s.DeviceId,
			&s.BaselinePressure,
			&s.SpikeProbability,
			&s.SpikeMagnitude,
			&s.NoiseLevel,
			&s.DriftRate,
			&s.CurrentPressure,
			&s.IsSpiking,
			&s.LastSpikeTime,
			&s.Trend,
			&s.NextRead,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	result := map[string]any{
		"data":        stats,
		"page":        page,
		"total_pages": totalPages,
		"total_rows":  totalRows,
	}
	return result, nil
}
