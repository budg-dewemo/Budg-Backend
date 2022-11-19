package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CreateTransactionResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)
	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}
	transaction := models.Transaction{}
	exp, err := transaction.GetTransactions(user.ID)

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

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)
	transactionID, getIdErr := strconv.Atoi(mux.Vars(r)["id"])

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
	transaction := models.Transaction{}
	exp, err := transaction.GetTransaction(user.ID, transactionID)

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

// create transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
	}

	transaction := models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	if transaction.Amount <= 0 || transaction.BudgetId <= 0 || transaction.CategoryId <= 0 || transaction.Description == "" || transaction.Date == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid transaction format expected: {\"amount\": , \"budgetId\": , \"categoryId\": , \"description\": \"\", \"date\": \"\"}"})
		return
	}

	transaction.UserId = user.ID
	exp, err := transaction.CreateTransaction()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateTransactionResponse{ID: int(exp), Status: "transaction created"})
	return
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
	}

	transaction := models.Transaction{}
	transactionID, getIdErr := strconv.Atoi(mux.Vars(r)["id"])

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: getIdErr.Error()})
		return
	}

	transaction.UserId = user.ID
	_, err := transaction.DeleteTransaction(transactionID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CreateTransactionResponse{ID: int(transactionID), Status: "transaction deleted"})
	return
}
