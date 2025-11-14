package common

import "time"

// Metrics defines common fields for different types of metrics
// that can be embedded in specific metric structs like
// CPU, Temperature, Humidity, Pressure and GPS

type Metrics struct {
	Id                string    `json:"id,omitempty"`
	DeviceId          string    `json:"device_id,omitempty"`
	Loc               string    `json:"loc,omitempty"`
	Model             string    `json:"model,omitempty"`
	DeviceName        string    `json:"device_name,omitempty"`
	HostName          string    `json:"hostname,omitempty"`
	CoreCount         int64     `json:"core_count,omitempty"`
	Frequency         float64   `json:"frequency,omitempty"`
	Manufacturer      string    `json:"manufacturer,omitempty"`
	CountManufacturer uint64    `json:"countManufacturer,omitempty"`
	InstallDate       time.Time `json:"install_date,omitempty"`
	CountRecords      uint64    `json:"countRecords,omitempty"`
	UniqFrequency     uint64    `json:"uniqFrequency,omitempty"`
	CountNoiseLevel   uint64    `json:"countNoiseLevel,omitempty"`
	Day               string    `json:"day,omitempty"`

	BaselineUsage       float64       `json:"baseline_usage,omitempty"`
	BaselineHumidity    float64       `json:"baseline_humidity,omitempty"`
	SumBaselineHumidity float64       `json:"sumBaselineHumidity,omitempty"`
	BaselineTemp        float64       `json:"baseline_temperature,omitempty"`
	SumBaselineTemp     float64       `json:"sumBaselineTemperature,omitempty"`
	BaselinePressure    float64       `json:"baseline_pressure,omitempty"`
	SumBaselinePressure float64       `json:"sumBaselinePressure,omitempty"`
	SpikeProbability    float64       `json:"spike_probability,omitempty"`
	SpikeMagnitude      float64       `json:"spike_magnitude,omitempty"`
	MaxSpikeMagnitude   float64       `json:"maxSpikeMagnitude,omitempty"`
	NoiseLevel          float64       `json:"noise_level,omitempty"`
	UpdateInterval      time.Duration `json:"update_interval_seconds,omitempty"`
	DriftRate           float64       `json:"drift_rate,omitempty"`
	SumDriftRate        float64       `json:"sumDriftRate,omitempty"`
	MinDriftRate        float64       `json:"minDriftRate,omitempty"`
	Trend               string        `json:"trend,omitempty"`
	CountTrend          uint64        `json:"count_trend,omitempty"`
	UniqTrend           uint64        `json:"uniqTrend,omitempty"`
	Latitude            float64       `json:"latitude,omitempty"`
	AvgLatitude         float64       `json:"avgLatitude,omitempty"`
	Longitude           float64       `json:"longitude,omitempty"`
	MaxLongitude        float64       `json:"maxLongitude,omitempty"`
	Altitude            float64       `json:"altitude,omitempty"`
	MaxAltitude         float64       `json:"maxAltitude,omitempty"`
	Speed               float64       `json:"speed,omitempty"`
	AvgSpeed            float64       `json:"avgSpeed,omitempty"`
	Heading             float64       `json:"heading,omitempty"`
	MaxHeading          float64       `json:"maxHeading,omitempty"`
	IsMoving            bool          `json:"is_moving,omitempty"`
	RollingAverage      float64       `json:"rolling_avg,omitempty"`
	Delta               float64       `json:"delta,omitempty"`

	CurrentUsage          float64   `json:"current_usage,omitempty"`
	AvgCurrentUsage       float64   `json:"avgCurrentUsage,omitempty"`
	CurrentTemp           float64   `json:"current_temperature,omitempty"`
	CurrentHumidity       float64   `json:"current_humidity,omitempty"`
	AvgCurrentHumidity    float64   `json:"avgCurrentHumidity,omitempty"`
	CurrentPressure       float64   `json:"current_pressure,omitempty"`
	AvgCurrentPressure    float64   `json:"avgCurrentPressure,omitempty"`
	Temperature           float64   `json:"cpu_temperature,omitempty"`
	AvgCurrentTemperature float64   `json:"avgCurrentTemperature,omitempty"`
	TotalCPUTemperature   float64   `json:"totalCPUTemperature,omitempty"`
	AvgCPUTemperature     float64   `json:"avgCPUTemperature,omitempty"`
	IsSpiking             bool      `json:"is_spiking,omitempty"`
	LastSpikeTime         time.Time `json:"last_spike_time,omitempty"`
	NextRead              time.Time `json:"next_read_time,omitempty"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
}
