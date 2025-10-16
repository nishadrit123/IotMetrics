package pressure

import (
	"iot/data_simulator/common"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Pressure struct {
	common.Metrics
}

func (p *Pressure) GenerateData(pressure_chan chan common.Metrics) {
	staticNum := rand.Intn(len(DeviceIds))

	p.Id = uuid.New().String()
	p.DeviceId = DeviceIds[staticNum]
	p.DeviceName = DeviceName
	p.Location = Locations[staticNum]
	p.Model = Models[staticNum]
	p.Manufacturer = Manufacturers[staticNum]
	p.InstallDate = InstallDates[staticNum]
	p.Unit = Units[staticNum]

	p.BaselinePressure = BaselinePressure[rand.Intn(len(BaselinePressure))]
	p.SpikeProbability = SpikeProbabilities[rand.Intn(len(SpikeProbabilities))]
	p.SpikeMagnitude = SpikeMagnitudes[rand.Intn(len(SpikeMagnitudes))]
	p.NoiseLevel = NoiseLevels[rand.Intn(len(NoiseLevels))]
	p.DriftRate = DriftRates[rand.Intn(len(DriftRates))]
	p.UpdateInterval = time.Duration(rand.Intn(PressureUpdateInterval))

	p.CurrentPressure = CurrentPressure[rand.Intn(len(CurrentPressure))]

	if p.SpikeMagnitude > SpikeThreshold {
		p.IsSpiking = true
		p.LastSpikeTime = time.Now()
	}

	if time.Now().After(p.NextRead) {
		p.NextRead = time.Now().Add(p.UpdateInterval)
	}

	switch {
	case p.CurrentPressure > p.BaselinePressure+5.0:
		p.Trend = "rising"
	case p.CurrentPressure < p.BaselinePressure-5.0:
		p.Trend = "falling"
	default:
		p.Trend = "stable"
	}

	pressure_chan <- p.Metrics
}
