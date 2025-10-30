package main

import (
	st "iot/internal/store"
	"net/http"
)

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
