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
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, filter, err := Paginate(r, *s.ch, "gps", "mergeTree")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT id,
		dictGetString(gps_metadatadict, 'loc', device_id) as loc, 
		dictGetString(gps_metadatadict, 'model', device_id) as model,
		dictGetString(gps_metadatadict, 'manufacturer', device_id) as manufacturer,
		dictGetDate(gps_metadatadict, 'install_date', device_id) as install_date,
		device_id, drift_rate, latitude, longitude, altitude, speed, heading, 
		is_moving, next_read_time, updated_at 
	FROM gps %s
	ORDER BY %s %s 
	LIMIT ? OFFSET ?`, filter, order, sort_way)

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

func (s *GPSStore) GetAggregationPerLocation(r *http.Request) (any, error) {
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, filter, err := Paginate(r, *s.ch, "GPS_PER_LOCATION", "incrementalLocMV")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT loc, maxMerge(maxLongitude), avgMerge(avgLatitude), minMerge(minDriftRate)
	FROM GPS_PER_LOCATION group by loc %s
	ORDER BY %s %s 
	LIMIT ? OFFSET ?`, filter, order, sort_way)

	rows, err := (*s.ch).Query(context.Background(), query, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}

	stats := []common.Metrics{}
	for rows.Next() {
		var s common.Metrics
		err := rows.Scan(
			&s.Loc,
			&s.Longitude,
			&s.Latitude,
			&s.DriftRate,
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

func (s *GPSStore) GetAggregationPerModel(r *http.Request) (any, error) {
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, filter, err := Paginate(r, *s.ch, "GPS_PER_MODEL", "incrementalModelMV")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT model, maxMerge(maxHeading), countMerge(countManufacturer)
	FROM GPS_PER_MODEL group by model %s
	ORDER BY %s %s 
	LIMIT ? OFFSET ?`, filter, order, sort_way)

	rows, err := (*s.ch).Query(context.Background(), query, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}

	stats := []common.Metrics{}
	for rows.Next() {
		var s common.Metrics
		err := rows.Scan(
			&s.Model,
			&s.Heading,
			&s.CountManufacturer,
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

func (s *GPSStore) GetDailyAggregationPerModel(r *http.Request) (any, error) {
	order, sort_way, totalPages, totalRows, offset, page, rowsPerPage, filter, err := Paginate(r, *s.ch, "gps_daily_summary", "refreshModelMV")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT model, day, avgMerge(avgSpeed), maxMerge(maxAltitude), 
	sumMerge(sumDriftRate), countMerge(countRecords)
	FROM gps_daily_summary group by (model, day) %s
	ORDER BY %s %s 
	LIMIT ? OFFSET ?`, filter, order, sort_way)

	rows, err := (*s.ch).Query(context.Background(), query, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}

	stats := []common.Metrics{}
	for rows.Next() {
		var s common.Metrics
		err := rows.Scan(
			&s.Model,
			&s.Day,
			&s.Speed,
			&s.Altitude,
			&s.DriftRate,
			&s.CountRecords,
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
