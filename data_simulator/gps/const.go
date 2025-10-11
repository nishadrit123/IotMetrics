package gps

import "time"

const (
	DeviceName        = "gps"
	GPSUpdateInterval = 3
)

var DeviceIds = []string{
	"GPS-1001A", "GPS-1002B", "GPS-1003C", "GPS-1004D", "GPS-1005E",
	"GPS-1006F", "GPS-1007G", "GPS-1008H", "GPS-1009I", "GPS-1010J",
	"GPS-1011K", "GPS-1012L", "GPS-1013M", "GPS-1014N", "GPS-1015O",
}

var Locations = []string{
	"Vehicle A", "Vehicle B", "Truck 1", "Truck 2", "Drone 1",
	"Drone 2", "Boat 1", "Boat 2", "Lab Test Unit 1", "Lab Test Unit 2",
	"Warehouse Vehicle 1", "Warehouse Vehicle 2", "Field Unit 1", "Field Unit 2", "Delivery Van 1",
}

var Models = []string{
	"GNSS-XYZ", "GNSS-ABC", "GPSPro-100", "GeoTrack-200", "NavSat-50",
	"SpeedTrack-300", "AeroNav-1", "HydroNav-5", "SmartGPS-900", "NavEdge-10",
	"MicroGPS-20", "UltraNav-15", "FieldTrack-33", "LabNav-77", "ProGPS-8",
}

var Manufacturers = []string{
	"NavTech", "GeoDynamics", "AeroInstruments", "HydroSystems", "MicroNav",
	"GPSify", "TrackLabs", "SatelliteCorp", "DataNav", "PrecisionGPS",
	"EnviroNav", "SmartSat", "ClimaTrack", "GeoEdge", "NavSystems",
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

var SpeedNoiseLevels = []float64{
	0.5, 0.7, 0.4, 0.6, 0.5, 0.8, 0.3, 0.9, 0.4, 0.6, 0.5, 0.7, 0.4, 0.6, 0.5,
}

var CoordNoiseLevels = []float64{
	0.0001, 0.0002, 0.00015, 0.0001, 0.00025, 0.0002, 0.0001, 0.0003, 0.00015, 0.0002,
	0.0001, 0.0002, 0.00025, 0.00015, 0.0002,
}

var MaxSpeeds = []float64{
	30.0, 25.0, 50.0, 45.0, 20.0,
	15.0, 35.0, 40.0, 10.0, 28.0,
	32.0, 22.0, 18.0, 40.0, 25.0,
}

var DriftRates = []float64{
	0.00001, 0.00002, 0.000015, 0.00001, 0.000025,
	0.00002, 0.00001, 0.00003, 0.000015, 0.00002,
	0.00001, 0.00002, 0.000025, 0.000015, 0.00002,
}

var Latitudes = []float64{
	19.0760, 28.7041, 12.9716, 17.3850, 13.0827,
	18.5204, 23.0225, 1.3521, 50.1109, 51.5074,
	40.7128, 37.7749, 35.6895, -33.8688, 43.6511,
}

var Longitudes = []float64{
	72.8777, 77.1025, 77.5946, 78.4867, 80.2707,
	73.8567, 72.5714, 103.8198, 8.6821, -0.1278,
	-74.0060, -122.4194, 139.6917, 151.2093, -79.3832,
}

var Altitudes = []float64{
	10.0, 15.0, 8.0, 12.0, 9.0,
	11.0, 7.0, 5.0, 20.0, 25.0,
	30.0, 18.0, 12.0, 6.0, 14.0,
}

var Speeds = []float64{
	0.0, 5.0, 12.0, 8.0, 3.0,
	10.0, 15.0, 7.0, 0.0, 20.0,
	25.0, 12.0, 5.0, 0.0, 8.0,
}

var Headings = []float64{
	0.0, 45.0, 90.0, 180.0, 270.0,
	30.0, 60.0, 120.0, 210.0, 300.0,
	15.0, 75.0, 135.0, 225.0, 315.0,
}

var IsMoving = []bool{
	false, true, true, true, false,
	true, true, true, false, true,
	true, true, false, false, true,
}
