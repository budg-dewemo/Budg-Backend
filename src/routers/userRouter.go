package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) *mux.Router {
	u := r.PathPrefix("").Subrouter()
	// allow CORS
	u.Use(mux.CORSMethodMiddleware(u))
	u.HandleFunc("", controllers.CreateUser).Methods("POST")
	return u
}
