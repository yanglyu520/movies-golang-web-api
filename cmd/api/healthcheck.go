package main

import (
	"net/http"

	"github.com/yanglyu520/movies-golang-web-api/internal/utils/debugutils"
)

const version = "1.0.0"

type systemInfo struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

type envStruct struct {
	Status     string     `json:"status"`
	SystemInfo systemInfo `json:"system_info"`
	BuiltSHA   string     `json:"built_sha"`
}

type envWithEnvelop struct {
	Env envStruct `json:"env"`
}

// HandleHealthCheck is the handler for the healthcheck endpoint
//
//	@Summary		Healthcheck
//	@Description	Returns a healthcheck response including the Git SHA that was
//	@Description	used to build the current binary. This does not use the HASH env
//	@Description	variable but rather the binary debug symbols.
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}	envWithEnvelop
//	@Failure		500	{object}	error
//	@Router			/v1/healthcheck [get]
func (app *application) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	env := envStruct{
		Status: "available",
		SystemInfo: systemInfo{
			Env:     app.cfg.env,
			Version: version,
		},
		BuiltSHA: debugutils.CommitSHA(),
	}

	envWithEnvelop := envWithEnvelop{env}

	err := app.writeJSON(w, http.StatusOK, envWithEnvelop, nil)
	if err != nil {
		app.logError(r, err)
	}

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
