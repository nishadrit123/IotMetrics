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

var CPUName = []string{
	"alpha-core-01", "alpha-core-02", "beta-node-01", "beta-node-02", "gamma-cpu-01", "gamma-cpu-02", "delta-engine-01", "delta-engine-02", "epsilon-core-01", "zeta-node-01", "eta-cpu-01", "theta-core-01", "iota-engine-01", "kappa-node-01", "lambda-cpu-01",
}

var CPULocation = []string{
	"Mumbai, India", "Delhi, India", "Bangalore, India", "Hyderabad, India", "Chennai, India",
	"Pune, India", "Ahmedabad, India", "Singapore", "Frankfurt, Germany", "London, UK",
	"New York, USA", "San Francisco, USA", "Tokyo, Japan", "Sydney, Australia", "Toronto, Canada",
}

var CPUModels = []string{
	"Intel Xeon E5-2690 v4", "Intel Xeon Gold 6230", "AMD EPYC 7742", "Intel Core i9-12900K",
	"AMD Ryzen 9 7950X", "Intel Xeon Platinum 8280", "AMD EPYC 7302P", "Intel Core i7-13700K",
	"Intel Xeon Silver 4214", "AMD Ryzen 7 5800X3D", "Intel Core i5-12600K", "Apple M2 Ultra",
	"AMD EPYC 9654", "Intel Xeon D-2183IT", "AMD Ryzen Threadripper PRO 5995WX",
}

var CPUCore = []int{
	4, 6, 8, 12, 16, 10, 8, 6, 4, 12, 8, 14, 6, 8, 16,
}

var CPUFrequency = []float64{
	2.4, 3.1, 2.9, 3.5, 3.8, 2.6, 2.8, 3.3, 3.0, 2.5, 3.7, 4.0, 2.2, 3.2, 3.9,
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
