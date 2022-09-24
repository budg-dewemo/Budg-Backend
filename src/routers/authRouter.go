package routers

import (
	"BudgBackend/src/controllers"
	"fmt"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	fmt.Println("Starting the application...")
	a.HandleFunc("/authenticate", controllers.CreateToken).Methods("POST")
	a.HandleFunc("/protected", controllers.ValidateMiddleware(controllers.ProtectedEndpoint)).Methods("GET")
	return a
}
