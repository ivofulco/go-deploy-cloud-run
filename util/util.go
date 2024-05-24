package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func IsValidCEP(cep string) bool {
	return len(cep) == 8 && IsDigit(cep)
}

func IsDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func GetEnvVariable(key string) string {
	fromOS := os.Getenv(key)

	if fromOS != "" {
		return fromOS
	}

	err := godotenv.Load(".env")
	if err == nil {
		return os.Getenv(key)
	}

	err = godotenv.Load("../.env")
	if err == nil {
		return os.Getenv(key)
	}

	err = godotenv.Load("../../.env")
	if err == nil {
		return os.Getenv(key)
	}

	return fromOS
}
