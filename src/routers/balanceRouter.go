package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func BalanceRouter(r *mux.Router) *mux.Router {
	b := r.PathPrefix("").Subrouter()
	// allow CORS
	b.Use(mux.CORSMethodMiddleware(b))
	b.HandleFunc("", controllers.GetBalance).Methods("GET")
	b.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return b
}
