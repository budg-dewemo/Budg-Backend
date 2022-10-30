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
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}
	category := models.Category{}
	categories, errCategory := category.GetCategories(user.ID)

	userInfo, errUser := user.GetUser()

	response.ExpenseCategories = categories
	response.User = userInfo

	if errUser != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errUser.Error()})
		return
	}

	if errCategory != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errCategory.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
