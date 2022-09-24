package models

type ITransaction interface {
	GetTransactions() []Transaction
	GetTransaction(id string) Transaction
	CreateTransaction(transaction Transaction) Transaction
	DeleteTransaction(id string) Transaction
}

type Transaction struct {
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

var Transactions []Transaction

func (t *Transaction) GetTransactions() []Transaction {
	return Transactions
}

func (t *Transaction) GetTransaction(id string) Transaction {
	for _, transaction := range Transactions {
		if transaction.Category == id {
			return transaction
		}
	}
	return Transaction{}
}

func (t *Transaction) CreateTransaction(transaction Transaction) Transaction {
	Transactions = append(Transactions, transaction)
	return transaction
}

func (t *Transaction) DeleteTransaction(id string) Transaction {
	for index, transaction := range Transactions {
		if transaction.Category == id {
			Transactions = append(Transactions[:index], Transactions[index+1:]...)
			return transaction
		}
	}
	return Transaction{}
}
