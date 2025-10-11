package humidity

import "time"

const (
	SpikeThreshold         = 10.0
	DeviceName             = "humidity"
	HumidityUpdateInterval = 4
)

var DeviceIds = []string{
	"HUM-9001A", "HUM-8423B", "HUM-7310C", "HUM-6598D", "HUM-5246E",
	"HUM-4875F", "HUM-3762G", "HUM-2691H", "HUM-1539I", "HUM-0452J",
	"HUM-9238K", "HUM-8725L", "HUM-7603M", "HUM-6421N", "HUM-5087O",
}

var Locations = []string{
	"Warehouse-Zone-A", "Warehouse-Zone-B", "Factory-Assembly-1", "Factory-Assembly-2", "Office-Floor-3",
	"Server-Room-1", "Server-Room-2", "Cold-Storage-Unit", "Greenhouse-1", "Greenhouse-2",
	"Basement-Lab", "Quality-Control-Room", "R&D-Lab-1", "Testing-Facility-2", "Warehouse-Zone-C",
}

var Models = []string{
	"HX-200", "HX-210", "HumidPro-5", "DrySense-X1", "MoistTrack-300",
	"EnviroHum-10", "AeroHumid-8", "NanoHumidity-2", "HygroSmart-500", "AtmosTrack-9",
	"HydroScan-7", "AirSense-H2", "EnviroWave-4", "HumidityEdge-12", "ClimaCheck-8",
}

var Manufacturers = []string{
	"EnviroTech", "AeroDynamics", "ThermoLogic", "NanoInstruments", "CoolSys",
	"Sensorify", "HumidLabs", "AtmosCorp", "MoistureMax", "DataTherm",
	"EnviroCore", "MicroSense", "AirSmart", "ClimaEdge", "HydroSystems",
}

var InstallDates = []time.Time{
	time.Date(2021, 3, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 8, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 2, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 4, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 11, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 9, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 1, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 3, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 7, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 10, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 5, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
}

var Units = []string{
	"%RH", "%RH", "%RH", "%RH", "%RH",
	"%RH", "%RH", "%RH", "%RH", "%RH",
	"%RH", "%RH", "%RH", "%RH", "%RH",
}

var BaselineHumidity = []float64{
	45.0, 55.0, 60.0, 50.0, 40.0,
	35.0, 38.0, 70.0, 80.0, 75.0,
	65.0, 55.0, 60.0, 50.0, 42.0,
}

var SpikeProbabilities = []float64{
	0.10, 0.08, 0.05, 0.12, 0.15, 0.09, 0.07, 0.20, 0.18, 0.11,
	0.10, 0.14, 0.09, 0.16, 0.08,
}

var SpikeMagnitudes = []float64{
	10.0, 8.5, 12.0, 7.0, 6.0,
	5.5, 9.0, 15.0, 18.0, 12.5,
	10.5, 8.0, 9.5, 11.0, 7.5,
}

var NoiseLevels = []float64{
	0.5, 0.7, 0.3, 0.8, 0.4,
	0.5, 0.6, 1.0, 0.9, 0.8,
	0.7, 0.6, 0.4, 0.8, 0.5,
}

var DriftRates = []float64{
	0.01, 0.03, 0.02, 0.01, 0.05,
	0.02, 0.03, 0.04, 0.06, 0.05,
	0.02, 0.03, 0.01, 0.04, 0.02,
}

var CurrentHumidity = []float64{
	46.2, 56.4, 62.0, 51.3, 39.8,
	36.7, 39.5, 72.4, 81.5, 76.1,
	66.8, 56.2, 61.4, 52.5, 43.3,
}
