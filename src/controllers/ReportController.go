package controllers

import (
	"BudgBackend/src/models"
	"BudgBackend/src/responses"
	"encoding/json"
	"net/http"
)

type chartDataResponse struct {
	ChartData models.Report `json:"chartData"`
}

func GetMonthlyReport(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}
	report := models.Report{}
	reports, err := report.GetMonthlyReport(user.ID, 12)
	chartData := chartDataResponse{ChartData: reports}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el reporte mensual"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chartData)
	return
}

func GetCategoryReports(w http.ResponseWriter, r *http.Request) {
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ErrorLogger.Println(errToken.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al validar el token"})
		return
	}
	report := models.Report{}
	reports, err := report.GetCategoryReport(user.ID, 7)
	chartData := chartDataResponse{ChartData: reports}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		ErrorLogger.Println(err.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener el reporte de categorias"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chartData)
	return
}
