package main

import (


	"github.com/VictorArruda22/api-golang/internal/application"
	"github.com/VictorArruda22/api-golang/internal/routers"
)

func main() {

	mux := routers.RouterManagement()
	application.InitApplication(mux)
}
