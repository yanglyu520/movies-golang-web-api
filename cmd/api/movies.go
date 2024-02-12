package main

import (
	"github.com/yanglyu520/movies-golang-web-api/internal/data"
	"math/rand"
	"net/http"
	"time"
)

// HandleMovieList  is the handler for getting a list of movies endpoint
//
//	@Summary		Get a list of movies
//	@Description	Get a list of movies
//	@Tags			movies
//	@Produce		json
//	@Success		200	{object}	data.MovieList
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies [get]
func (app *application) HandleMovieList(w http.ResponseWriter, r *http.Request) {
	var err error

	movie := data.Movie{
		ID:        int64(rand.Int()),
		CreatedAt: time.Now().Unix(),
		Title:     "something",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	movieList := data.MovieList{[]data.Movie{movie}}

	err = app.writeJSON(w, http.StatusOK, movieList, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		http.Error(w, "the server has encountered a problem and could not process your request", http.StatusInternalServerError)

	}
}

type movieInput struct {
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime int32    `json:"runtime"`
	Genres  []string `json:"genres"`
}

// HandleMoviePost is the handler for creating a movie endpoint
//
//	@Summary		Create a new movie
//	@Description	Create a new movie
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies [post]
func (app *application) HandleMoviePost(w http.ResponseWriter, r *http.Request) {
	var err error
	var input movieInput

	err = app.readJSONInput(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	movie := data.Movie{
		ID:        int64(rand.Int()),
		CreatedAt: time.Now().Unix(),
		Title:     "something",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	movieWithEnvelop := map[string]any{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusOK, movieWithEnvelop, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		http.Error(w, "the server has encountered a problem and could not process your request", http.StatusInternalServerError)

	}
}

// HandleMovieGet  is the handler for getting a specific movie endpoint
//
//	@Summary		Get a specific movie
//	@Description	Get  a specific movie
//	@Tags			movies
//	@Param			id	path	string	true	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/id [get]
func (app *application) HandleMovieGet(w http.ResponseWriter, r *http.Request) {
	var err error

	var id int64
	id, err = app.readMovieIDParam(r)
	if err != nil {
		app.logger.Printf("failed to read the movie id with error: %s\n", err.Error())
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now().Unix(),
		Title:     "something",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	movieWithEnvelop := map[string]any{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusOK, movieWithEnvelop, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		http.Error(w, "the server has encountered a problem and could not process your request", http.StatusInternalServerError)

	}

}

// HandleMoviePut is the handler for the update a specific movie endpoint
//
//	@Summary		Update a specific movie
//	@Description	Update a specific movie
//	@Tags			movies
//	@Param			id	path	string	true	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/id [put]
func (app *application) HandleMoviePut(w http.ResponseWriter, r *http.Request) {
	var err error

	var id int64
	id, err = app.readMovieIDParam(r)
	if err != nil {
		app.logger.Printf("failed to read the movie id with error: %s\n", err.Error())
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now().Unix(),
		Title:     "something",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	movieWithEnvelop := map[string]any{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusOK, movieWithEnvelop, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		http.Error(w, "the server has encountered a problem and could not process your request", http.StatusInternalServerError)

	}
}

// HandleMovieDelete is the handler for the delete a specific movie endpoint
//
//	@Summary		Delete a specific movie
//	@Description	Delete a specific movie
//	@Tags			movies
//	@Param			id	path	string	true	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/id [delete]
func (app *application) HandleMovieDelete(w http.ResponseWriter, r *http.Request) {
	var err error

	var id int64
	id, err = app.readMovieIDParam(r)
	if err != nil {
		app.logger.Printf("failed to read the movie id with error: %s\n", err.Error())
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now().Unix(),
		Title:     "something",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	movieWithEnvelop := map[string]any{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusOK, movieWithEnvelop, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		http.Error(w, "the server has encountered a problem and could not process your request", http.StatusInternalServerError)

	}
}
