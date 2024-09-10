package main

import (
	"fmt"
	"greenlight/alan.com/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        1,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Year:      1942,
		Runtime:   102,
		Genres:    []string{"Drama", "Romance", "War"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
