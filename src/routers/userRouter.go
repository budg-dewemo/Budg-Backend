package routers

//func StudentsRouter(r *mux.Router) *mux.Router {
//	i := r.PathPrefix("/issues").Subrouter()
//	// allow CORS
//	i.Use(mux.CORSMethodMiddleware(i))
//	i.HandleFunc("", controllers.GetUsers).Methods("GET")
//	i.HandleFunc("", controllers.AddUser).Methods("POST")
//	//login
//	i.HandleFunc("/login", controllers.Login).Methods("POST")
//	i.HandleFunc("/{id}", controllers.DeleteUSer).Methods("DELETE")
//	return i
//}
