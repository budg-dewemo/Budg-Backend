package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type IBudget interface {
	GetCurrentBalance(userId int) []Budget
	CreateBalance(name string, userId int, amount float32, startDate string, endDate string) int64
	checkIfUserExists(userId int) bool
}

type Budget struct {
	BudgetId  int     `json:"budgetId"`
	Name      string  `json:"name"`
	UserId    int     `json:"userId"`
	Amount    float32 `json:"amount"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
}

func checkIfUserExists(userId int) bool {
	query := fmt.Sprintf("SELECT id FROM User WHERE id = %d and active = true", userId)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	id := 0
	for rows.Next() {
		i++
		err = rows.Scan(&id)
	}
	switch i {
	case 0:
		return false
	default:
		if id == userId {
			return true
		} else {
			return false
		}
	}
}

func (b *Budget) GetCurrentBalance(userId int) (Budget, error) {
	if checkIfUserExists(userId) {
		var budget Budget
		InfoLogger.Println("Getting current balance for user with id: ", userId)
		query := fmt.Sprintf("SELECT id as budgetId, name as name, user_id as userId, amount as amount, start_date as startDate, end_date as endDate FROM Budget WHERE user_id = %d AND current_budget = true", userId)
		rows, err := database.QueryDB(query)
		if err != nil {
			fmt.Println(err)
		}

		i := 0
		for rows.Next() {
			i++
			err = rows.Scan(&budget.BudgetId, &budget.Name, &budget.UserId, &budget.Amount, &budget.StartDate, &budget.EndDate)
			if err != nil {
				fmt.Println(err)
			}
		}

		switch i {
		case 0:
			ErrorLogger.Println("No budget for user with id: ", userId)
			return budget, fmt.Errorf("No budgets for user with id: %d", userId)
		case 1:
			InfoLogger.Println("Found active budget for user with id: ", userId)
			return budget, nil
		default:
			ErrorLogger.Println("Multiple active budgets for user with id: ", userId)
			return budget, fmt.Errorf("Multiple active budgets for user with id: %d", userId)
		}
	} else {
		ErrorLogger.Println("User with id: ", userId, " does not exist")
		return Budget{}, fmt.Errorf("User with id: %d does not exist or is not active", userId)
	}
}
