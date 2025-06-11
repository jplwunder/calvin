package api

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Default database status
	db_status := "available"

	// Create a context with a 5-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext to test the database connection.
	err := app.db.Ping(ctx)
	if err != nil {
		db_status = "unavailable"
	}

	env := envelope{
		"status": "available",
		"system_info": map[string]any{ // Changed to map[string]any
			"environment":     app.config.environment,
			"version":         version,
			"trusted_origins": strings.Join(app.config.cors.trustedOrigins, ", "),
			"database": map[string]string{
				"status": db_status,
			},
		},
	}

	err = app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
