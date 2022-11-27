package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

type CreateCategoryResponse struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}
	category := models.Category{}
	categories, err := category.GetCategories(user.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener las categorias"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	return
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)
	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}
	category := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al decodificar el json"})
		return
	}
	category.UserId = user.ID
	categoryId, err := category.CreateCategory()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear la categoria"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateCategoryResponse{ID: categoryId, Status: "Created"})
	return
}
