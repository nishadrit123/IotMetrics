package store

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var Tables = []string{
	"cpu_daily_summary", "temperature_daily_summary",
	"pressure_daily_summary", "humidity_daily_summary",
}

var TableMap map[string]string

func init() {
	TableMap = make(map[string]string)
	TableMap["cpu_daily_summary"] = "avgCurrentUsage"
	TableMap["temperature_daily_summary"] = "avgCurrentTemperature"
	TableMap["pressure_daily_summary"] = "avgCurrentPressure"
	TableMap["humidity_daily_summary"] = "avgCurrentHumidity"
}

func Paginate(r *http.Request, ch clickhouse.Conn, tableName, tableType string) (string, string, int, int, int, string, []any, error) {
	var (
		totalRows uint64
		query     string
		order     string
		groupBy   string
		filter    string
		final     string
		args      []any
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
		filter, args = setFilter(f, tableType)
	}

	if filter == "" {
		if tableType != "mergeTree" {
			final = "FINAL"
		}
		query = fmt.Sprintf("SELECT count() FROM %v %v", tableName, final)
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

	// Mitigation from SQL injection attack
	err := ch.QueryRow(context.Background(), query, args...).Scan(&totalRows)
	if err != nil {
		return order, sort_way, 0, 0, 0, filter, args, err
	}

	rowsPerPage := 10
	totalPages := int((totalRows + uint64(rowsPerPage) - 1) / uint64(rowsPerPage))
	offset := (page - 1) * rowsPerPage

	args = append(args, rowsPerPage)
	args = append(args, offset)

	return order, sort_way, totalPages, int(totalRows), page, filter, args, nil
}

func setFilter(f, tableType string) (string, []any) {
	var (
		filter string
		op     string
		col    string
		args   []any
	)
	placeHolder := "?"
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
			col = parts[0]
			op = "="
			args = append(args, fmt.Sprintf("%s", parts[1]))
		}
		if strings.Contains(fs, "~") {
			parts := strings.Split(fs, "~")
			col = parts[0]
			op = "LIKE"
			args = append(args, fmt.Sprintf("%%%s%%", parts[1]))
		}
		if strings.Contains(fs, ">") {
			parts := strings.Split(fs, ">")
			intPart1, _ := strconv.Atoi(parts[1])
			col = parts[0]
			op = ">"
			args = append(args, intPart1)
		}
		if strings.Contains(fs, "<") {
			parts := strings.Split(fs, "<")
			intPart1, _ := strconv.Atoi(parts[1])
			col = parts[0]
			op = "<"
			args = append(args, intPart1)
		}
		if strings.Contains(tableType, "MV") {
			col = getCombinators(col)
		}

		// This is done since we reply on dictGet() for loc and model in case of mergeTree
		// and parameterized '?' does not work with it
		if strings.Contains(tableType, "mergeTree") && (col == "loc" || col == "model") {
			filter += fmt.Sprintf("%s %s %s", col, op, fmt.Sprintf("'%v'", args[len(args)-1]))
			args = args[:len(args)-1]
		} else {
			filter += fmt.Sprintf("%s %s %s", col, op, placeHolder)
		}
	}
	return filter, args
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
