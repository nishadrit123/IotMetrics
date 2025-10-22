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
