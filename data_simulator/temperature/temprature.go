package temperature

import (
	"iot/data_simulator/common"
	"time"

	"math/rand"

	"github.com/google/uuid"
)

type Temperature struct {
	common.Metrics
}

func (t *Temperature) GenerateData(temprature_chan chan common.Metrics) {
	staticnum := rand.Intn(len(DeviceIds))

	// static
	t.Id = uuid.New().String()
	t.DeviceId = DeviceIds[staticnum]
	t.DeviceName = DeviceName
	t.Location = Locations[staticnum]
	t.Model = Models[staticnum]
	t.Manufacturer = Manufacturers[staticnum]
	t.InstallDate = InstallDates[staticnum]
	t.Unit = Units[staticnum]

	// static per run
	t.BaselineTemp = BaselineTemps[rand.Intn(len(BaselineTemps))]
	t.SpikeProbability = SpikeProbabilities[rand.Intn(len(SpikeProbabilities))]
	t.SpikeMagnitude = SpikeMagnitudes[rand.Intn(len(SpikeMagnitudes))]
	t.NoiseLevel = NoiseLevels[rand.Intn(len(NoiseLevels))]
	t.UpdateInterval = time.Duration(rand.Intn(TempratureUpdateInterval))
	t.DriftRate = DriftRates[rand.Intn(len(DriftRates))]

	// dynamic
	t.CurrentTemp = CurrentTemps[rand.Intn(len(CurrentTemps))]
	switch {
	case t.CurrentTemp > t.BaselineTemp+1.0:
		t.Trend = "rising"
	case t.CurrentTemp < t.BaselineTemp-1.0:
		t.Trend = "falling"
	default:
		t.Trend = "stable"
	}
	if t.SpikeMagnitude > SpikeThreshold {
		t.IsSpiking = true
		t.LastSpikeTime = time.Now()
	}
	if time.Now().After(t.NextRead) {
		t.NextRead = time.Now().Add(t.UpdateInterval)
	}

	temprature_chan <- t.Metrics
}
