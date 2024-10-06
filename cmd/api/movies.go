package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mahesh-singh/greenlight/internal/data"
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

	movies := data.Movies{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movies}, nil)

	if err != nil {
		app.logger.Error(err.Error())

		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
