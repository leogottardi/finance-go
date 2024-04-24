package domain

type Transaction struct {
	AccountID int
	Amount uint
	Type string
	Installments uint
}

func CreateTransaction(accountID int, amount uint, transactionType string, installments uint) Transaction {
	return Transaction{
		AccountID: accountID,
		Amount: amount,
		Type: transactionType,
		Installments: installments,
	}
}