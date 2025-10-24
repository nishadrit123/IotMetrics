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
