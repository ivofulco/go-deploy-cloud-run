package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/ivofulco/go-deploy-cloud-run/pkg/cep"
	"github.com/ivofulco/go-deploy-cloud-run/pkg/weather"
	"github.com/ivofulco/go-deploy-cloud-run/util"
)

func NewWebServer(weather weather.Weather, cep cep.CEP) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlerHealth)
	r.Get("/{cep}", handlerCEP(weather, cep))

	return r
}

type CEPResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	defer log.Println("App health checked")
	w.Write([]byte("Server is running"))
}

func handlerCEP(weather weather.Weather, cep cep.CEP) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cepParams := strings.TrimSpace(r.URL.Path[1:])

		if !util.IsValidCEP(cepParams) {
			message := "Invalid CEP"
			log.Println(message)
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(message))
			return
		}

		location, err := cep.FindLocation(cepParams)
		if err != nil {
			message := "CEP not Found"
			log.Println(message)
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(message))
			return
		}

		temperature, err := weather.GetTemperature(location)
		if err != nil {
			message := "Internal Server Error"
			log.Println(message)
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(message))
			return
		}

		response := CEPResponse{
			TempC: temperature.Celsius,
			TempF: temperature.Fahrenheit,
			TempK: temperature.Kelvin,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
}
