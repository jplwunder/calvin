package data

var Contacts = []Contact{
	{ID: "1", Name: "John Doe", Phone: 1234567890, Email: "john.doe@example.com"},
	{ID: "2", Name: "Jane Smith", Phone: 9876543210, Email: "jane.smith@example.com"},
	{ID: "3", Name: "Alice Johnson", Phone: 5551234567, Email: "alice.johnson@example.com"},
}

type ContactModel struct{}

func (m ContactModel) GetAll() []Contact {
	return Contacts
}

func (m ContactModel) GetByID(id string) *Contact {
	for _, contact := range Contacts {
		if contact.ID == id {
			return &contact
		}
	}
	return nil
}

func (m ContactModel) Insert(contact Contact) {
	Contacts = append(Contacts, contact)
}
