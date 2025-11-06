package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type HeatMapStore struct {
	ch *clickhouse.Conn
}

type Locations struct {
	Locs []string `json:"locs,omitempty"`
}

func (s *HeatMapStore) GetHeatMap(payload []byte) (any, error) {
	var (
		l  Locations
		hm [][]float64
	)

	hm = make([][]float64, len(Tables))
	for k := range len(Tables) {
		hm[k] = make([]float64, len(Tables))
	}

	err := json.Unmarshal(payload, &l)
	if err != nil {
		log.Printf("Error unmarshalling payload for heatmap %v", err)
		return nil, err
	}

	for i := 0; i < len(Tables); i++ {
		for j := i; j < len(Tables); j++ {
			if i == j {
				hm[i][j] = 1
			} else {
				corr, err := calculateCorrelation(*s.ch, l, i, j)
				if err != nil {
					continue
				}
				hm[i][j] = corr
				hm[j][i] = corr
			}
		}
	}

	result := map[string]any{
		"sequence": Tables,
		"data":     hm,
	}
	return result, nil
}

func calculateCorrelation(ch clickhouse.Conn, l Locations, i, j int) (float64, error) {
	var (
		whereClause string
		correlation float64
	)

	if len(l.Locs) > 1 {
		whereClause += "WHERE loc IN ( "
		for i, loc := range l.Locs {
			whereClause += fmt.Sprintf("'%v'", loc)
			if i != len(l.Locs)-1 {
				whereClause += ", "
			}
		}
		whereClause += " )"
	}

	query := fmt.Sprintf(`
		SELECT corr(var1, var2) FROM 
		( 
			SELECT loc, day, 
			avgMerge(%v) as var1, avgMerge(%v) as var2
			FROM %v JOIN %v USING (loc, day) 
			%v
			GROUP BY (loc, day) 
		)
	`, TableMap[Tables[i]], TableMap[Tables[j]], Tables[i], Tables[j], whereClause)

	err := ch.QueryRow(context.Background(), query).Scan(&correlation)
	if err != nil {
		log.Printf("Error calcuating corr() %v", err)
		return 0, err
	}
	return correlation, nil
}
