package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	auth := r.PathPrefix("/api/authenticate").Subrouter()
	expenses := r.PathPrefix("/api/expenses").Subrouter()
	signup := r.PathPrefix("/api/signup").Subrouter()
	categories := r.PathPrefix("/api/categories").Subrouter()
	userPreferences := r.PathPrefix("/api/userPreferences").Subrouter()
	//enableCORS(r)
	InfoLogger.Println("CORS enabled")
	AuthRouter(auth)
	InfoLogger.Println("Auth router enabled at /api/authenticate")
	TransactionRouter(expenses)
	InfoLogger.Println("Expense router enabled at /api/expenses")
	SignUpRouter(signup)
	InfoLogger.Println("User router enabled at /api/signup")
	CategoriesRouter(categories)
	InfoLogger.Println("Category router enabled at /api/categories")
	UserPreferencesRouter(userPreferences)
	InfoLogger.Println("User preferences router enabled at /api/userPreferences")
	return r
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
