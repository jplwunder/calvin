package data

import ()

type CustomerModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone int64  `json:"phone"`
	Email string `json:"email"`
}

var customers = []CustomerModel{
	{ID: "1", Name: "John Doe", Phone: 1234567890, Email: "john.doe@example.com"},
	{ID: "2", Name: "Jane Smith", Phone: 9876543210, Email: "jane.smith@example.com"},
	{ID: "3", Name: "Alice Johnson", Phone: 5551234567, Email: "alice.johnson@example.com"},
}

func (m CustomerModel) GetAll() []CustomerModel {
	return customers
}

func (m CustomerModel) GetByID(id string) (CustomerModel, bool) {
	for _, customer := range customers {
		if customer.ID == id {
			return customer, true
		}
	}
	return CustomerModel{}, false
}

func (m CustomerModel) Insert(customer CustomerModel) {
	customers = append(customers, customer)
}
