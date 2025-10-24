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
