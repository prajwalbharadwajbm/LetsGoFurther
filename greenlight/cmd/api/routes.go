package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/prajwalbharadwajbm/LetsGoFurther/greenlight/docs" // Import the generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// Set custom error handlers that adhere to http.Handler interface as required by httprouter.
	// These handlers will be called when a route is not found (404) or when a method is not allowed (405).
	// The handlers must implement ServeHTTP(ResponseWriter, *Request) as per http.Handler interface.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

	return router
}
