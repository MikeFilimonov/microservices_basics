package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		ErrorOccured: false,
		Payload:      "Generic message to the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}
