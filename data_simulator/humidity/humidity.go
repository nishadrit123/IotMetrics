package humidity

import (
	"encoding/json"
	"fmt"
	"iot/data_simulator/common"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Humidity struct {
	common.Metrics
}

func (h *Humidity) GenerateData(humidity_chan chan string) {
	staticNum := rand.Intn(len(DeviceIds))

	h.Id = uuid.New().String()
	h.DeviceId = DeviceIds[staticNum]
	h.DeviceName = DeviceName
	h.Location = Locations[staticNum]
	h.Model = Models[staticNum]
	h.Manufacturer = Manufacturers[staticNum]
	h.InstallDate = InstallDates[staticNum]
	h.Unit = Units[staticNum]

	h.BaselineHumidity = BaselineHumidity[rand.Intn(len(BaselineHumidity))]
	h.SpikeProbability = SpikeProbabilities[rand.Intn(len(SpikeProbabilities))]
	h.SpikeMagnitude = SpikeMagnitudes[rand.Intn(len(SpikeMagnitudes))]
	h.NoiseLevel = NoiseLevels[rand.Intn(len(NoiseLevels))]
	h.DriftRate = DriftRates[rand.Intn(len(DriftRates))]
	h.UpdateInterval = time.Duration(rand.Intn(HumidityUpdateInterval))

	h.CurrentHumidity = CurrentHumidity[rand.Intn(len(CurrentHumidity))]

	if h.SpikeMagnitude > SpikeThreshold {
		h.IsSpiking = true
		h.LastSpikeTime = time.Now()
	}

	if time.Now().After(h.NextRead) {
		h.NextRead = time.Now().Add(h.UpdateInterval)
	}

	switch {
	case h.CurrentHumidity > h.BaselineHumidity+2:
		h.Trend = "rising"
	case h.CurrentHumidity < h.BaselineHumidity-2:
		h.Trend = "falling"
	default:
		h.Trend = "stable"
	}

	humidityData, err := json.Marshal(h)
	if err != nil {
		fmt.Printf("Error marshaling humidity data: %v", err)
		return
	}
	humidity_chan <- string(humidityData)
	// fmt.Printf("humidity: %v\n\n", string(humidityData))
}
