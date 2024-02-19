package main

import (
	_ "embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.MethodNotAllowed(app.methodNotAllowedResponse)
	r.NotFound(app.notFoundResponse)
	r.Use(middleware.Timeout(60 * time.Second))

	printRoutes(r)

	r.Get("/", app.HandleRootGet)

	r.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/", app.HandleRootGet)
		r.Get("/healthcheck", app.handleHealthCheck)
		r.Mount("/movies", app.movieRouter())
		app.RouteAPIDocs(r)
	})

	return r
}

func (app *application) movieRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", app.HandleMovieList)
	r.Post("/", app.HandleMoviePost)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", app.HandleMovieGet)
		r.Put("/", app.HandleMoviePut)
		r.Delete("/", app.HandleMovieDelete)
	})
	return r
}

func (app *application) HandleRootGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte(`movies Web API, see <a href="/v1/apidocs">API Docs</a> for documentation.`))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func printRoutes(r chi.Routes) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		slog.Info("registered route", "method", method, "route", route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		slog.Info("chi.Walk failed", "error", err)
	}
}
