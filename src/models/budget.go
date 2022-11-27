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

func (b *Budget) GetCurrentBudget(userId int) (Budget, error) {
	if checkIfUserExists(userId) {
		var budget Budget
		InfoLogger.Println("Getting current balance for user with id: ", userId)
		query := fmt.Sprintf("SELECT id as budgetId, name as name, user_id as userId, amount as amount, start_date as startDate, end_date as endDate FROM Budget WHERE user_id = %d AND current_budget = true", userId)
		rows, err := database.QueryDB(query)
		if err != nil {
			ErrorLogger.Println(err.Error())
		}

		i := 0
		for rows.Next() {
			i++
			err = rows.Scan(&budget.BudgetId, &budget.Name, &budget.UserId, &budget.Amount, &budget.StartDate, &budget.EndDate)
			if err != nil {
				ErrorLogger.Println(err.Error())
			}
		}

		if i == 0 {
			WarningLogger.Println("No budget for user with id: ", userId)
			return budget, fmt.Errorf("No budgets for user with id: %d", userId)
		}
		if i == 1 {
			return budget, nil
		}
		WarningLogger.Println("Multiple active budgets for user with id: ", userId)
		return budget, fmt.Errorf("Multiple active budgets for user with id: %d", userId)

	} else {
		ErrorLogger.Println("User with id: ", userId, " does not exist")
		return Budget{}, fmt.Errorf("User with id: %d does not exist or is not active", userId)
	}
}

func (b *Budget) CreateBudget() (int64, error) {
	if checkIfUserExists(b.UserId) {
		InfoLogger.Println("Creating budget for user with id: ", b.UserId)
		query := fmt.Sprintf("INSERT INTO Budget (name, user_id, amount, start_date, end_date, current_budget) VALUES ('%s', %d, %f, '%s', '%s', false)", b.Name, b.UserId, b.Amount, b.StartDate, b.EndDate)
		res, err := database.InsertDB(query)
		if err != nil {
			ErrorLogger.Println("Error creating budget: ", err)
			return -1, err
		}
		b.changeCurrentBudget(res)
		return res, nil
	} else {
		ErrorLogger.Println("User with id: ", b.UserId, " does not exist")
		return -1, fmt.Errorf("User with id: %d does not exist or is not active", b.UserId)
	}

}

func (b *Budget) changeCurrentBudget(budgetID int64) (string, error) {
	if checkIfUserExists(b.UserId) {
		InfoLogger.Println("Changing current budget for user with id: ", b.UserId)
		actualBudget, err := b.GetCurrentBudget(b.UserId)
		if err != nil {
			ErrorLogger.Println("Error getting current budget: ", err)
			return "", err
		}
		query := fmt.Sprintf("UPDATE Budget SET current_budget = false WHERE id = %d", actualBudget.BudgetId)
		_, err = database.UpdateDB(query)
		if err != nil {
			ErrorLogger.Println("Error changing current budget: ", err)
			return "", err
		}
		query = fmt.Sprintf("UPDATE Budget SET current_budget = true WHERE id = %d", budgetID)
		_, err = database.UpdateDB(query)
		if err != nil {
			ErrorLogger.Println("Error changing current budget: ", err)
			return "", err
		}
		return "Current budget changed", nil

	} else {
		ErrorLogger.Println("User with id: ", b.UserId, " does not exist")
		return "", fmt.Errorf("User with id: %d does not exist or is not active", b.UserId)
	}
}
