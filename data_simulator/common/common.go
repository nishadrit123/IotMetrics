package common

import "time"

// Metrics defines common fields for different types of metrics
// that can be embedded in specific metric structs like
// CPU, Temperature, Humidity, Pressure and GPS

type Metrics struct {
	Id           string    `json:"id,omitempty"`
	DeviceId     string    `json:"device_id,omitempty"`
	Loc          string    `json:"loc,omitempty"`
	Model        string    `json:"model,omitempty"`
	DeviceName   string    `json:"device_name,omitempty"`
	HostName     string    `json:"host_name,omitempty"`
	CoreCount    int64     `json:"core_count,omitempty"`
	Frequency    float64   `json:"frequency_ghz,omitempty"`
	Manufacturer string    `json:"manufacturer,omitempty"`
	InstallDate  time.Time `json:"install_date,omitempty"`
	Count        uint64    `json:"count,omitempty"`
	CountNoise   uint64    `json:"count_noise,omitempty"`
	Day          string    `json:"day,omitempty"`

	BaselineUsage    float64       `json:"baseline_usage,omitempty"`
	BaselineHumidity float64       `json:"baseline_humidity,omitempty"`
	BaselineTemp     float64       `json:"baseline_temperature,omitempty"`
	BaselinePressure float64       `json:"baseline_pressure,omitempty"`
	SpikeProbability float64       `json:"spike_probability,omitempty"`
	SpikeMagnitude   float64       `json:"spike_magnitude,omitempty"`
	NoiseLevel       float64       `json:"noise_level,omitempty"`
	UpdateInterval   time.Duration `json:"update_interval_seconds,omitempty"`
	DriftRate        float64       `json:"drift_rate,omitempty"`
	Trend            string        `json:"trend,omitempty"`
	Latitude         float64       `json:"latitude,omitempty"`
	Longitude        float64       `json:"longitude,omitempty"`
	Altitude         float64       `json:"altitude_meters,omitempty"`
	Speed            float64       `json:"speed_kmh,omitempty"`
	Heading          float64       `json:"heading_degrees,omitempty"`
	IsMoving         bool          `json:"is_moving,omitempty"`

	CurrentUsage    float64   `json:"current_usage,omitempty"`
	CurrentTemp     float64   `json:"current_temperature,omitempty"`
	CurrentHumidity float64   `json:"current_humidity,omitempty"`
	CurrentPressure float64   `json:"current_pressure,omitempty"`
	Temperature     float64   `json:"temperature_celsius,omitempty"`
	IsSpiking       bool      `json:"is_spiking,omitempty"`
	LastSpikeTime   time.Time `json:"last_spike_time,omitempty"`
	NextRead        time.Time `json:"next_read_time,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
