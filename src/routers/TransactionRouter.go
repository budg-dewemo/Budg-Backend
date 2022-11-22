package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func TransactionRouter(r *mux.Router) *mux.Router {
	e := r.PathPrefix("").Subrouter()

	e.HandleFunc("", controllers.GetTransactions).Methods("GET")
	e.HandleFunc("/{id}", controllers.GetTransaction).Methods("GET")
	e.HandleFunc("", controllers.CreateTransaction).Methods("POST")
	e.HandleFunc("/{id}", controllers.DeleteTransaction).Methods("DELETE")
	e.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return e
}
