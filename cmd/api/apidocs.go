package main

import (
	_ "embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed dist/swagger.yaml
var apiDocsYAML []byte

//go:embed dist/index.html
var apiDocsIndex []byte

//go:embed dist/logo.png
var apiDocsLogo []byte

func (app *application) RouteAPIDocs(r chi.Router) {
	r.Get("/static/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		app.HandleStaticFile(w, r, apiDocsYAML, "text/yaml")
	})
	r.Get("/static/logo.png", func(w http.ResponseWriter, r *http.Request) {
		app.HandleStaticFile(w, r, apiDocsLogo, "image/png")
	})
	r.Get("/apidocs", func(w http.ResponseWriter, r *http.Request) {
		app.HandleStaticFile(w, r, apiDocsIndex, "text/html")
	})
}

// HandleStaticFile is a helper function to serve static files
func (app *application) HandleStaticFile(w http.ResponseWriter, r *http.Request, file []byte, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "no-store")

	_, err := w.Write(file)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
