package main

import (
	st "iot/internal/store"
	"net/http"
)

func (app *application) getPressureStatistics(w http.ResponseWriter, r *http.Request) {
	pressureStatistics, err := app.store.Pressure.(*st.PressureStore).GetStatistics(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, pressureStatistics); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPressureAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	pressureAggregation, err := app.store.Pressure.(*st.PressureStore).GetAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, pressureAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPressureAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	pressureAggregation, err := app.store.Pressure.(*st.PressureStore).GetAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, pressureAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPressureDailyAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	pressureDailyAggregation, err := app.store.Pressure.(*st.PressureStore).GetDailyAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, pressureDailyAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
