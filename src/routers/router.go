package routers

import (
	"BudgBackend/src/responses"
	"encoding/json"
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
	enableCORS(r)

	//api version 1
	v1 := r.PathPrefix("/api/v1").Subrouter()

	auth := v1.PathPrefix("/authenticate").Subrouter()
	transactions := v1.PathPrefix("/transactions").Subrouter()
	signup := v1.PathPrefix("/signup").Subrouter()
	categories := v1.PathPrefix("/categories").Subrouter()
	userPreferences := v1.PathPrefix("/userPreferences").Subrouter()
	budget := v1.PathPrefix("/budget").Subrouter()
	balance := v1.PathPrefix("/balance").Subrouter()
	report := v1.PathPrefix("/report").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	InfoLogger.Println("CORS enabled")

	AuthRouter(auth)
	InfoLogger.Println("Auth router enabled at /api/v1/authenticate")
	TransactionRouter(transactions)
	InfoLogger.Println("Expense router enabled at /api/v1/transactions")
	SignUpRouter(signup)
	InfoLogger.Println("User router enabled at /api/v1/signup")
	CategoriesRouter(categories)
	InfoLogger.Println("Category router enabled at /api/v1/categories")
	UserPreferencesRouter(userPreferences)
	InfoLogger.Println("User preferences router enabled at /api/v1/userPreferences")
	BudgetRouter(budget)
	InfoLogger.Println("Budget router enabled at /api/v1/budget")
	BalanceRouter(balance)
	InfoLogger.Println("Balance router enabled at /api/v1/balance")
	ReportRouter(report)
	InfoLogger.Println("Report router enabled at /api/v1/report")
	return r
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Exception{Message: "path not found"})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses.Exception{Message: "method not allowed"})
	}
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
