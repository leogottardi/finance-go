package domain

import "time"

type Customer struct {
	ID        string
	Name      string
	Document  string
	CreatedAt time.Time
}

func CreateCustomer(id string, name string, document string) Customer {
	return Customer{
		ID:        id,
		Name:      name,
		Document:  document,
		CreatedAt: time.Now(),
	}
}
