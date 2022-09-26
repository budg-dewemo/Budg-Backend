package main

import (
	"BudgBackend/src/config"
	"BudgBackend/src/routers"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	////get active budget by id
	//budget := models.Budget{}
	//bd, errb := budget.GetCurrentBalance(1)
	//if errb != nil {
	//	fmt.Println(errb)
	//} else {
	//	fmt.Println(bd)
	//}
	//
	//// Examples Expenses
	//expenses := models.Expense{}
	//expense, err := expenses.GetExpense(8)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(expense)
	//}
	//// get all expenses for user
	//trs, errt := expenses.GetExpenses(2)
	//if errt != nil {
	//	fmt.Println(errt)
	//} else {
	//	fmt.Println(trs)
	//}

	// create Expenses
	//id, errct := expenses.CreateExpense(2, 1, 100, "test", 1)
	//if errct != nil {
	//	fmt.Println(errct)
	//} else {
	//	fmt.Println(id)
	//}

	router := routers.Routers()
	log.Println("Server started at address", config.ServerAddress+":8080")
	log.Fatal(http.ListenAndServe(config.ServerAddress+":8080", router))

}
