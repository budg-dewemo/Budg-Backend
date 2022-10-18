package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
)

func UserPreferencesRouter(r *mux.Router) *mux.Router {
	i := r.PathPrefix("").Subrouter()
	// allow CORS
	i.Use(mux.CORSMethodMiddleware(i))
	i.HandleFunc("", controllers.ValidateMiddleware(controllers.GetUserPreferences)).Methods("GET")
	return i
}
