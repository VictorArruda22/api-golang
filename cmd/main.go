package main

import (
	"log"
	"os"

	"github.com/VictorArruda22/api-golang/internal/application"
	"github.com/VictorArruda22/api-golang/internal/db"
	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/VictorArruda22/api-golang/internal/repository"
	"github.com/VictorArruda22/api-golang/internal/routers"
	"github.com/VictorArruda22/api-golang/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	config := db.CreateDBConfig()

	DB, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close(DB)

	customerRepo := repository.NewCustomerRepository(DB)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := routers.RouterCustomerManagement(*customerHandler)
	application.InitApplication(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
}
