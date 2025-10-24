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
