package data

type Models struct {
	Customers CustomerModel
}

func NewModels() Models {
	return Models{
		Customers: CustomerModel{},
	}
}
