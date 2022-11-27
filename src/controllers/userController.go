package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

type CreateUserResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al decodificar el usuario"})
		return
	}

	id, errCreate := user.CreateUser()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(errCreate)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear el usuario"})
		return
	}

	if id == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(errCreate)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear el usuario"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{ID: int(id), Status: "User created"})
	return
}
