package main

import (
	"context"
	_ "embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

//go:embed dist/swagger.yaml
var apiDocsYAML []byte

//go:embed dist/index.html
var apiDocsIndex []byte

//go:embed dist/logo.png
var apiDocsLogo []byte

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", app.HandleRootGet)

	// API version 1
	r.Route("/v1", func(r chi.Router) {
		r.Use(app.apiVersionCtx("v1"))
		r.Get("/", app.HandleRootGet)
		r.Get("/apidocs", func(w http.ResponseWriter, r *http.Request) {
			app.HandleStaticFile(w, r, apiDocsIndex, "text/html")
		})

		r.Mount("/movies", app.movieRouter())
		r.Get("/static/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
			app.HandleStaticFile(w, r, apiDocsYAML, "text/yaml")
		})
		r.Get("/static/logo.png", func(w http.ResponseWriter, r *http.Request) {
			app.HandleStaticFile(w, r, apiDocsLogo, "image/png")
		})
	})

	return r
}

func (app *application) apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func (app *application) movieRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", app.HandleMovieList)
	r.Post("/", app.HandleMoviePost)
	r.Route("/{movieID}", func(r chi.Router) {
		r.Get("/", app.HandleMovieGet)
		r.Put("/", app.HandleMoviePut)
		r.Delete("/", app.HandleMovieDelete)

	})
	return r
}

//func (app *application) staticFileRouter() http.Handler {
//	r := chi.NewRouter()
//
//	r.Get("/logo.png", func(w http.ResponseWriter, r *http.Request) {
//		app.HandleStaticFile(w, r, apiDocsLogo, "image/png")
//	})
//	r.Get("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
//		app.HandleStaticFile(w, r, apiDocsYAML, "text/yaml")
//	})
//	return r
//}

func (app *application) HandleRootGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte(`movies Web API, see <a href="/v1/apidocs">API Docs</a> for documentation.`))
	if err != nil {
		app.handleServerError(w, r, err)
		return
	}
}
