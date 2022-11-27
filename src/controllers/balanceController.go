package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"math"
	"net/http"
)

type balanceResponse struct {
	CurrentBalance float64 `json:"currentBalance"`
	TotalIncome    float64 `json:"totalIncome"`
	TotalExpenses  float64 `json:"totalExpenses"`
	TotalBudget    float64 `json:"totalBudget"`
}

func GetBalance(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "No se pudo validar el token"})
		return
	}

	budget := models.Budget{}
	currentBudget, err := budget.GetCurrentBudget(user.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el budget activo"})
		return
	}

	transaction := models.Transaction{}
	transaction.BudgetId = currentBudget.BudgetId
	transaction.UserId = user.ID
	transactions, err := transaction.GetTransactions(-1)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener las transacciones"})
		return
	}
	var totalIncome float64
	var totalExpenses float64
	for _, t := range transactions {
		if t.Type == "income" {
			totalIncome += float64(t.Amount)

		} else {
			totalExpenses += float64(t.Amount)
		}
	}

	currentBalance := (float64(currentBudget.Amount) + totalIncome) - totalExpenses
	balanceResponse := balanceResponse{CurrentBalance: roundFloat(currentBalance, 2), TotalIncome: roundFloat(totalIncome, 2), TotalExpenses: roundFloat(totalExpenses, 2), TotalBudget: math.Floor(float64(currentBudget.Amount*100) / 100)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balanceResponse)
	return
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
