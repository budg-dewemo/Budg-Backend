package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type ITransaction interface {
	GetTransactions(userId int) []Transaction
	GetTransaction(id string) Transaction
	CreateTransaction(userId int, budgetId int, amount float32, description string, categoryId int) Transaction
	DeleteTransaction(id int) Transaction
}

type Transaction struct {
	Id          int     `json:"id"`
	UserId      int     `json:"userId"`
	BudgetId    int     `json:"budgetId"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CategoryId  int     `json:"categoryId"`
	Date        string  `json:"date"`
}

var Transactions []Transaction

func (t *Transaction) GetTransactions(userId int) ([]Transaction, error) {
	var transactions []Transaction
	InfoLogger.Println("Getting transactions")
	query := fmt.Sprintf("SELECT id as id, user_id as userId, budget_id as budgetId, amount as amount, description as description, category_id as categoryId, date as date FROM Expense WHERE user_id = %d", userId)
	rows, err := database.QueryDB(query)
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for rows.Next() {
		i++
		var transaction Transaction
		err = rows.Scan(&transaction.Id, &transaction.UserId, &transaction.BudgetId, &transaction.Amount, &transaction.Description, &transaction.CategoryId, &transaction.Date)
		if err != nil {
			fmt.Println(err)
		}
		transactions = append(transactions, transaction)
	}

	switch i {
	case 0:
		ErrorLogger.Println("No transactions for user with id: ", userId)
		return transactions, fmt.Errorf("No transactions for user with id: %d", userId)
	default:
		InfoLogger.Println("Found transactions for user with id: ", userId)
		return transactions, nil
	}
}

func (t *Transaction) GetTransaction(id int) (Transaction, error) {
	var transaction Transaction
	InfoLogger.Println("Getting transaction with id: ", id)
	query := fmt.Sprintf("SELECT id as id, user_id as userId, budget_id as budgetId, amount as amount, description as description, category_id as categoryId, date as date FROM Expense WHERE id = %d", id)
	rows, err := database.QueryDB(query)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for rows.Next() {
		i++
		err = rows.Scan(&transaction.Id, &transaction.UserId, &transaction.BudgetId, &transaction.Amount, &transaction.Description, &transaction.CategoryId, &transaction.Date)
		if err != nil {
			fmt.Println(err)
		}
	}
	switch i {
	case 0:
		ErrorLogger.Println("No transaction with id: ", id)
		return Transaction{}, fmt.Errorf("No transaction with id: %d", id)
	case 1:
		InfoLogger.Println("Found transaction with id: ", id)
		return transaction, nil
	default:
		ErrorLogger.Println("Multiple transactions with id: ", id)
		return Transaction{}, fmt.Errorf("Multiple transactions with id: %d", id)
	}

}

func (t *Transaction) CreateTransaction(userId int, budgetId int, amount float32, description string, categoryId int) (int64, error) {
	//transaction := Transaction{UserId: userId, BudgetId: budgetId, Amount: amount, Description: description, CategoryId: categoryId}
	insert := fmt.Sprintf("INSERT INTO expense (user_id, budget_id, amount, description, category_id,date) VALUES (%d, %d, %f, '%s', %d,CURDATE())", userId, budgetId, amount, description, categoryId)
	id, err := database.InsertDB(insert)
	if err != nil {
		ErrorLogger.Println("Error creating transaction: ", err)
		return 0, err

	}

	return id, nil
}

func (t *Transaction) DeleteTransaction(id int) Transaction {
	for index, transaction := range Transactions {
		if transaction.CategoryId == id {
			Transactions = append(Transactions[:index], Transactions[index+1:]...)
			return transaction
		}
	}
	return Transaction{}
}
