package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

type CreateBudgetResponse struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func CreateBudget(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	budget := models.Budget{}
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al decodificar el json"})
		return
	}

	budget.UserId = user.ID
	budgetId, err := budget.CreateBudget()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear el budget"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateBudgetResponse{ID: budgetId, Status: "success"})
	return
}

func GetCurrentBudget(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	budget := models.Budget{}
	currentBudget, err := budget.GetCurrentBudget(user.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el budget"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currentBudget)
	return
}
