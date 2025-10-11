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
	cpuSensor         MetricsType
	temperatureSensor MetricsType
	HumiditySensor    MetricsType
	PressureSensor    MetricsType
	GPSSensor         MetricsType
)

func SimulateData() {
	cpuSensor = &cpu.CPU{}
	temperatureSensor = &temperature.Temperature{}
	HumiditySensor = &humidity.Humidity{}
	PressureSensor = &pressure.Pressure{}
	GPSSensor = &gps.GPS{}

	for i := 0; i < NumDevices; i++ {
		go SimulateCPUData(cpuSensor)
		go SimulateTemperatureData(temperatureSensor)
		go SimulateHumidityData(HumiditySensor)
		go SimulatePressureData(PressureSensor)
		go SimulateGPSData(GPSSensor)
		time.Sleep(2 * time.Second) // Stagger the start times slightly
	}
}

func SimulateCPUData(sensor MetricsType) {
	for {
		sensor.GenerateData()
		if cpuTyped, ok := sensor.(*cpu.CPU); ok {
			time.Sleep(cpuTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateTemperatureData(sensor MetricsType) {
	for {
		sensor.GenerateData()
		if tempTyped, ok := sensor.(*temperature.Temperature); ok {
			time.Sleep(tempTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateHumidityData(sensor MetricsType) {
	for {
		sensor.GenerateData()
		if humTyped, ok := sensor.(*humidity.Humidity); ok {
			time.Sleep(humTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulatePressureData(sensor MetricsType) {
	for {
		sensor.GenerateData()
		if presTyped, ok := sensor.(*pressure.Pressure); ok {
			time.Sleep(presTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func SimulateGPSData(sensor MetricsType) {
	for {
		sensor.GenerateData()
		if gpsTyped, ok := sensor.(*gps.GPS); ok {
			time.Sleep(gpsTyped.UpdateInterval * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
