package temperature

import "time"

const (
	SpikeThreshold           = 5.0
	TempratureUpdateInterval = 4
	DeviceName               = "temperature"
)

var DeviceIds = []string{
	"TEMP-9012A", "TEMP-8347B", "TEMP-7231C", "TEMP-6458D", "TEMP-5193E",
	"TEMP-4872F", "TEMP-3920G", "TEMP-2756H", "TEMP-1683I", "TEMP-0549J",
	"TEMP-9234K", "TEMP-8721L", "TEMP-7563M", "TEMP-6312N", "TEMP-5024O",
}

var Locations = []string{
	"Mumbai/Lab-1", "Delhi/Server-Room", "Bangalore/Factory-2", "Hyderabad/Office-1", "Chennai/Data-Center",
	"Pune/Lab-3", "Ahmedabad/Warehouse", "Singapore/R&D", "Frankfurt/Plant-1", "London/HQ",
	"New York/Lab-A", "San Francisco/Lab-B", "Tokyo/Plant-4", "Sydney/Server-Room", "Toronto/Lab-2",
}

var Models = []string{
	"TX-90", "TX-85", "HTR-100", "NanoTemp-X1", "AeroSense-5",
	"ThermoPro-7", "TempEdge-300", "MicroHeat-2", "SenseX-T100", "EnviroMax-9",
	"ThermIQ-10", "TempMate-400", "HeatScan-8", "CoolSense-500", "EcoTherm-6",
}

var Manufacturers = []string{
	"ThermoTech", "EnviroSense", "HeatMaster", "AeroDynamics", "CoolSys",
	"NanoInstruments", "TempEdge Corp", "Sensorify", "HeatLabs", "ThermaVision",
	"EnviroCore", "MicroHeat", "SenseLogic", "DataTherm", "AeroSense Ltd",
}

var InstallDates = []time.Time{
	time.Date(2021, 6, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 11, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 9, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 7, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 5, 30, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 12, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 2, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 8, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 10, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 2, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 6, 21, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 1, 11, 0, 0, 0, 0, time.UTC),
}

var Units = []string{
	"°C", "°C", "°C", "°C", "°C", "°C", "°C", "°C", "°C", "°C",
	"°C", "°C", "°C", "°C", "°C",
}

var BaselineTemps = []float64{
	24.5, 27.0, 22.3, 25.1, 29.4, 26.5, 23.8, 30.2, 21.9, 28.0,
	24.7, 25.5, 26.8, 27.9, 23.5,
}

var SpikeProbabilities = []float64{
	0.10, 0.05, 0.08, 0.12, 0.15, 0.09, 0.06, 0.20, 0.18, 0.07,
	0.11, 0.13, 0.05, 0.17, 0.10,
}

var SpikeMagnitudes = []float64{
	3.5, 4.0, 2.5, 5.0, 6.5, 4.2, 3.0, 7.0, 5.8, 4.8,
	6.2, 3.9, 2.8, 5.6, 6.0,
}

var NoiseLevels = []float64{
	0.3, 0.4, 0.2, 0.5, 0.6, 0.3, 0.4, 0.7, 0.5, 0.4,
	0.6, 0.3, 0.2, 0.5, 0.4,
}

var UpdateIntervals = []time.Duration{
	1 * time.Second, 2 * time.Second, 3 * time.Second, 1 * time.Second, 5 * time.Second,
	4 * time.Second, 2 * time.Second, 3 * time.Second, 1 * time.Second, 2 * time.Second,
	3 * time.Second, 4 * time.Second, 2 * time.Second, 1 * time.Second, 5 * time.Second,
}

var DriftRates = []float64{
	0.02, 0.05, 0.03, 0.04, 0.01, 0.02, 0.05, 0.03, 0.06, 0.02,
	0.04, 0.01, 0.03, 0.05, 0.02,
}

var CurrentTemps = []float64{
	25.1, 26.8, 22.9, 25.6, 30.1, 27.3, 24.5, 31.0, 22.7, 29.0,
	25.3, 26.2, 27.4, 28.1, 23.9,
}
