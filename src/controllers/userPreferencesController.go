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
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
	category := models.Category{}
	categories, errCategory := category.GetCategories(user.ID)
	userInfo, errUser := user.GetUser()

	response.ExpenseCategories = categories
	response.User = userInfo

	if errCategory != nil || errUser != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: errCategory.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
