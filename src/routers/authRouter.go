package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.CreateToken).Methods("POST")
	return a
}
