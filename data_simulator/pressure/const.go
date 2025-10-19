package pressure

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
