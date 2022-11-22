package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func ReportRouter(r *mux.Router) *mux.Router {
	report := r.PathPrefix("").Subrouter()
	// allow CORS
	report.Use(mux.CORSMethodMiddleware(report))
	report.HandleFunc("/monthly", controllers.GetMonthlyReport).Methods("GET")
	report.HandleFunc("/category", controllers.GetCategoryReports).Methods("GET")
	//c.HandleFunc("", controllers.CreateCategory).Methods("POST")
	//i.HandleFunc("/{id}", controllers.DeleteExpense).Methods("DELETE")
	report.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return report
}
