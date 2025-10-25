package store

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Paginate(r *http.Request, ch clickhouse.Conn, tableName string) (string, string, int, int, int, int, int, string, error) {
	var (
		totalRows uint64
		query     string
	)
	page := 1
	order := "device_id"
	sort_way := "asc"
	filter := ""

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
	if f := q.Get("filter"); f != "" {
		filter = setFilter(f)
	}

	if filter == "" {
		query = fmt.Sprintf("SELECT count() FROM %v", tableName)
	} else {
		locStr := fmt.Sprintf("dictGetString(%v_metadatadict, 'loc', device_id)", tableName)
		modelStr := fmt.Sprintf("dictGetString(%v_metadatadict, 'model', device_id)", tableName)
		filter = strings.Replace(filter, "loc", locStr, -1)
		filter = strings.Replace(filter, "model", modelStr, -1)
		query = fmt.Sprintf("SELECT count() FROM %v %v", tableName, filter)
	}
	err := ch.QueryRow(context.Background(), query).Scan(&totalRows)
	if err != nil {
		return order, sort_way, 0, 0, 0, 0, 0, filter, err
	}
	rowsPerPage := 10
	totalPages := int((totalRows + uint64(rowsPerPage) - 1) / uint64(rowsPerPage))
	offset := (page - 1) * rowsPerPage

	return order, sort_way, totalPages, int(totalRows), offset, page, rowsPerPage, filter, nil
}

func setFilter(f string) string {
	var filter string
	filterSlice := strings.Split(f, ":")
	for i, fs := range filterSlice {
		if i == 0 {
			filter += "WHERE "
		} else {
			filter += " AND "
		}
		if strings.Contains(fs, "=") {
			parts := strings.Split(fs, "=")
			filter += fmt.Sprintf("%s = '%s'", parts[0], parts[1])
		}
		if strings.Contains(fs, "~") {
			parts := strings.Split(fs, "~")
			filter += fmt.Sprintf("%s LIKE '%%%s%%'", parts[0], parts[1])
		}
		if strings.Contains(fs, ">") {
			parts := strings.Split(fs, ">")
			intPart1, _ := strconv.Atoi(parts[1])
			filter += fmt.Sprintf("%s > %v", parts[0], intPart1)
		}
		if strings.Contains(fs, "<") {
			parts := strings.Split(fs, "<")
			intPart1, _ := strconv.Atoi(parts[1])
			filter += fmt.Sprintf("%s < %v", parts[0], intPart1)
		}
	}
	return filter
}
