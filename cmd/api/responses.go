package main

import (
	"fmt"
	"net/http"
)

// send server error 500, log the error and send error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.genericErrorResponse(w, r, http.StatusInternalServerError, errServerMessage)
}

// send 404 not found error, log the error and send error response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.genericErrorResponse(w, r, http.StatusNotFound, errNotFoundMessage)
}

// send 405 method not allowed error, log the error and send error response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	app.genericErrorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("method %s %s", r.Method, errMessageNotAllowed))
}

// send 400 Bad Request error
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request) {
	app.genericErrorResponse(w, r, http.StatusBadRequest, errNotFoundMessage)
}

// send 422 unprocessable entity
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.genericErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
