package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yanglyu520/movies-golang-web-api/internal/validator"

	"github.com/yanglyu520/movies-golang-web-api/internal/data"
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

	var input struct {
		Title  string
		Genres []string
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()
	input.Title = app.readString(qs, "title", "")
	input.Genres = app.readCSV(qs, "genres", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movies, metadata, err := app.models.Movies.GetAll(input.Title, input.Genres, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, data.MovieList{Movies: movies, Metadata: metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
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
		app.logError(r, err)
		app.badRequestResponse(w, r)
		return
	}

	log.Println(input)

	movie := &data.Movie{
		Title:     input.Title,
		Year:      input.Year,
		Genres:    input.Genres,
		Runtime:   input.Runtime,
		CreatedAt: time.Now().Unix(),
	}

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	movieWithEnvelop := map[string]any{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusCreated, movieWithEnvelop, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// HandleMovieGet  is the handler for getting a specific movie endpoint
//
//	@Summary		Get a specific movie
//	@Description	Get  a specific movie
//	@Tags			movies
//	@Param			id	path	string	false	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/{id} [get]
func (app *application) HandleMovieGet(w http.ResponseWriter, r *http.Request) {
	var err error

	var id int64
	id, err = app.readMovieIDParam(r)
	if err != nil {
		app.logError(r, err)
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}

// HandleMoviePut is the handler for the update a specific movie endpoint
//
//	@Summary		Update a specific movie
//	@Description	Update a specific movie
//	@Tags			movies
//	@Param			id	path	string	false	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/{id} [put]
func (app *application) HandleMoviePut(w http.ResponseWriter, r *http.Request) {
	var err error

	var id int64
	id, err = app.readMovieIDParam(r)
	if err != nil {
		app.logError(r, err)
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
		app.serverErrorResponse(w, r, err)
	}
}

// HandleMovieDelete is the handler for the delete a specific movie endpoint
//
//	@Summary		Delete a specific movie
//	@Description	Delete a specific movie
//	@Tags			movies
//	@Param			id	path	string	false	"movie ID"
//	@Produce		json
//	@Success		200	{object}	data.Movie
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/v1/movies/{id} [delete]
func (app *application) HandleMovieDelete(w http.ResponseWriter, r *http.Request) {
	var err error

	_, err = app.readMovieIDParam(r)
	if err != nil {
		app.logError(r, err)
		app.notFoundResponse(w, r)
	}

	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
