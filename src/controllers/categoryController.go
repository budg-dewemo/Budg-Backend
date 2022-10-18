package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

type CreateCategoryResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
	category := models.Category{}
	categories, err := category.GetCategories(user.ID)

	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

//
//// create expense
//func CreateExpense(w http.ResponseWriter, r *http.Request) {
//	user, errToken := validateToken(r)
//
//	if errToken != nil {
//		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusUnauthorized)
//	}
//
//	expense := models.Expense{}
//	err := json.NewDecoder(r.Body).Decode(&expense)
//	if err != nil {
//		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	expense.UserId = user.ID
//	exp, err := expense.CreateExpense()
//	if err != nil {
//		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(CreateExpenseResponse{ID: int(exp), Status: "Expense created"})
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	return
//}
