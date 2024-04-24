package main

import (
	"finance-go/domain"
	"fmt"
)


func main() {
	customer := domain.CreateCustomer("1", "John Doe", "12345678900")
	account := domain.CreateAccount("1", customer.ID, 1000, 5)
	account.AddInvoice()

	account.AddTransaction(10, "credit")
	account.AddTransaction(1, "credit")

	fmt.Println(account.CurrentInvoice().TotalCredit())
}