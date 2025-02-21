package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prajwalbharadwajbm/LetsGoFurther/greenlight/internal/data"
)

// Movie represents a movie in the system
// @Description Movie information
type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   int32     `json:"runtime"`
	CreatedAt time.Time `json:"created_at"`
}

// @Summary Create a new movie
// @Description Create a new movie with the input payload
// @Tags movies
// @Accept json
// @Produce json
// @Success 201 {object} Movie
// @Failure 400 {object} map[string]interface{}
// @Router /v1/movies [post]
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new Movie")
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
		http.NotFound(w, r)
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

	app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)

	fmt.Fprintf(w, "show the details of movie with id: %d\n", id)
}
