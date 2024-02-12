package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	errServerMessage     = "the server encountered a problem and could not process your request"
	errNotFoundMessage   = "the requested resource could not be found"
	errMessageNotAllowed = "method is not supported for this resource"
)

// APIError
type APIError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	CreatedAt    int64  `json:"created_at"`
}

//type APIErrorList struct {
//	Errors []APIError `json:"errors"`
//}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	env := APIError{
		status,
		message,
		time.Now().Unix(),
	}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// send server error 500, log the error and send error response
func (app *application) handleServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(w, r, http.StatusInternalServerError, errServerMessage)
}

// send 400 Bad Request error
func (app *application) handleBadRequest(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusBadRequest, errNotFoundMessage)
}

// send 404 not found error, log the error and send error response
func (app *application) handleNotFound(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, errNotFoundMessage)
}

// send 405 method not allowed error, log the error and send error response
func (app *application) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("method %s %s", r.Method, errMessageNotAllowed))
}
