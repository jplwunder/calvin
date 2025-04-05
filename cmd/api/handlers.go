package api

import (
	"calvin/internal/data"
	"net/http"
)

type contextKey string

const httprouterParamsKey = contextKey("httprouterParams")

func (app *application) listCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers := app.models.Customers.GetAll()

	err := app.writeJSON(w, http.StatusOK, envelope{"data": customers}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var newCustomer data.CustomerModel

	err := app.readJSON(w, r, &newCustomer)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.models.Customers.Insert(newCustomer)

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": newCustomer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httprouterParamsKey).(map[string]string)
	id := params["id"]

	customer, found := app.models.Customers.GetByID(id)
	if !found {
		app.notFoundResponse(w, r)
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": customer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
