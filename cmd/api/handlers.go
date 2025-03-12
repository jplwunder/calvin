package api

import (
	"contacts-manager/internal/data"
	"net/http"
)

type contextKey string

const httprouterParamsKey = contextKey("httprouterParams")

func (app *application) listContactsHandler(w http.ResponseWriter, r *http.Request) {
	contacts := app.models.Contacts.GetAll()

	err := app.writeJSON(w, http.StatusOK, envelope{"data": contacts}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createContactHandler(w http.ResponseWriter, r *http.Request) {
	var newContact data.Contact

	err := app.readJSON(w, r, &newContact)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.models.Contacts.Insert(newContact)

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": newContact}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getContactHandler(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httprouterParamsKey).(map[string]string)
	id := params["id"]

	contact := app.models.Contacts.GetByID(id)
	if contact == nil {
		app.notFoundResponse(w, r)
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": contact}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
