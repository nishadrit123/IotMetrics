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

func (s *GPSStore) GetStatistics(r *http.Request) (any, error) {
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
	err := (*s.ch).QueryRow(context.Background(), "SELECT count() FROM gps").Scan(&totalRows)
	if err != nil {
		return nil, err
	}
	rowsPerPage := 10
	totalPages := int((totalRows + uint64(rowsPerPage) - 1) / uint64(rowsPerPage))
	offset := (page - 1) * rowsPerPage

	query := fmt.Sprintf(`
	SELECT id,
		dictGetString(gps_metadatadict, 'loc', device_id) as loc, 
		dictGetString(gps_metadatadict, 'model', device_id) as model,
		dictGetString(gps_metadatadict, 'manufacturer', device_id) as manufacturer,
		dictGetDate(gps_metadatadict, 'install_date', device_id) as install_date,
		device_id, drift_rate, latitude, longitude, altitude, speed, heading, 
		is_moving, next_read_time, updated_at 
	FROM gps 
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
			&s.DriftRate,
			&s.Latitude,
			&s.Longitude,
			&s.Altitude,
			&s.Speed,
			&s.Heading,
			&s.IsMoving,
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
