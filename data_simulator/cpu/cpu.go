package cpu

import (
	"encoding/json"
	"iot/data_simulator/common"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CPU struct {
	common.Metrics
}

func (c *CPU) GenerateData(cpu_chan chan string) {
	staticnum := rand.Intn(len(CPUId))

	// static
	c.Id = uuid.New().String()
	c.DeviceId = CPUId[staticnum]
	c.DeviceName = DeviceName

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

	cpuData, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error marshaling CPU data: %v", err)

	}
	cpu_chan <- string(cpuData)
	// fmt.Printf("cpu: %v\n\n", string(cpuData))
}
