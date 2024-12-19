package main

import (
	"log"

	"github.com/VictorArruda22/api-golang/internal/application"
	"github.com/VictorArruda22/api-golang/internal/db"
	"github.com/VictorArruda22/api-golang/internal/handler"
	"github.com/VictorArruda22/api-golang/internal/repository"
	"github.com/VictorArruda22/api-golang/internal/routers"
	"github.com/VictorArruda22/api-golang/internal/service"
)

func main() {

	config := db.CreateDBConfig()

	DB, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close(DB)

	categoryRepository := repository.NewCategoryRepository(DB)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	mux := routers.RouterManagement(categoryHandler)
	application.InitApplication(mux)
}
