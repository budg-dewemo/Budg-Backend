package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
)

func CategoriesRouter(r *mux.Router) *mux.Router {
	i := r.PathPrefix("").Subrouter()
	// allow CORS
	i.Use(mux.CORSMethodMiddleware(i))
	i.HandleFunc("", controllers.GetCategories).Methods("GET")
	//i.HandleFunc("/{id}", controllers.GetExpense).Methods("GET")
	//i.HandleFunc("", controllers.CreateCategory).Methods("POST")
	//i.HandleFunc("/{id}", controllers.DeleteExpense).Methods("DELETE")
	return i
}
