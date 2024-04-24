package domain

import (
	"strconv"
	"time"
)

type Account struct {
	ID         string
	CustomerID string
	DueDay     uint
	Limit      uint
	Invoices  []Invoice
	CreatedAt  time.Time
}

func CreateAccount(id string, customerID string, limit uint, dueDay uint) Account {
	return Account{
		ID:         id,
		CustomerID: customerID,
		Limit:      limit,
		DueDay:     dueDay,
		CreatedAt:  time.Now(),
	}
}

func (account *Account) CurrentInvoice() *Invoice {
		currentTime := time.Now().Format("200601")
		termUint, _ := strconv.ParseUint(currentTime, 10, 64)
	for _, invoice := range account.Invoices {
		if invoice.Term == uint(termUint) {
			return &invoice
		}
	}
	panic("Invoice not found")
}

func (account *Account) AddInvoice() {
	invoice := CreateInvoice("1", account.ID)
	account.Invoices = append(account.Invoices, invoice)
}

func (account *Account) AddTransaction(amount uint, transactionType string) {
	if ((transactionType == "credit") && (account.Limit - amount) < 0) {
		panic("Limit exceeded")
	}
	if transactionType == "credit" {
		account.Limit -= amount
	} else {
		account.Limit += amount
	}
	transaction := CreateTransaction("1", amount, transactionType)
	invoice := account.CurrentInvoice()
	invoice.AddTransaction(transaction)
}
