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
