package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, errCreate := user.CreateUser()
	if err != nil {
		json.NewEncoder(w).Encode(responses.Exception{Message: errCreate.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CreateUserResponse{ID: int(id), Status: "User created"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}
