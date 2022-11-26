package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/repository"
	"BudgBackend/src/responses"
	"encoding/json"
	"fmt"
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

	var budgetId = 0
	id := r.URL.Query().Get("budgetId")
	if id == "" {
		budgetId = 0
	} else {
		var errGetBudget error
		budgetId, errGetBudget = strconv.Atoi(id)
		if errGetBudget != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responses.Exception{Message: errGetBudget.Error()})
			return
		}
	}

	transaction := models.Transaction{}
	transaction.UserId = user.ID
	transaction.BudgetId = budgetId

	var quantityLimit = -1
	quantity := r.URL.Query().Get("quantity")
	if quantity == "" {
		quantityLimit = -1
	} else {
		var errGetQuantity error
		quantityLimit, errGetQuantity = strconv.Atoi(quantity)
		if errGetQuantity != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responses.Exception{Message: errGetQuantity.Error()})
			return
		}
	}

	var transactions []models.Transaction
	var errorGetTransactions error
	if transaction.BudgetId == 0 {
		transactions, errorGetTransactions = transaction.GetAllTransactions(quantityLimit)
	} else {
		transactions, errorGetTransactions = transaction.GetTransactions(quantityLimit)
	}

	if errorGetTransactions != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errorGetTransactions.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
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
	transaction.UserId = user.ID
	transactions, err := transaction.GetTransaction(transactionID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
	return
}

// create transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
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
	trs, err := transaction.CreateTransaction()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateTransactionResponse{ID: int(trs), Status: "transaction created"})
	return
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
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
func PutFile(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	//get transaction id
	//transactionID, getIdErr := strconv.Atoi(mux.Vars(r)["transaction"])
	id := r.URL.Query().Get("id")
	transactionID, getIdErr := strconv.Atoi(id)

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: getIdErr.Error()})
		return
	}

	//read file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	//filename := handler.Filename

	response, error := repository.PutFile(handler, file, transactionID)
	if error != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(error)
		return
	}
	fmt.Printf("Uploaded File: %+v	", handler.Filename)

	//check if transaction exist
	transaction := models.Transaction{}
	transaction.UserId = user.ID
	trx, err := transaction.GetTransaction(transactionID)

	if trx.Id == transactionID {
		//update transaction with filepath
		transaction.FilePath = response
		filePathUpdated, errorUpdate := transaction.UpdateImagePath(transactionID)
		if errorUpdate != nil {
			fmt.Println("Error loading File")
			fmt.Println(errorUpdate)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filePathUpdated)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Transaction not found"})
		return
	}

}
