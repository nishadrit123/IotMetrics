package main

import (
	st "iot/internal/store"
	"net/http"
)

func (app *application) getHumidityStatistics(w http.ResponseWriter, r *http.Request) {
	humidityStatistics, err := app.store.Humidity.(*st.HumidityStore).GetStatistics(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, humidityStatistics); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getHumidityAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	humidityAggregation, err := app.store.Humidity.(*st.HumidityStore).GetAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, humidityAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getHumidityAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	humidityAggregation, err := app.store.Humidity.(*st.HumidityStore).GetAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, humidityAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getHumidityDailyAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	humidityDailyAggregation, err := app.store.Humidity.(*st.HumidityStore).GetDailyAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, humidityDailyAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
