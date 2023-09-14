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

	//init server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
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
