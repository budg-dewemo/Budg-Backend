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
	log.Println("Server started at address", config.ServerAddress+":8080")
	log.Fatal(http.ListenAndServe(config.ServerAddress+":8080", router))

}
