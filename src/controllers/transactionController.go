package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/repository"
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
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
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
			ErrorLogger.Println(errGetBudget.Error())
			json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el id del budget"})
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
			ErrorLogger.Println(errGetQuantity.Error())
			json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener la cantidad de transacciones"})
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
		ErrorLogger.Println(errorGetTransactions.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener las transacciones"})
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
		json.NewEncoder(w).Encode(responses.Exception{Message: getIdErr.Error()})
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el id de la transacción"})
		return
	}

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}
	transaction := models.Transaction{}
	transaction.UserId = user.ID
	transactions, err := transaction.GetTransaction(transactionID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener la transacción"})
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
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	transaction := models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el json"})
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
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear la transacción"})
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
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	transaction := models.Transaction{}
	transactionID, getIdErr := strconv.Atoi(mux.Vars(r)["id"])

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(getIdErr.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el id de la transacción"})
		return
	}

	transaction.UserId = user.ID
	_, err := transaction.DeleteTransaction(transactionID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al eliminar la transacción"})
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
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	id := r.URL.Query().Get("id")
	transactionID, getIdErr := strconv.Atoi(id)

	if getIdErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(getIdErr.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el id de la transacción"})
		return
	}

	//read file
	file, handler, err := r.FormFile("file")
	if err != nil {
		ErrorLogger.Println("Error Retrieving the File ", err.Error())

		return
	}
	defer file.Close()
	//filename := handler.Filename

	response, error := repository.PutFile(handler, file, transactionID)
	if error != nil {
		ErrorLogger.Println("Error Retrieving the File ", error.Error())

		return
	}

	//check if transaction exist
	transaction := models.Transaction{}
	transaction.UserId = user.ID
	trx, err := transaction.GetTransaction(transactionID)

	if trx.Id == transactionID {
		//update transaction with filepath
		transaction.FilePath = response
		filePathUpdated, errorUpdate := transaction.UpdateImagePath(transactionID)
		if errorUpdate != nil {
			ErrorLogger.Println("Error Retrieving the File ", errorUpdate.Error())
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
