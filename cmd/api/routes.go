package api

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) wrapHandler(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		paramsMap := make(map[string]string)
		for _, param := range ps {
			paramsMap[param.Key] = param.Value
		}

		ctx := context.WithValue(r.Context(), httprouterParamsKey, paramsMap)
		r = r.WithContext(ctx)

		h(w, r)
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/healthcheck", app.wrapHandler(app.healthcheckHandler))
	router.GET("/contacts", app.wrapHandler(app.listContactsHandler))
	router.POST("/contacts", app.wrapHandler(app.createContactHandler))
	router.GET("/contacts/:id", app.wrapHandler(app.getContactHandler))

	return app.recoverPanic(app.rateLimit(app.enableCORS(router)))
}
