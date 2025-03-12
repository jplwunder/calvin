package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/healthcheck", app.healthcheckHandler)
	router.GET("/contacts", app.getContacts)
	router.POST("/contacts", app.createContact)
	router.GET("/contacts/:id", app.getContactByID)

	return app.recoverPanic(app.rateLimit(app.enableCORS(router)))
}
