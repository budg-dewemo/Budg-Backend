package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

func GetUserPreferences(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)
	response := responses.UserPreferencesResponse{}

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}

	userInfo, errUser := user.GetUser()

	budget := models.Budget{}
	currentBudget, errBudget := budget.GetCurrentBudget(user.ID)

	if errUser != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(errUser.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el usuario"})
		return
	}

	if errBudget != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(errBudget.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el budget"})
		return
	}

	response.User = userInfo
	response.BudgetId = currentBudget.BudgetId

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
