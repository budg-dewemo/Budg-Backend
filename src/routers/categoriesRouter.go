package routers

import (
	"BudgBackend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func CategoriesRouter(r *mux.Router) *mux.Router {
	c := r.PathPrefix("").Subrouter()
	// allow CORS
	c.Use(mux.CORSMethodMiddleware(c))
	c.HandleFunc("", controllers.GetCategories).Methods("GET")
	//i.HandleFunc("/{id}", controllers.GetExpense).Methods("GET")
	c.HandleFunc("", controllers.CreateCategory).Methods("POST")
	//i.HandleFunc("/{id}", controllers.DeleteExpense).Methods("DELETE")
	c.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return c
}
