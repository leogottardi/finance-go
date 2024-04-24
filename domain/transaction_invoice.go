package domain

import "time"

type TransactionInvoice struct {
	ID        string
	InvoiceId      string
	Transaction Transaction
	CreatedAt time.Time
}

func CreateTransactionInvoice(id string, transaction Transaction, invoiceId string) TransactionInvoice {
	return TransactionInvoice{
		ID:        id,
		Transaction: transaction,
		InvoiceId:      invoiceId,
		CreatedAt: time.Now(),
	}
}

