package store

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Paginate(r *http.Request, ch clickhouse.Conn, tableName, tableType string) (string, string, int, int, int, int, int, string, error) {
	var (
		totalRows uint64
		query     string
		order     string
		groupBy   string
	)

	if tableType == "mergeTree" {
		order = "device_id"
	} else if tableType == "incrementalLocMV" || tableType == "refreshLocMV" {
		order = "loc"
	} else if tableType == "incrementalModelMV" || tableType == "refreshModelMV" {
		order = "model"
	}

	page := 1
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
		if strings.Contains(tableType, "MV") {
			order = getCombinators(order)
		}
	}
	if s := q.Get("sort"); s != "" {
		sort_way = s
	}
	if f := q.Get("filter"); f != "" {
		filter = setFilter(f, tableType)
	}

	if filter == "" {
		query = fmt.Sprintf("SELECT count() FROM %v", tableName)
	} else {
		if tableType == "mergeTree" {
			locStr := fmt.Sprintf("dictGetString(%v_metadatadict, 'loc', device_id)", tableName)
			modelStr := fmt.Sprintf("dictGetString(%v_metadatadict, 'model', device_id)", tableName)
			filter = strings.Replace(filter, "loc", locStr, -1)
			filter = strings.Replace(filter, "model", modelStr, -1)
			query = fmt.Sprintf("SELECT count() FROM %v %v", tableName, filter)
		} else if strings.Contains(tableType, "MV") {
			if strings.Contains(tableType, "incremental") {
				if strings.Contains(strings.ToLower(tableName), "loc") {
					groupBy = "loc"
				} else if strings.Contains(strings.ToLower(tableName), "model") {
					groupBy = "model"
				}
				query = fmt.Sprintf("SELECT count() FROM (SELECT %v FROM %v GROUP BY %v %v)", groupBy, tableName, groupBy, filter)
			} else if strings.Contains(tableType, "refresh") {
				if strings.Contains(strings.ToLower(tableName), "gps") {
					groupBy = "model"
				} else {
					groupBy = "loc"
				}
				query = fmt.Sprintf("SELECT count() FROM (SELECT %v FROM %v GROUP BY (%v, day) %v)", groupBy, tableName, groupBy, filter)
			}
		}
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

func setFilter(f, tableType string) string {
	var filter string
	filterSlice := strings.Split(f, ":")
	for i, fs := range filterSlice {
		if i == 0 {
			if tableType == "mergeTree" {
				filter += "WHERE "
			} else if strings.Contains(tableType, "MV") {
				filter += " HAVING "
			}
		} else {
			filter += " AND "
		}
		if strings.Contains(fs, "=") {
			parts := strings.Split(fs, "=")
			if tableType == "mergeTree" {
				filter += fmt.Sprintf("%s = '%s'", parts[0], parts[1])
			} else if strings.Contains(tableType, "MV") {
				filter += fmt.Sprintf("%s = '%s'", getCombinators(parts[0]), parts[1])
			}
		}
		if strings.Contains(fs, "~") {
			parts := strings.Split(fs, "~")
			if tableType == "mergeTree" {
				filter += fmt.Sprintf("%s LIKE '%%%s%%'", parts[0], parts[1])
			} else if strings.Contains(tableType, "MV") {
				filter += fmt.Sprintf("%s LIKE '%%%s%%'", getCombinators(parts[0]), parts[1])
			}
		}
		if strings.Contains(fs, ">") {
			parts := strings.Split(fs, ">")
			intPart1, _ := strconv.Atoi(parts[1])
			if tableType == "mergeTree" {
				filter += fmt.Sprintf("%s > %v", parts[0], intPart1)
			} else if strings.Contains(tableType, "MV") {
				filter += fmt.Sprintf("%s > %v", getCombinators(parts[0]), intPart1)
			}
		}
		if strings.Contains(fs, "<") {
			parts := strings.Split(fs, "<")
			intPart1, _ := strconv.Atoi(parts[1])
			if tableType == "mergeTree" {
				filter += fmt.Sprintf("%s < %v", parts[0], intPart1)
			} else if strings.Contains(tableType, "MV") {
				filter += fmt.Sprintf("%s < %v", getCombinators(parts[0]), intPart1)
			}
		}
	}
	return filter
}

func getCombinators(order string) string {
	if strings.Contains(order, "max") {
		return fmt.Sprintf("maxMerge(%s)", order)
	} else if strings.Contains(order, "min") {
		return fmt.Sprintf("minMerge(%s)", order)
	} else if strings.Contains(order, "avg") {
		return fmt.Sprintf("avgMerge(%s)", order)
	} else if strings.Contains(order, "sum") || strings.Contains(order, "total") {
		return fmt.Sprintf("sumMerge(%s)", order)
	} else if strings.Contains(order, "count") {
		return fmt.Sprintf("countMerge(%s)", order)
	} else if strings.Contains(order, "uniq") {
		return fmt.Sprintf("uniqMerge(%s)", order)
	}
	return order
}
