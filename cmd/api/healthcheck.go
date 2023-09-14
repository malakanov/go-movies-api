package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
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
