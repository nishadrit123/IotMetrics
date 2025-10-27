package main

import (
	st "iot/internal/store"
	"net/http"
)

func (app *application) getCPUStatistics(w http.ResponseWriter, r *http.Request) {
	cpuStatistics, err := app.store.CPU.(*st.CPUStore).GetStatistics(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cpuStatistics); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getCPUAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	cpuAggregation, err := app.store.CPU.(*st.CPUStore).GetAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cpuAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getCPUAggregationPerModel(w http.ResponseWriter, r *http.Request) {
	cpuAggregation, err := app.store.CPU.(*st.CPUStore).GetAggregationPerModel(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cpuAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getCPUDailyAggregationPerLocation(w http.ResponseWriter, r *http.Request) {
	cpuDailyAggregation, err := app.store.CPU.(*st.CPUStore).GetDailyAggregationPerLocation(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cpuDailyAggregation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
