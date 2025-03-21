package main

import (
	"net/http"
)

// @Summary Health check endpoint
// @Description Get the current status of the API
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//Create a map which holds the information that we want to send in the response.
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
