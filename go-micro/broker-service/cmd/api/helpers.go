package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonResponse struct {
	ErrorOccured bool   `json:"error"`
	Payload      string `json:"payload"`
	Hint         any    `json:"hint,omitempty"`
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {

	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.ErrorOccured = true
	payload.Payload = err.Error()

	return app.writeJSON(w, statusCode, payload)
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Response, data any) error {

	maxSize := 1048576 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSize))

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("body must have only a single JSON value")
	}

	return nil

}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set(HeaderContentType, JSONContentType)
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil

}
