package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// app version number
const version = "1.0.0"

// configuration settings of application
type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func (app application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "status: available")
	if err != nil {
		app.logger.Println(err)
	}

	_, err = fmt.Fprintf(w, "environment: %s\n", app.config.env)
	if err != nil {
		app.logger.Println(err)
	}

	_, err = fmt.Fprintf(w, "version: %s\n", version)
	if err != nil {
		app.logger.Println(err)
	}
}

func main() {
	// init config
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.StringVar(&cfg.env, "env", "development", "env (development|staging|production)")
	flag.Parse()

	// init logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// init application
	app := &application{
		config: cfg,
		logger: logger,
	}

	//init mux
	mux := http.NewServeMux()

	//handlers
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//init server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	//start server
	logger.Printf("starting %s server %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()

	if err != nil {
		logger.Fatalf("%s", err)
	}

}
