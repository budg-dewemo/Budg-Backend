package main

import (
	"BudgBackend/src/config"
	"BudgBackend/src/models"
	"BudgBackend/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	//get active budget by id
	budget := models.Budget{}
	bd, errb := budget.GetCurrentBalance(1)
	if errb != nil {
		fmt.Println(errb)
	} else {
		fmt.Println(bd)
	}

	// Examples transactions
	transactions := models.Transaction{}
	transaction, err := transactions.GetTransaction(8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(transaction)
	}
	// get all transaction for user
	trs, errt := transactions.GetTransactions(2)
	if errt != nil {
		fmt.Println(errt)
	} else {
		fmt.Println(trs)
	}

	// create transaction
	//id, errct := transactions.CreateTransaction(2, 1, 100, "test", 1)
	//if errct != nil {
	//	fmt.Println(errct)
	//} else {
	//	fmt.Println(id)
	//}

	router := routers.Routers()
	log.Fatal(http.ListenAndServe(config.ServerAddress+":8080", router))

}
