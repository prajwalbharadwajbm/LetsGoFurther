package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// version is the current version of the application
// for now its a hard-coded global constant
const version = "1.0.0"

// config struct to hold all the configuration settings for our application
type config struct {
	port int
	env  string
}

// appication struct is used for all http handlers
// it will contain the dependencies for each route
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}
	// Creates a  servemux and add a /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler method.
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Creates an HTTP server with timeout settings
	// Listens on port provided in the config
	// Handles incoming requests using the ServeMux declared above
	// IdleTimeout is the maximum amount of time to wait for the next request when a connection is idle
	// ReadTimeout is the maximum duration for reading the entire request, including the body
	// WriteTimeout is the maximum duration before timing out the request write operation
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start the HTTP server
	logger.Printf("starting %s server on http://localhost%s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
