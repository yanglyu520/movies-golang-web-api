package main

import (
	"net/http"
)

const (
	errServerMessage     = "the server encountered a problem and could not process your request"
	errNotFoundMessage   = "the requested resource could not be found"
	errMessageNotAllowed = "method is not supported for this resource"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err.Error())
}

type envelop map[string]any

func (app *application) genericErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelop{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}
