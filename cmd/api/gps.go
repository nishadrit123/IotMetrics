package main

import (
	"encoding/json"
	st "iot/internal/store"
	"log"
	"net/http"
)

type Delta struct {
	Preceding int `json:"preceding,omitempty"`
	Following int `json:"following,omitempty"`
}

func (app *application) getGPSStatistics(w http.ResponseWriter, r *http.Request) {
	gpsStatistics, err := app.store.GPS.(*st.GPSStore).GetStatistics(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsStatistics); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getGPSAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	gpsAggregation, err := app.store.GPS.(*st.GPSStore).GetAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getGPSAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	gpsAggregation, err := app.store.GPS.(*st.GPSStore).GetAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getGPSDailyAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	gpsDailyAggregation, err := app.store.GPS.(*st.GPSStore).GetDailyAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsDailyAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getGPSDelta(w http.ResponseWriter, r *http.Request) {
	var payload Delta
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	strDelta, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json payload %v\n", err)
	}

	gpsDelta, err := app.store.GPS.(*st.GPSStore).GetDelta(r, strDelta)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsDelta); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
