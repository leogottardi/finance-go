package domain

import "time"

type Transaction struct {
	ID        string
	Amount    uint
	Type      string
	CreatedAt time.Time
}

func CreateTransaction(id string, amount uint, transactionType string) Transaction {
	return Transaction{
		ID:        id,
		Amount:    amount,
		Type:      transactionType,
		CreatedAt: time.Now(),
	}
}