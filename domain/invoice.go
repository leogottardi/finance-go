package domain

import (
	"strconv"
	"time"
)

type Invoice struct {
	ID        string
	AccountID string
	Amount    uint
	Term      uint
	TransactionInvoices []TransactionInvoice
	CreatedAt time.Time
}

func CreateInvoice(id string, accountId string) Invoice {
	yearMonthFormat := "200601"
	currentTime := time.Now().Format(yearMonthFormat)
	termUint, _ := strconv.ParseUint(currentTime, 10, 64)
	return Invoice{
		ID:         id,
		AccountID:  accountId,
		Amount:     0,
		Term:       uint(termUint),
		CreatedAt:  time.Now(),
	}
}

func (invoice *Invoice) AddTransaction(transaction Transaction) {
	if transaction.Type == "credit" {
		invoice.Amount += transaction.Amount
	} else {
		invoice.Amount -= transaction.Amount
	}
	transactionInvoice := CreateTransactionInvoice("1", transaction, invoice.ID)
	invoice.TransactionInvoices = append(invoice.TransactionInvoices, transactionInvoice)
}

func (invoice *Invoice) TotalCredit() uint {
	var total uint
	for _, transactionInvoice := range invoice.TransactionInvoices {
		print(transactionInvoice)
		transaction := transactionInvoice.Transaction
		if transactionInvoice.Transaction.Type == "credit" {
			total += transaction.Amount
		}
	}
	return total
}

func (invoice *Invoice) TotalDebit() uint {
	var total uint
	for _, transactionInvoice := range invoice.TransactionInvoices {
		transaction := transactionInvoice.Transaction
		if transactionInvoice.Transaction.Type == "debit" {
			total += transaction.Amount
		}
	}
	return total
}

func (invoice *Invoice) Balance() int {
	return int(invoice.TotalCredit()) - int(invoice.TotalDebit())
}