package main

import (
	"iot/data_simulator/cpu"
	"iot/data_simulator/gps"
	"iot/data_simulator/humidity"
	"iot/data_simulator/pressure"
	"iot/data_simulator/temperature"
	"log"
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
	cpu_chan := make(chan string, CPUChanSize)

	temperatureSensor = &temperature.Temperature{}
	temperature_chan := make(chan string, TemperatureChanSize)

	HumiditySensor = &humidity.Humidity{}
	humidity_chan := make(chan string, HumidityChanSize)

	PressureSensor = &pressure.Pressure{}
	pressure_chan := make(chan string, PressureChanSize)

	GPSSensor = &gps.GPS{}
	gps_chan := make(chan string, GPSChanSize)

	go func() {
		for i := 0; i < NumDevices; i++ {
			go SimulateCPUData(cpuSensor, cpu_chan)
			go SimulateTemperatureData(temperatureSensor, temperature_chan)
			go SimulateHumidityData(HumiditySensor, humidity_chan)
			go SimulatePressureData(PressureSensor, pressure_chan)
			go SimulateGPSData(GPSSensor, gps_chan)
			time.Sleep(2 * time.Second) // Stagger the start times slightly
		}
	}()

	go func() {
		for {
			select {
			case data := <-cpu_chan:
				log.Printf("cpu channel: %v\n\n", data)
			case data := <-temperature_chan:
				log.Printf("temperature channel: %v\n\n", data)
			case data := <-humidity_chan:
				log.Printf("humidity channel: %v\n\n", data)
			case data := <-pressure_chan:
				log.Printf("pressure channel: %v\n\n", data)
			case data := <-gps_chan:
				log.Printf("gps channel: %v\n\n", data)
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
