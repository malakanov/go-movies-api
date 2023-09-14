package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show a movie")
}
