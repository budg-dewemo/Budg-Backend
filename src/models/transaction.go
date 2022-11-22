package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type ITransaction interface {
	GetTransactions() ([]Transaction, error)
	GetTransaction(id string) (Transaction, error)
	CreateTransaction() (Transaction, error)
	DeleteTransaction(id int) (Transaction, error)
}

type Transaction struct {
	Id          int     `json:"id"`
	UserId      int     `json:"userId"`
	BudgetId    int     `json:"budgetId"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CategoryId  int     `json:"categoryId"`
	Date        string  `json:"date"`
	Type        string  `json:"type"`
	FilePath    string  `json:"filePath"`
}

var Transactions []Transaction

func (t *Transaction) GetTransactions(limit int) ([]Transaction, error) {
	var transactions []Transaction
	InfoLogger.Println("Getting transactions")
	query := ""
	if limit == -1 {
		query = fmt.Sprintf("SELECT id as Id, user_id as UserId, budget_id as BudgetId, amount as Amount, description as Description, category_id as CategoryId, date as Date FROM User_transaction WHERE user_id = %d ORDER BY date DESC", t.UserId)

	} else {
		query = fmt.Sprintf("SELECT id as Id, user_id as UserId, budget_id as BudgetId, amount as Amount, description as Description, category_id as CategoryId, date as Date FROM User_transaction WHERE user_id = %d ORDER BY date DESC LIMIT %d", t.UserId, limit)
	}
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

	if i == 0 {
		return transactions, fmt.Errorf("No transactions for user with id: %d", t.UserId)
	}
	return transactions, nil
}

func (t *Transaction) GetTransaction(transactionID int) (Transaction, error) {
	var transaction Transaction
	InfoLogger.Println("Getting transaction with id: ", transactionID)
	query := fmt.Sprintf("SELECT id as id, user_id as userId, budget_id as budgetId, amount as amount, description as description, category_id as categoryId, date as date FROM User_transaction WHERE id = %d and user_id = %d", transactionID, t.UserId)
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

	if i == 0 {
		return Transaction{}, fmt.Errorf("transaction %d not found", transactionID)
	}
	if i == 1 {
		return transaction, nil
	}
	ErrorLogger.Println("transaction not found", transactionID)
	return Transaction{}, fmt.Errorf("Multiple transactions with id %d", transactionID)

}

func (t *Transaction) CreateTransaction() (int64, error) {
	//transaction := transaction{UserId: userId, BudgetId: budgetId, Amount: amount, Description: description, CategoryId: categoryId}

	insert := fmt.Sprintf("INSERT INTO User_transaction (user_id, budget_id, amount, description, category_id,type,date) VALUES (%d, %d, %f, '%s', %d, '%s', CURDATE())", t.UserId, t.BudgetId, t.Amount, t.Description, t.CategoryId, t.Type)
	print(insert)
	id, err := database.InsertDB(insert)
	if err != nil {
		ErrorLogger.Println("Error creating transaction: ", err)
		return 0, err

	}

	return id, nil
}

func (e *Transaction) DeleteTransaction(id int) (int64, error) {

	query := fmt.Sprintf("DELETE FROM User_transaction WHERE id = %d and user_id = %d", id, e.UserId)
	_, err := database.DeleteDB(query)
	if err != nil {
		ErrorLogger.Println("Error deleting transaction: ", err)
		return 0, err
	}
	return int64(e.Id), nil
}
