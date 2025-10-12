package main

const (
	NumDevices          = 5
	CPUChanSize         = 500
	TemperatureChanSize = 500
	HumidityChanSize    = 500
	PressureChanSize    = 500
	GPSChanSize         = 500
)

type MetricsType interface {
	GenerateData(metricsChan chan string)
}
