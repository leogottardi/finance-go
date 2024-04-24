package domain

type Customer struct {
	ID   int
	Name string
}


func CreateCustomer(id int, name string) Customer {
	return Customer{
		ID:   id,
		Name: name,
	}
}
