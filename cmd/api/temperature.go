package main

import (
	st "iot/internal/store"
	"net/http"
)

func (app *application) getTemperatureStatistics(w http.ResponseWriter, r *http.Request) {
	temperatureStatistics, err := app.store.Temperature.(*st.TemperatureStore).GetStatistics(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, temperatureStatistics); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getTemperatureAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	temperatureAggregation, err := app.store.Temperature.(*st.TemperatureStore).GetAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, temperatureAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getTemperatureAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	temperatureAggregation, err := app.store.Temperature.(*st.TemperatureStore).GetAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, temperatureAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getTemperatureDailyAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	temperatureDailyAggregation, err := app.store.Temperature.(*st.TemperatureStore).GetDailyAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, temperatureDailyAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
