package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/prajwalbharadwajbm/LetsGoFurther/greenlight/docs" // Import the generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

	return router
}
