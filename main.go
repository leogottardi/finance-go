package main

import (
	"finance-go/domain"
	"fmt"
)


func main() {
	customer := domain.CreateCustomer(1, "John Doe")
	account := domain.CreateAccount(1, customer.ID, 1000)
	account.AddTransaction( 500, "credit", 1)
	account.AddTransaction(500, "credit", 3)

	fmt.Println(account.GetBalance())
}