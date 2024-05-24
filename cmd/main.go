package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ivofulco/go-deploy-cloud-run/pkg/cep"
	"github.com/ivofulco/go-deploy-cloud-run/pkg/server"
	"github.com/ivofulco/go-deploy-cloud-run/pkg/weather"
	"github.com/ivofulco/go-deploy-cloud-run/util"
)

func main() {
	port := util.GetEnvVariable("PORT")
	apiKey := util.GetEnvVariable("WEATHER_API_KEY")

	if port == "" {
		log.Println("Missing environment variable PORT, falling back to 8080")
		port = "8080"
	}

	if apiKey == "" {
		log.Fatal("Missing environment variable WEATHER_API_KEY, there is no fall back")
	}

	weatherApi := weather.InstanceWeatherApi(apiKey)
	viaCep := cep.InstanceViaCep()

	r := server.NewWebServer(weatherApi, viaCep)

	log.Println("Starting web server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
