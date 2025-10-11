package cpu

import (
	"encoding/json"
	"fmt"
	"iot/data_simulator/common"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CPU struct {
	common.Metrics
}

func (c *CPU) GenerateData() {
	staticnum := rand.Intn(len(CPUId))

	// static
	c.Id = uuid.New().String()
	c.DeviceId = CPUId[staticnum]
	c.DeviceName = DeviceName
	c.HostName = CPUName[staticnum]
	c.Location = CPULocation[staticnum]
	c.Model = CPUModels[staticnum]
	c.CoreCount = CPUCore[staticnum]
	c.Frequency = CPUFrequency[staticnum]

	// static per run
	c.BaselineUsage = CPUbaselineUsage[rand.Intn(len(CPUbaselineUsage))]
	c.SpikeProbability = CPUSpikeProbability[rand.Intn(len(CPUSpikeProbability))]
	c.SpikeMagnitude = CPUSpikeMagnitude[rand.Intn(len(CPUSpikeMagnitude))]
	c.NoiseLevel = CPUNoiseLevel[rand.Intn(len(CPUNoiseLevel))]
	c.UpdateInterval = time.Duration(rand.Intn(CPUUpdateInterval))

	// dynamic
	c.CurrentUsage = CPUCurrentUsage[rand.Intn(len(CPUCurrentUsage))]
	c.Temperature = CPUTemperature[rand.Intn(len(CPUTemperature))]
	if c.SpikeMagnitude > SpikeTreshold {
		c.IsSpiking = true
		c.LastSpikeTime = time.Now()
	}
	if time.Now().After(c.NextRead) {
		c.NextRead = time.Now().Add(c.UpdateInterval)
	}
	str, _ := json.Marshal(c)
	fmt.Printf("cpu: %v\n\n", string(str))
}
