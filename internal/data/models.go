package data

import (
	"calvin/connections"
)

type Models struct {
	Customers CustomerModel
}

func NewModels() Models {
	// Get the MongoDB database connection
	db := connections.GetDB()

	return Models{
		Customers: CustomerModel{DB: db},
	}
}

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone int64  `json:"phone"`
	Email string `json:"email"`
}
