package main

import (
	"contacts-manager/internal/generate"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var contacts = []generate.Contact{
	{ID: "1", Name: "Fredi William Wunderlich", Phone: 5511989669526, Email: "fredi.wunder@gmail.com"},
	{ID: "2", Name: "Letícia Rahel Lopes Wunderlich", Phone: 5511995265188, Email: "leticia.wunder@gmail.com"},
	{ID: "3", Name: "Izalira Ferreira Lopes Wunderlich", Phone: 5511989694053, Email: "izaliralopes@gmail.com"},
}

func (app *Application) getContacts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	env := envelope{
		"data": contacts,
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) createContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newContact generate.Contact
	app.logger.Info("Creating contact with the JSON data")

	err := app.readJSON(w, r, &newContact)
	if err != nil {
		app.logger.Error("failed to decode contact data", "error", err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	app.logger.Info("Storing contact info")
	contacts = append(contacts, newContact)

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"data": newContact,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	app.logger.Info("New contact created")
}

func (app *Application) getContactByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	for _, contact := range contacts {
		if contact.ID == id {
			app.writeJSON(w, http.StatusOK, envelope{
				"data": contact,
			}, nil)
			return
		}
	}
	app.notFoundResponse(w, r)
}
