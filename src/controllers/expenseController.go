package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strings"
)

type CreateExpenseResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func validateToken(r *http.Request) (models.User, error) {
	//obtener el token desde el header Authorization
	auth := r.Header.Get("Authorization")
	//separar el token del string "Bearer "
	bearerToken := strings.Split(auth, " ")[1]

	// validar el token
	token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return jwtToken, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		mapstructure.Decode(claims, &user)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			mapstructure.Decode(claims, &user)
		}
		return user, nil
	} else {
		return models.User{}, fmt.Errorf("Invalid authorization token")
	}
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
	expense := models.Expense{}
	exp, err := expense.GetExpenses(user.ID)

	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(exp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// create expense
func CreateExpense(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}

	expense := models.Expense{}
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense.UserId = user.ID
	exp, err := expense.CreateExpense()
	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CreateExpenseResponse{ID: int(exp), Status: "Expense created"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}
