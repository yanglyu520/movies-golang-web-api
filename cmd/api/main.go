package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverIdleTimeout  = time.Minute
	serverReadTimeout  = 5 * time.Second
	serverWriteTimeout = 10 * time.Second
)

type config struct {
	env  string
	port int
}
type application struct {
	cfg    config
	logger *log.Logger
}

//	@title			Movies Web API
//	@version		1.0
//	@description	This is a Movies Web API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://github.com/yanglyu520
//	@contact.email	yang.lyu@hotmail.com

//	@servers.url			/
//	@servers.description	API server behind the api docs
func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "dev", "Env(dev/stage/prod)")
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	flag.Parse()

	app := &application{
		cfg:    cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  serverIdleTimeout,
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
	}

	app.logger.Printf("starting %s env server on port %d", app.cfg.env, app.cfg.port)

	err := server.ListenAndServe()

	app.logger.Fatalf("server failed with err: %w", err)

}
