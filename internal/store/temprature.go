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

type TemperatureStore struct {
	ch *clickhouse.Conn
}

func (s *TemperatureStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO temperature (id, device_name, device_id, baseline_temperature, spike_probability, spike_magnitude, noise_level, updated_interval, drift_rate, current_temperature, is_spiking, last_spike_time, next_read_time, trend)")
	if err != nil {
		log.Printf("Error preparing Temperature batch: %v", err)
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
			log.Printf("Error appending to Temperature batch: %v", err)
			return err
		}
	}
	return batch.Send()
}

func (s *TemperatureStore) GetStatistics(r *http.Request) (any, error) {
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, _, err := Paginate(r, *s.ch, "temperature")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT id,
		dictGetString(temperature_metadatadict, 'loc', device_id) as loc, 
		dictGetString(temperature_metadatadict, 'model', device_id) as model,
		dictGetString(temperature_metadatadict, 'manufacturer', device_id) as manufacturer,
		dictGetDate(temperature_metadatadict, 'install_date', device_id) as install_date,
		device_id, baseline_temperature, spike_probability, spike_magnitude, 
	    noise_level, drift_rate, current_temperature, is_spiking, 
		last_spike_time, trend, next_read_time, updated_at 
	FROM temperature 
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
			&s.BaselineTemp,
			&s.SpikeProbability,
			&s.SpikeMagnitude,
			&s.NoiseLevel,
			&s.DriftRate,
			&s.CurrentTemp,
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
