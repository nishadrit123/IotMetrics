package store

import (
	"context"
	"fmt"
	"iot/data_simulator/common"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type CPUStore struct {
	ch *clickhouse.Conn
}

func (s *CPUStore) InsertBatch(data []common.Metrics) error {
	batch, err := (*s.ch).PrepareBatch(context.Background(), "INSERT INTO cpu (id, device_name, device_id, baseline_usage, spike_probability, spike_magnitude, noise_level, updated_interval, current_usage, cpu_temperature, is_spiking, last_spike_time, next_read_time)")
	if err != nil {
		log.Printf("Error preparing CPU batch: %v", err)
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
			log.Printf("Error appending to CPU batch: %v", err)
			return err
		}
	}
	return batch.Send()
}

func (s *CPUStore) GetStatistics(r *http.Request) (any, error) {
	// Parse page number from query params, default to 1
	page := 1
	order := "device_id"
	sort_way := "asc"

	q := r.URL.Query()
	if p := q.Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		}
	}
	if o := q.Get("order"); o != "" {
		order = o
	}
	if s := q.Get("sort"); s != "" {
		sort_way = s
	}

	var totalRows uint64
	err := (*s.ch).QueryRow(context.Background(), "SELECT count() FROM cpu").Scan(&totalRows)
	if err != nil {
		return nil, err
	}
	rowsPerPage := 10
	totalPages := int((totalRows + uint64(rowsPerPage) - 1) / uint64(rowsPerPage))
	offset := (page - 1) * rowsPerPage

	query := fmt.Sprintf(`
	SELECT id, 
		dictGetString(cpu_metadatadict, 'hostname', device_id) as hostname, 
		dictGetString(cpu_metadatadict, 'loc', device_id) as loc, 
		dictGetString(cpu_metadatadict, 'model', device_id) as model,
		dictGetInt64(cpu_metadatadict, 'core_count', device_id) as core_count,
		dictGetFloat64(cpu_metadatadict, 'frequency', device_id) as frequency,
		device_id, baseline_usage, spike_probability, spike_magnitude, 
	    noise_level, current_usage, cpu_temperature, is_spiking, 
		last_spike_time, next_read_time, updated_at
	FROM cpu 
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
			&s.HostName,
			&s.Loc,
			&s.Model,
			&s.CoreCount,
			&s.Frequency,
			&s.DeviceId,
			&s.BaselineUsage,
			&s.SpikeProbability,
			&s.SpikeMagnitude,
			&s.NoiseLevel,
			&s.CurrentUsage,
			&s.Temperature,
			&s.IsSpiking,
			&s.LastSpikeTime,
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
