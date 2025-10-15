package main

import (
	"iot/data_simulator/cpu"
	"iot/data_simulator/gps"
	"iot/data_simulator/humidity"
	"iot/data_simulator/pressure"
	"iot/data_simulator/temperature"
	"time"
)

var (
	cpuBuffer         []string
	temperatureBuffer []string
	humidityBuffer    []string
	pressureBuffer    []string
	gpsBuffer         []string
)

func SimulateData() {
	cpu_chan := make(chan string, CPUChanSize)
	temperature_chan := make(chan string, TemperatureChanSize)
	humidity_chan := make(chan string, HumidityChanSize)
	pressure_chan := make(chan string, PressureChanSize)
	gps_chan := make(chan string, GPSChanSize)

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
					// InsertBatchtoClickHouse("cpu_metrics", cpuBuffer)
				}
			case data := <-temperature_chan:
				temperatureBuffer = append(temperatureBuffer, data)
				if len(temperatureBuffer) > TemperatureBatchSize {
					// InsertBatchtoClickHouse("temperature_metrics", temperatureBuffer)
				}
			case data := <-humidity_chan:
				humidityBuffer = append(humidityBuffer, data)
				if len(humidityBuffer) > HumidityBatchSize {
					// InsertBatchtoClickHouse("humidity_metrics", humidityBuffer)
				}
			case data := <-pressure_chan:
				pressureBuffer = append(pressureBuffer, data)
				if len(pressureBuffer) > PressureBatchSize {
					// InsertBatchtoClickHouse("pressure_metrics", pressureBuffer)
				}
			case data := <-gps_chan:
				gpsBuffer = append(gpsBuffer, data)
				if len(gpsBuffer) > GPSBatchSize {
					// InsertBatchtoClickHouse("gps_metrics", gpsBuffer)
				}
			default:
				time.Sleep(100 * time.Millisecond) // Prevent busy waiting
			}
		}
	}()
}

func SimulateCPUData(sensor MetricsType, cpu_chan chan string) {
	for {
		sensor.GenerateData(cpu_chan)
		if cpuTyped, ok := sensor.(*cpu.CPU); ok {
			time.Sleep(cpuTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateTemperatureData(sensor MetricsType, temperature_chan chan string) {
	for {
		sensor.GenerateData(temperature_chan)
		if tempTyped, ok := sensor.(*temperature.Temperature); ok {
			time.Sleep(tempTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateHumidityData(sensor MetricsType, humidity_chan chan string) {
	for {
		sensor.GenerateData(humidity_chan)
		if humTyped, ok := sensor.(*humidity.Humidity); ok {
			time.Sleep(humTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulatePressureData(sensor MetricsType, pressure_chan chan string) {
	for {
		sensor.GenerateData(pressure_chan)
		if presTyped, ok := sensor.(*pressure.Pressure); ok {
			time.Sleep(presTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateGPSData(sensor MetricsType, gps_chan chan string) {
	for {
		sensor.GenerateData(gps_chan)
		if gpsTyped, ok := sensor.(*gps.GPS); ok {
			time.Sleep(gpsTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
