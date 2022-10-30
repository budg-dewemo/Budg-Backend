package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.CreateToken).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}
