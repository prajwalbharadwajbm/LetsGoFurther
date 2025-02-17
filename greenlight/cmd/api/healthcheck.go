package main

import (
	"fmt"
	"net/http"
)

// healthcheckHandler is a HTTP handler
// returns a 200 OK response and prints message which includes the status, environment and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
