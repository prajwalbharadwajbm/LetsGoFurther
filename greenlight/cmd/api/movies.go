package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prajwalbharadwajbm/LetsGoFurther/greenlight/internal/data"
)

// @Summary Create a new movie
// @Description Create a new movie with the input payload
// @Tags movies
// @Accept json
// @Produce json
// @Success 201 {object} Movie
// @Failure 400 {object} map[string]interface{}
// @Router /v1/movies [post]
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Struct fields must start with capital letter(It must be exported), so that they're visible to encoding/josn package.
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// must use a non-nil pointer for Decode(), else we find json.InvalidUnmarshalError error at runtime.
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

// @Summary Get a movie by ID
// @Description Get details of a specific movie by its ID
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} Movie
// @Failure 404 {object} map[string]interface{}
// @Router /v1/movies/{id} [get]
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "casablanca",
		Year:      2001,
		Runtime:   145,
		Genres:    []string{"comedy", "romance"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
