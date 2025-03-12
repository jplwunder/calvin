package data

type Models struct {
	Contacts ContactModel
}

func NewModels() Models {
	return Models{
		Contacts: ContactModel{},
	}
}

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone int64  `json:"phone"`
	Email string `json:"email"`
}
