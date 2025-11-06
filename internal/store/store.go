package store

import "github.com/ClickHouse/clickhouse-go/v2"

type Store struct {
	CPU         interface{}
	GPS         interface{}
	Humidity    interface{}
	Pressure    interface{}
	Temperature interface{}
	HeatMap     interface{}
}

func NewStore(ch *clickhouse.Conn) Store {
	return Store{
		CPU:         &CPUStore{ch},
		GPS:         &GPSStore{ch},
		Humidity:    &HumidityStore{ch},
		Pressure:    &PressureStore{ch},
		Temperature: &TemperatureStore{ch},
		HeatMap:     &HeatMapStore{ch},
	}
}
