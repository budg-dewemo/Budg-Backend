package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CreateExpenseResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}
	expense := models.Expense{}
	exp, err := expense.GetExpenses(user.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exp)
	return
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)
	expenseID, getIdErr := strconv.Atoi(mux.Vars(r)["id"])

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}
	expense := models.Expense{}
	exp, err := expense.GetExpense(user.ID, expenseID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exp)
	return
}

// create expense
func CreateExpense(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
	}

	expense := models.Expense{}
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	if expense.Amount <= 0 || expense.BudgetId <= 0 || expense.CategoryId <= 0 || expense.Description == "" || expense.Date == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid expense format expected: {\"amount\": , \"budgetId\": , \"categoryId\": , \"description\": \"\", \"date\": \"\"}"})
		return
	}

	expense.UserId = user.ID
	exp, err := expense.CreateExpense()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateExpenseResponse{ID: int(exp), Status: "Expense created"})
	return
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
	}

	expense := models.Expense{}
	expenseID, getIdErr := strconv.Atoi(mux.Vars(r)["id"])

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: getIdErr.Error()})
		return
	}

	expense.UserId = user.ID
	_, err := expense.DeleteExpense(expenseID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CreateExpenseResponse{ID: int(expenseID), Status: "Expense deleted"})
	return
}
