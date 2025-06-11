package api

import (
	"net/http"
	"strings"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment":     app.config.environment,
			"version":         version,
			"trusted_origins": strings.Join(app.config.cors.trustedOrigins, ", "),
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
