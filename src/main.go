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
	router := routers.Routers()
	log.Fatal(http.ListenAndServe(config.ServerAddress+":8080", router))
}
