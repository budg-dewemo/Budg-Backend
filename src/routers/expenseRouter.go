package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
)

func TransactionRouter(r *mux.Router) *mux.Router {
	i := r.PathPrefix("").Subrouter()
	// allow CORS
	//i.Use(mux.CORSMethodMiddleware(i))
	//i.HandleFunc("", controllers.ValidateMiddleware(controllers.GetExpenses)).Methods("GET")
	i.HandleFunc("", controllers.GetExpenses).Methods("GET")
	//i.HandleFunc("/{id}", controllers.GetExpense).Methods("GET")
	i.HandleFunc("", controllers.CreateExpense).Methods("POST")
	//i.HandleFunc("/{id}", controllers.DeleteExpense).Methods("DELETE")
	return i
}
