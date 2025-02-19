package main

import (
	"fmt"
	"net/http"
)

// healthcheckHandler is a HTTP handler
// returns a 200 OK response and prints message which includes the status, environment and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version":%q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("charset", "utf-8")

	w.Write([]byte(js))
}
