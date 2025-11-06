package main

import (
	"encoding/json"
	st "iot/internal/store"
	"log"
	"net/http"
)

type Locations struct {
	Locs []string `json:"locs,omitempty"`
}

func (app *application) getHeatMap(w http.ResponseWriter, r *http.Request) {
	var payload Locations
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	strLocs, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json payload for heatmap %v\n", err)
	}

	gpsDelta, err := app.store.HeatMap.(*st.HeatMapStore).GetHeatMap(strLocs)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, gpsDelta); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
