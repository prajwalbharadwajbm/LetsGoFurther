package main

import (
	"net/http"
)

// healthcheckHandler is a HTTP handler
// returns a 200 OK response and prints message which includes the status, environment and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//Create a map which holds the information that we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}
