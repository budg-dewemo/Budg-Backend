package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func BudgetRouter(r *mux.Router) *mux.Router {
	b := r.PathPrefix("").Subrouter()
	// allow CORS
	b.Use(mux.CORSMethodMiddleware(b))
	b.HandleFunc("", controllers.GetCurrentBudget).Methods("GET")
	b.HandleFunc("", controllers.CreateBudget).Methods("POST")
	b.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return b
}
