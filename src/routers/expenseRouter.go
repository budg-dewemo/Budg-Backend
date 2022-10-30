package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func TransactionRouter(r *mux.Router) *mux.Router {
	e := r.PathPrefix("").Subrouter()
	// allow CORS
	//i.Use(mux.CORSMethodMiddleware(i))
	//i.HandleFunc("", controllers.ValidateMiddleware(controllers.GetExpenses)).Methods("GET")
	e.HandleFunc("", controllers.GetExpenses).Methods("GET")
	e.HandleFunc("/{id}", controllers.GetExpense).Methods("GET")
	e.HandleFunc("", controllers.CreateExpense).Methods("POST")
	e.HandleFunc("/{id}", controllers.DeleteExpense).Methods("DELETE")
	e.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return e
}
