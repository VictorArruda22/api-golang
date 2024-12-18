package main

import (
	"log"

	"github.com/VictorArruda22/api-golang/internal/application"
	"github.com/VictorArruda22/api-golang/internal/db"
	"github.com/VictorArruda22/api-golang/internal/routers"
)

func main() {


	config := db.CreateDBConfig()

	DB, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close(DB)

	mux := routers.RouterManagement()
	application.InitApplication(mux)
}



