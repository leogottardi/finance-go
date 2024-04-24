package domain

type Account struct {
	ID  int
	CustomerID int
	Limit uint
	Transactions []Transaction
}

func CreateAccount(id int, customerID int, limit uint) Account {
	return Account{
		ID: id,
		CustomerID: customerID,
		Limit: limit,
	}
}

func (account *Account) AddTransaction(amount uint, transactionType string, installments uint) {
	balance := account.GetBalance()
	if balance < amount {
		println("Transaction declined")
		return
	}
	account.Transactions = append(account.Transactions, Transaction{account.ID, amount, transactionType, installments})
}

func (account *Account) GetBalance() uint {
	var balance uint
	for _, transaction := range account.Transactions {
		if transaction.Type == "credit" {
			balance += transaction.Amount
		} else {
			balance -= transaction.Amount
		}
	}
	return account.Limit - balance
}