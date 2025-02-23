package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/contacts", getContacts)
	router.POST("/contacts", postContacts)
	router.GET("/contacts/:id", getContactByID)

	router.Run("0.0.0.0:8080")
}

type contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone int64  `json:"phone"`
	Email string `json:"email"`
}

// constacts slice to seed record album data.
var contacts = []contact{
	{ID: "1", Name: "Fredi William Wunderlich", Phone: 5511989669526, Email: "fredi.wunder@gmail.com"},
	{ID: "2", Name: "Letícia Rahel Lopes Wunderlich", Phone: 5511995265188, Email: "leticia.wunder@gmail.com"},
	{ID: "3", Name: "Izalira Ferreira Lopes Wunderlich", Phone: 5511989694053, Email: "izaliralopes@gmail.com"},
}

// getContacts responds with the list of all contacts as JSON.
func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

// postAlbums adds an album from JSON received in the request body.
func postContacts(c *gin.Context) {
	var newContact contact

	// Call BindJSON to bind the received JSON to
	// newContact.
	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	// Add the new album to the slice.
	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

// getContactByID locates the contact whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getContactByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range contacts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found"})
}
