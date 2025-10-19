package humidity

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
