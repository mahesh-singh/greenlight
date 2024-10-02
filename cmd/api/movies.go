package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movies")
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParm(r)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the detail of movies %d\n", id)
}
