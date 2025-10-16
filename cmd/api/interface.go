package main

import "iot/data_simulator/common"

const (
	NumDevices = 5

	CPUChanSize  = 500
	CPUBatchSize = 100

	TemperatureChanSize  = 500
	TemperatureBatchSize = 100

	HumidityChanSize  = 500
	HumidityBatchSize = 100

	PressureChanSize  = 500
	PressureBatchSize = 100

	GPSChanSize  = 500
	GPSBatchSize = 100
)

type MetricsType interface {
	GenerateData(metricsChan chan common.Metrics)
}
