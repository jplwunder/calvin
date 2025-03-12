package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment":     app.config.env,
			"version":         version,
			"trusted_origins": strings.Join(app.config.cors.trustedOrigins, ", "),
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
