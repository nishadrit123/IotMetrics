package pressure

import "time"

const (
	SpikeThreshold         = 5.0
	PressureUpdateInterval = 3
	DeviceName             = "pressure"
)

var DeviceIds = []string{
	"PRS-1001A", "PRS-1002B", "PRS-1003C", "PRS-1004D", "PRS-1005E",
	"PRS-1006F", "PRS-1007G", "PRS-1008H", "PRS-1009I", "PRS-1010J",
	"PRS-1011K", "PRS-1012L", "PRS-1013M", "PRS-1014N", "PRS-1015O",
}

var Locations = []string{
	"Factory Floor-1", "Factory Floor-2", "Warehouse Zone-A", "Warehouse Zone-B", "Laboratory-1",
	"Laboratory-2", "Assembly Line-1", "Assembly Line-2", "Test Chamber-1", "Test Chamber-2",
	"Storage Room-1", "Storage Room-2", "Maintenance Area", "Packaging Unit", "Shipping Dock",
}

var Models = []string{
	"PX-500", "PX-510", "PressurePro-100", "PressureMax-200", "EnviroPress-50",
	"ProPress-300", "SensorPress-X", "AeroPressure-1", "HydroPress-5", "IndustrialPX-900",
	"PressureEdge-10", "SmartPress-20", "TestPX-77", "LabPress-33", "PressureTrack-8",
}

var Manufacturers = []string{
	"EnviroTech", "AeroDynamics", "ThermoLogic", "NanoInstruments", "PressSys",
	"Sensorify", "HydroLabs", "AtmosCorp", "ProSensors", "DataTherm",
	"MicroSense", "AirSmart", "ClimaEdge", "PressureWorks", "HydroSystems",
}

var InstallDates = []time.Time{
	time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 5, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 2, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 3, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 8, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 9, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 4, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2022, 7, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2020, 10, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2023, 5, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
}

var BaselinePressure = []float64{
	101325, 100800, 102000, 101500, 100900,
	101200, 101000, 102200, 101800, 100700,
	101100, 101400, 101700, 101050, 100950,
}

var SpikeProbabilities = []float64{
	0.05, 0.07, 0.03, 0.10, 0.08, 0.06, 0.04, 0.12, 0.09, 0.05,
	0.07, 0.03, 0.10, 0.06, 0.08,
}

var SpikeMagnitudes = []float64{
	5.0, 4.0, 6.0, 3.5, 4.5, 5.5, 4.2, 6.2, 3.8, 5.1,
	4.3, 5.0, 6.5, 4.0, 4.8,
}

var NoiseLevels = []float64{
	0.5, 0.6, 0.4, 0.7, 0.5, 0.6, 0.3, 0.8, 0.5, 0.7,
	0.4, 0.6, 0.5, 0.7, 0.4,
}

var DriftRates = []float64{
	0.01, 0.02, 0.03, 0.01, 0.04,
	0.02, 0.03, 0.01, 0.05, 0.02,
	0.03, 0.01, 0.04, 0.02, 0.03,
}

var CurrentPressure = []float64{
	101400, 100850, 102050, 101600, 100950,
	101250, 101050, 102250, 101850, 100720,
	101150, 101450, 101750, 101080, 100980,
}
