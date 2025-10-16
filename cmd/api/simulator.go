package main

import (
	"iot/data_simulator/common"
	"iot/data_simulator/cpu"
	"iot/data_simulator/gps"
	"iot/data_simulator/humidity"
	"iot/data_simulator/pressure"
	"iot/data_simulator/temperature"
	"iot/internal/store"
	"time"
)

var (
	cpuBuffer         []common.Metrics
	temperatureBuffer []common.Metrics
	humidityBuffer    []common.Metrics
	pressureBuffer    []common.Metrics
	gpsBuffer         []common.Metrics
)

func SimulateData(st store.Store) {
	cpu_chan := make(chan common.Metrics, CPUChanSize)
	temperature_chan := make(chan common.Metrics, TemperatureChanSize)
	humidity_chan := make(chan common.Metrics, HumidityChanSize)
	pressure_chan := make(chan common.Metrics, PressureChanSize)
	gps_chan := make(chan common.Metrics, GPSChanSize)

	go func() {
		for i := 0; i < NumDevices; i++ {
			go SimulateCPUData(&cpu.CPU{}, cpu_chan)
			go SimulateTemperatureData(&temperature.Temperature{}, temperature_chan)
			go SimulateHumidityData(&humidity.Humidity{}, humidity_chan)
			go SimulatePressureData(&pressure.Pressure{}, pressure_chan)
			go SimulateGPSData(&gps.GPS{}, gps_chan)
			time.Sleep(2 * time.Second) // Stagger the start times slightly
		}
	}()

	go func() {
		for {
			select {
			case data := <-cpu_chan:
				cpuBuffer = append(cpuBuffer, data)
				if len(cpuBuffer) > CPUBatchSize {
					st.CPU.(*store.CPUStore).InsertBatch(cpuBuffer)
					cpuBuffer = cpuBuffer[:0]
				}
			case data := <-temperature_chan:
				temperatureBuffer = append(temperatureBuffer, data)
				if len(temperatureBuffer) > TemperatureBatchSize {
					st.Temperature.(*store.TemperatureStore).InsertBatch(temperatureBuffer)
					temperatureBuffer = temperatureBuffer[:0]
				}
			case data := <-humidity_chan:
				humidityBuffer = append(humidityBuffer, data)
				if len(humidityBuffer) > HumidityBatchSize {
					st.Humidity.(*store.HumidityStore).InsertBatch(humidityBuffer)
					humidityBuffer = humidityBuffer[:0]
				}
			case data := <-pressure_chan:
				pressureBuffer = append(pressureBuffer, data)
				if len(pressureBuffer) > PressureBatchSize {
					st.Pressure.(*store.PressureStore).InsertBatch(pressureBuffer)
					pressureBuffer = pressureBuffer[:0]
				}
			case data := <-gps_chan:
				gpsBuffer = append(gpsBuffer, data)
				if len(gpsBuffer) > GPSBatchSize {
					st.GPS.(*store.GPSStore).InsertBatch(gpsBuffer)
					gpsBuffer = gpsBuffer[:0]
				}
			default:
				time.Sleep(100 * time.Millisecond) // Prevent busy waiting
			}
		}
	}()
}

func SimulateCPUData(sensor MetricsType, cpu_chan chan common.Metrics) {
	for {
		sensor.GenerateData(cpu_chan)
		if cpuTyped, ok := sensor.(*cpu.CPU); ok {
			time.Sleep(cpuTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateTemperatureData(sensor MetricsType, temperature_chan chan common.Metrics) {
	for {
		sensor.GenerateData(temperature_chan)
		if tempTyped, ok := sensor.(*temperature.Temperature); ok {
			time.Sleep(tempTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateHumidityData(sensor MetricsType, humidity_chan chan common.Metrics) {
	for {
		sensor.GenerateData(humidity_chan)
		if humTyped, ok := sensor.(*humidity.Humidity); ok {
			time.Sleep(humTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulatePressureData(sensor MetricsType, pressure_chan chan common.Metrics) {
	for {
		sensor.GenerateData(pressure_chan)
		if presTyped, ok := sensor.(*pressure.Pressure); ok {
			time.Sleep(presTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateGPSData(sensor MetricsType, gps_chan chan common.Metrics) {
	for {
		sensor.GenerateData(gps_chan)
		if gpsTyped, ok := sensor.(*gps.GPS); ok {
			time.Sleep(gpsTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
