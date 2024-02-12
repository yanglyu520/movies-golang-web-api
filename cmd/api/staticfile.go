package main

import "net/http"

func (app *application) HandleStaticFile(w http.ResponseWriter, r *http.Request, file []byte, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "no-store")

	_, err := w.Write(file)
	if err != nil {
		app.handleServerError(w, r, err)
	}
}
