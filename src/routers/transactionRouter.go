package routers

//func TransactionRouter(r *mux.Router) *mux.Router {
//	i := r.PathPrefix("/transaction").Subrouter()
//	// allow CORS
//	i.Use(mux.CORSMethodMiddleware(i))
//	i.HandleFunc("", controllers.GetTransactions).Methods("GET")
//	i.HandleFunc("/id", controllers.GetTransaction).Methods("GET")
//	i.HandleFunc("", controllers.CreateTransaction).Methods("POST")
//	i.HandleFunc("/{id}", controllers.DeleteTransaction).Methods("DELETE")
//	return i
//}
