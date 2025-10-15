package cpu

const (
	SpikeTreshold     = 30.0
	CPUUpdateInterval = 3
	DeviceName        = "cpu"
)

var CPUId = []string{
	"CPU-9f3a1c72", "CPU-47b8d2ef", "CPU-a13f5e90", "CPU-6b2c4d11", "CPU-f09a7b34",
	"CPU-3e7c9a52", "CPU-b4f1d689", "CPU-28d6a9c3", "CPU-c8e2f934", "CPU-5a7b3d80",
	"CPU-d49f8a15", "CPU-7c2d1e63", "CPU-fb8a6d27", "CPU-9a3e2b74", "CPU-1f6d4e58",
}

var CPUbaselineUsage = []float64{
	12.5, 20.0, 35.0, 10.0, 50.0, 25.0, 18.0, 40.0, 60.0, 28.0,
}

var CPUSpikeProbability = []float64{
	0.10, 0.15, 0.20, 0.05, 0.30, 0.12, 0.08, 0.25, 0.35, 0.18,
}

var CPUSpikeMagnitude = []float64{
	25.0, 30.0, 45.0, 15.0, 50.0, 20.0, 22.0, 35.0, 55.0, 28.0,
}

var CPUNoiseLevel = []float64{
	3.0, 4.0, 5.0, 2.0, 6.0, 3.5, 2.8, 5.5, 7.0, 4.2,
}

var CPUCurrentUsage = []float64{
	14.2, 22.1, 37.6, 11.3, 54.5, 27.8, 19.4, 44.1, 63.9, 30.6,
}

var CPUTemperature = []float64{
	45.5, 50.2, 58.8, 42.0, 63.4, 49.1, 46.3, 60.7, 68.2, 52.9,
}
