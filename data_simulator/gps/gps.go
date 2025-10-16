package gps

import (
	"iot/data_simulator/common"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type GPS struct {
	common.Metrics
}

func (g *GPS) GenerateData(gps_chan chan common.Metrics) {
	staticNum := rand.Intn(len(DeviceIds))

	// static
	g.Id = uuid.New().String()
	g.DeviceId = DeviceIds[staticNum]
	g.DeviceName = DeviceName
	g.Location = Locations[staticNum]
	g.Model = Models[staticNum]
	g.Manufacturer = Manufacturers[staticNum]
	g.InstallDate = InstallDates[staticNum]

	// static per run
	g.UpdateInterval = time.Duration(rand.Intn(GPSUpdateInterval))
	speedNoise := SpeedNoiseLevels[rand.Intn(len(SpeedNoiseLevels))]
	coordNoise := CoordNoiseLevels[rand.Intn(len(CoordNoiseLevels))]
	maxSpeed := MaxSpeeds[rand.Intn(len(MaxSpeeds))]
	drift := DriftRates[rand.Intn(len(DriftRates))]
	g.Latitude = Latitudes[rand.Intn(len(Latitudes))]
	g.Longitude = Longitudes[rand.Intn(len(Longitudes))]
	g.Altitude = Altitudes[rand.Intn(len(Altitudes))]
	g.Speed = Speeds[rand.Intn(len(Speeds))]
	g.Heading = Headings[rand.Intn(len(Headings))]
	g.IsMoving = IsMoving[rand.Intn(len(IsMoving))]

	g.Latitude += (rand.Float64()*2 - 1) * coordNoise
	g.Longitude += (rand.Float64()*2 - 1) * coordNoise

	g.Speed += (rand.Float64()*2 - 1) * speedNoise
	if g.Speed < 0 {
		g.Speed = 0
	} else if g.Speed > maxSpeed {
		g.Speed = maxSpeed
	}

	g.Latitude += (rand.Float64()*2 - 1) * drift
	g.Longitude += (rand.Float64()*2 - 1) * drift

	// Update NextRead timestamp
	if time.Now().After(g.NextRead) {
		g.NextRead = time.Now().Add(g.UpdateInterval)
	}

	gps_chan <- g.Metrics
}
