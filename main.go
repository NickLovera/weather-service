package main

import (
	"log"
	"net/http"

	"github.com/NickLovera/weather-service/mgr"
	"github.com/NickLovera/weather-service/repo"
	"github.com/NickLovera/weather-service/web"
)

func main() {

	pointRepo := repo.NewPointsRepo()
	gridPointRepo := repo.NewGridPointsRepo()

	weatherSvc := mgr.NewWeatherService(pointRepo, gridPointRepo)

	server := web.NewWeatherServer(weatherSvc)

	mux := server.InitServer()

	log.Println("Serving Swagger UI at http://localhost:8080/swagger-ui/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
