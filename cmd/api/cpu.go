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
