package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type IExpense interface {
	GetExpenses(userId int) ([]Expense, error)
	GetExpense(id string) (Expense, error)
	CreateExpense(userId int, budgetId int, amount float32, description string, categoryId int) (Expense, error)
	DeleteExpense(id int) (Expense, error)
}

type Expense struct {
	Id          int     `json:"id"`
	UserId      int     `json:"userId"`
	BudgetId    int     `json:"budgetId"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CategoryId  int     `json:"categoryId"`
	Date        string  `json:"date"`
}

var Transactions []Expense

func (t *Expense) GetExpenses(userId int) ([]Expense, error) {
	var expenses []Expense
	InfoLogger.Println("Getting transactions")
	//query := fmt.Sprintf("SELECT id as id, user_id as userId, budget_id as budgetId, amount as amount, description as description, category_id as categoryId, date as date FROM Expense WHERE user_id = %d", userId)
	query := fmt.Sprintf("SELECT id as Id, user_id as UserId, budget_id as BudgetId, amount as Amount, description as Description, category_id as CategoryId, date as Date FROM Expense WHERE user_id = %d", userId)
	rows, err := database.QueryDB(query)
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for rows.Next() {
		i++
		var expense Expense
		err = rows.Scan(&expense.Id, &expense.UserId, &expense.BudgetId, &expense.Amount, &expense.Description, &expense.CategoryId, &expense.Date)
		if err != nil {
			fmt.Println(err)
		}
		expenses = append(expenses, expense)
	}

	if i == 0 {
		return expenses, fmt.Errorf("No transactions for user with id: %d", userId)
	}
	return expenses, nil
}

func (t *Expense) GetExpense(userID int, expenseID int) (Expense, error) {
	var transaction Expense
	InfoLogger.Println("Getting expense with id: ", expenseID)
	query := fmt.Sprintf("SELECT id as id, user_id as userId, budget_id as budgetId, amount as amount, description as description, category_id as categoryId, date as date FROM Expense WHERE id = %d and user_id = %d", expenseID, userID)
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
		return Expense{}, fmt.Errorf("Expense %d not found", expenseID)
	}
	if i == 1 {
		return transaction, nil
	}
	ErrorLogger.Println("Expense not found", expenseID)
	return Expense{}, fmt.Errorf("Multiple expenses with id %d", expenseID)

}

func (t *Expense) CreateExpense() (int64, error) {
	//transaction := Expense{UserId: userId, BudgetId: budgetId, Amount: amount, Description: description, CategoryId: categoryId}
	insert := fmt.Sprintf("INSERT INTO Expense (user_id, budget_id, amount, description, category_id,date) VALUES (%d, %d, %f, '%s', %d,CURDATE())", t.UserId, t.BudgetId, t.Amount, t.Description, t.CategoryId)
	id, err := database.InsertDB(insert)
	if err != nil {
		ErrorLogger.Println("Error creating transaction: ", err)
		return 0, err

	}

	return id, nil
}

func (e *Expense) DeleteExpense(id int) (int64, error) {

	query := fmt.Sprintf("DELETE FROM Expense WHERE id = %d and user_id = %d", id, e.UserId)
	_, err := database.DeleteDB(query)
	if err != nil {
		ErrorLogger.Println("Error deleting transaction: ", err)
		return 0, err
	}
	return int64(e.Id), nil
}
